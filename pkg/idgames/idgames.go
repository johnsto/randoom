package idgames

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// FetchDirs returns a list of directories in the given path.
func FetchDirs(name string, cfg *Config) ([]Dir, error) {
	params := make(url.Values)
	params.Set("action", "getdirs")
	params.Set("name", name)

	dirs := make([]Dir, 0)
	return dirs, fetch(params, cfg, &dirs)
}

// FetchFiles returns a list of files in the given path.
func FetchFiles(path string, cfg *Config) ([]File, error) {
	params := make(url.Values)
	params.Set("action", "getfiles")
	params.Set("name", path)

	tmp := struct {
		Files []File `json:"file"`
	}{}
	return tmp.Files, fetch(params, cfg, &tmp)
}

// Fetch returns a file entry for the given file ID.
func Fetch(id int, cfg *Config) (*File, error) {
	params := make(url.Values)
	params.Set("action", "get")
	params.Set("id", strconv.Itoa(id))

	file := new(File)
	return file, fetch(params, cfg, &file)
}

func fetch(params url.Values, cfg *Config, v interface{}) error {
	params.Set("out", "json")

	// FIXME: apiUrl *might* contain query part already
	resp, err := http.Get(cfg.apiUrl() + "?" + params.Encode())
	if err != nil {
		return err
	}

	tmp := struct {
		Meta struct {
			Version int `json:"version"`
		} `json:"meta"`
		Content json.RawMessage `json:"content"`
	}{}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&tmp); err != nil {
		return errors.Wrap(err, "JSON response malformed")
	}

	if tmp.Meta.Version != 3 {
		fmt.Printf("Warning: unexpected version %d", tmp.Meta.Version)
	}

	return json.Unmarshal(tmp.Content, v)
}
