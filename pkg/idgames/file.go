package idgames

import (
	"encoding/json"
	"strings"
)

// Dir is a directory of files in the idgames repository.
type Dir struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// File is an entry in the idgames repository.
type File struct {
	ID int `json:"id"`

	Author      stringy `json:"author"`
	Date        *Date   `json:"date"`
	Description string  `json:"description"`
	Directory   string  `json:"dir"`
	Email       string  `json:"email"`
	Filename    string  `json:"filename"`
	IdGamesURL  string  `json:"idgamesurl"`
	Rating      float32 `json:"rating"`
	Size        int     `json:"size"`
	Title       stringy `json:"title"`
	URL         string  `json:"url"`
	Votes       int     `json:"votes"`
}

// GetMirrorURL returns a URL for this file, given an appropriate idgames mirror URL.
func (f File) GetMirrorURL(mirror string) string {
	return strings.Replace(f.IdGamesURL, "idgames://", mirror, 1)
}

// stringy contains a value parse from a JSON document that may or may not be a
// string, but we want it to be. This exists as idgames somethings returns
// non-strings in properties that should be strings, such as the 'author' field
// in entry 7775 (curse you, '8130423'!)
type stringy string

func (s *stringy) UnmarshalJSON(b []byte) error {
	var ss string
	if err := json.Unmarshal(b, &ss); err == nil {
		// Hallelujah, it really is a string!
		*s = stringy(ss)
		return nil
	}

	// It's a number, a bool, or a cucumber or something, I don't know, just
	// make it stringy again.
	*s = stringy(b)

	return nil
}
