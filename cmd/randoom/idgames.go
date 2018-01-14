package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"

	"github.com/johnsto/randoom/pkg/idgames"
	wadio "github.com/johnsto/randoom/pkg/wad"
	"github.com/pkg/errors"
)

type (
	Game = wadio.Game
	GameType string
)

const (
	Doom    Game = "doom"
	Doom2            = "doom2"
	Heretic          = "heretic"
	Hexen            = "hexen"
	Strife           = "strife"

	Mission    GameType = "mission"
	Deathmatch          = "deathmatch"
)

// GetLevelsPath returns an idgames path for the given game and mode.
// If the "Mission" game type is chosen, a random sub-directory is returned
// from the alphabetical categories.
func GetLevelsPath(n Game, t GameType) (string, error) {
	missionDirs := []string{"a-c", "d-f", "g-i", "j-l", "m-o", "p-r", "s-u", "v-z", "0-9"}

	switch t {
	case Mission:
		// FIXME: maps have non-uniform chance of being selected as each mission directory
		// contains a different number of files. Perhaps build up a cache each time GetRandomLevel
		// is called and select from the cache instead.
		dir := missionDirs[rand.Intn(len(missionDirs))]
		return "levels/" + string(n) + "/" + string(dir) + "/", nil
	case Deathmatch:
		return "levels/" + string(n) + "/deathmatch/", nil
	default:
		return "", errors.Errorf("unsupported game type %v", t)
	}
}

// GetLevelInPath returns a random file from the given path.
func GetRandomLevel(path string, cfg *idgames.Config) (*idgames.File, error) {
	files, err := idgames.FetchFiles(path, cfg)
	if err != nil {
		return nil, err
	}

	file := files[rand.Intn(len(files))]

	return &file, nil
}

// FetchArchive retrieves and parses the zip-encoded archive at the given URL.
func FetchArchive(url string) (*Archive, error) {
	buf := new(bytes.Buffer)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't fetch URL %s")
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/zip" {
		return nil, errors.Errorf("unexpected content type %s", contentType)
	}

	// Copy into buffer
	_, err = io.Copy(buf, resp.Body)

	// Parse archive in memory.
	fmt.Println("* Deciphering content...")
	rd := bytes.NewReader(buf.Bytes())
	return Parse(rd, rd.Size())
}

// SaveWad saves the WAD to the specified file.
func SaveWad(filename string, wad *wadio.Wad) (err error) {
	wadfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer wadfile.Close()

	_, err = io.Copy(wadfile, wad.Reader())
	return err
}

// SanitizeFilename removes any path information from the given path.
func SanitizeFilename(n string) string {
	base := path.Base(n)
	if base == "." {
		return ""
	}
	return base
}
