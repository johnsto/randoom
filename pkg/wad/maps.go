package wad

import (
	"strconv"

	"github.com/pkg/errors"
)

var (
	ErrInvalidMapName = errors.New("invalid map name")
)

// IsMapName returns true if the given lump name is a valid map name.
func IsMapName(name string) bool {
	if len(name) == 4 &&
		name[0] == 'E' && name[2] == 'M' &&
		name[1] >= '0' && name[1] <= '9' &&
		name[3] >= '0' && name[3] <= '9' {
		return true
	} else if len(name) == 5 &&
		name[0:3] == "MAP" &&
		name[3] >= '0' && name[3] <= '9' &&
		name[4] >= '0' && name[4] <= '9' {
		return true
	}
	return false
}

// SplitMapName splits the given map name (such as E1M2 or MAP03) into it's
// parts (such as [1, 2], or [0, 3] respectively).
func SplitMapName(name string) (int, int, error) {
	if len(name) == 4 && name[0] == 'E' && name[2] == 'M' {
		// ExMy
		episode, err := strconv.Atoi(name[1:2])
		if err != nil {
			return 0, 0, errors.Wrapf(ErrInvalidMapName, "invalid episode index %s", name[1:2])
		}
		mission, err := strconv.Atoi(name[3:4])
		if err != nil {
			return 0, 0, errors.Wrapf(ErrInvalidMapName, "invalid mission index %s", name[3:4])
		}
		return episode, mission, nil
	} else if len(name) == 5 && name[0:3] == "MAP" {
		// MAPzz
		mission, err := strconv.Atoi(name[3:5])
		if err != nil {
			return 0, 0, errors.Wrapf(ErrInvalidMapName, "invalid map index %s", name[3:])
		}
		return 0, mission, nil
	} else {
		return 0, 0, errors.Wrapf(ErrInvalidMapName, "unknown map name %s", name)
	}
}
