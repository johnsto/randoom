package wad

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var CommonIWADs = []string{
	// Doom
	"doom.wad",
	"doomu.wad",
	"bfgdoom.wad",
	"doombfg.wad",
	// Doom2
	"doom2f.wad",
	"doom2.wad",
	"bfgdoom2.wad",
	"doom2bfg.wad",
	// Heretic
	"heretic.wad",
	"hereticsr.wad",
	// Hexen
	"hexen.wad",
	"hexdd.wad",
	// Strife
	"strife1.wad",
}

// GetPaths returns a list of paths where Doom IWAD files are typically searched for,
// in order of preference, with most preferred paths first.
func GetPaths() []string {
	paths := make([]string, 0)

	paths = append(paths, filepath.SplitList(os.Getenv("DOOMWADDIR"))...)
	paths = append(paths, filepath.SplitList(os.Getenv("DOOMWADPATH"))...)

	if runtime.GOOS != "windows" {
		paths = append(paths,
			"/usr/share/games/doom",
			"/usr/local/share/games/doom")
	}

	return paths
}

// FindIWADs searches the given paths for Doom IWADs and returns a list of found files.
func FindIWADs(paths []string) (map[string]GameVersion, error) {
	iwads := map[string]GameVersion{}
	for _, path := range paths {
		if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if os.IsNotExist(err) {
				// Path doesn't exist, ignore it.
				return nil
			} else if err != nil {
				return err
			}

			if info.IsDir() {
				// Don't recurse.
				return nil
			}

			if strings.ToLower(filepath.Ext(info.Name())) != ".wad" {
				// (Probably) not a WAD file
				return nil
			}

			f, err := os.Open(path)
			if err != nil {
				// Can't open file
				return err
			}
			defer f.Close()

			// Check if likely IWAD
			if magic, err := ReadMagic(f); err != nil {
				return err
			} else if magic != "IWAD" {
				// Not an IWAD
				return nil
			} else {
				// Reset file pointer
				f.Seek(0, 0)
			}

			// Calculate file hash.
			h := md5.New()
			if _, err := io.Copy(h, f); err != nil {
				return err
			}

			sum := hex.EncodeToString(h.Sum(nil))

			if i, ok := IWADs[sum]; ok {
				// A known IWAD!
				iwads[path] = i
			}

			return nil
		}); err != nil {
			return nil, err
		}
	}
	return iwads, nil
}
