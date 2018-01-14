package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	wadio "github.com/johnsto/go-randoom/pkg/wad"
	"github.com/pkg/errors"
)

type Archive struct {
	Wads     map[string]*wadio.Wad
	Dehacked map[string][]byte
}

func (a *Archive) GetFirstPlayableWad() (name string, wadfile *wadio.Wad, level string) {
	for n, w := range a.Wads {
		for _, lump := range w.Lumps {
			if wadio.IsMapName(lump.Name) {
				if level == "" || level > lump.Name {
					wadfile = w
					level = lump.Name
					name = n
				}
			}
		}
	}
	return
}

func Parse(r io.ReaderAt, size int64) (*Archive, error) {
	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}

	archive := &Archive{
		Wads: make(map[string]*wadio.Wad),
		Dehacked: make(map[string][]byte),
	}

	for _, f := range zr.File {
		fr, err := f.Open()
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't open zip entry %s", f.Name)
		}
		defer fr.Close()

		ext := filepath.Ext(strings.ToLower(f.Name))
		switch ext {
		case ".wad":
			wad, err := wadio.Read(fr)
			if err != nil {
				return nil, errors.Wrapf(err, "couldn't parse WAD in %s", f.Name)
			}
			archive.Wads[f.Name] = wad
		case ".deh":
			// Probable Dehacked patch
			data, err := ioutil.ReadAll(fr)
			if err != nil {
				return nil, errors.Wrapf(err, "couldn't read DEH patchh in %s", f.Name)
			}
			archive.Dehacked[f.Name] = data
		}
	}

	return archive, nil
}
