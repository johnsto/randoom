package main

import (
	"math/rand"
	"strings"
)

// mapFlag parses multiple arguments of the form <key>=<value> under
// the same flag as a map.
type mapFlag map[string]string

func (f *mapFlag) String() string {
	var s string
	for k, v := range *f {
		if s != "" {
			s = s + " "
		}
		s = s + k
		if v != "" {
			s = s + "=" + v
		}
	}
	return s
}

func (f *mapFlag) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) == 1 {
		(*f)[parts[0]] = ""
	} else {
		(*f)[parts[0]] = parts[1]
	}
	return nil
}

func (f *mapFlag) Num() int {
	return len(*f)
}

func (f *mapFlag) First() (string, string) {
	for k, v := range *f {
		return k, v
	}
	return "", ""
}

func (f *mapFlag) Random() (string, string) {
	if len(*f) > 0 {
		n := rand.Intn(len(*f))
		for k, v := range *f {
			if n == 0 {
				return k, v
			}
			n--
		}
	}
	return "", ""
}
