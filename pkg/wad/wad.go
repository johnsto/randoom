package wad

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"io"
	"strings"

	"github.com/pkg/errors"
)

const (
	IWAD Type = "IWAD"
	PWAD      = "PWAD"

	Doom    Game = "doom"
	Doom2        = "doom2"
	Heretic      = "heretic"
	Hexen        = "hexen"
	Strife       = "strife"
)

type (
	Type string
	Game string

	GameVersion struct {
		Game    Game
		Weight  int
		Version string
	}

	// Wad represents the data of a WAD file.
	Wad struct {
		Type  Type
		Lumps []Lump

		data []byte
	}

	Lump struct {
		Name string
	}
)

// Reader returns a new reader containing the raw WAD data.
func (w *Wad) Reader() *bytes.Reader {
	return bytes.NewReader(w.data)
}

// MD5 returns the MD5 of the WAD file.
func (w *Wad) MD5() (string, error) {
	h := md5.New()
	if _, err := w.Reader().WriteTo(h); err != nil {
		return "", err
	}
	return string(h.Sum(nil)), nil
}

// GameVersion returns the game and version of the file, assuming it's a known revision
// of a game's IWAD.
func (w *Wad) GameVersion() (GameVersion, error) {
	h, err := w.MD5()
	if err != nil {
		return GameVersion{}, err
	}

	if ver, ok := IWADs[h]; ok {
		return ver, nil
	}

	return GameVersion{}, nil
}

func ReadMagic(r io.Reader) (string, error) {
	var magic [4]byte

	if err := binary.Read(r, binary.LittleEndian, &magic); err != nil {
		return "", errors.Wrapf(err, "couldn't parse WAD header")
	}

	return string(magic[:]), nil
}

// Read reads a WAD file from the given Reader.
func Read(r io.Reader) (*Wad, error) {
	// Read WAD into RAM so it can be seeked.
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r); err != nil {
		return nil, errors.Wrapf(err, "couldn't read WAD into memory")
	}

	rd := bytes.NewReader(buf.Bytes())

	// Read header.
	header := struct {
		Magic           [4]byte
		Entries         int32
		DirectoryOffset int32
	}{}

	if err := binary.Read(rd, binary.LittleEndian, &header); err != nil {
		return nil, errors.Wrapf(err, "couldn't parse WAD header")
	}

	if string(header.Magic[:]) != "PWAD" {
		return nil, errors.Errorf("unknown file type %s", string(header.Magic[:]))
	}

	// Read directory.
	if _, err := rd.Seek(int64(header.DirectoryOffset), io.SeekStart); err != nil {
		return nil, errors.Wrapf(err, "couldn't seek to directory")
	}

	type entry struct {
		Offset int32
		Size   int32
		Name   [8]byte
	}

	directory := make([]entry, header.Entries)

	if err := binary.Read(rd, binary.LittleEndian, &directory); err != nil {
		return nil, errors.Wrapf(err, "couldn't parse WAD directory")
	}

	// Generate return value.
	wad := &Wad{
		Type:  Type(header.Magic[:]),
		Lumps: make([]Lump, header.Entries),
		data:  buf.Bytes(),
	}
	for i, e := range directory {
		wad.Lumps[i].Name = strings.TrimRight(string(e.Name[:]), "\x00")
	}

	return wad, nil
}
