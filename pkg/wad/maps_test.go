package wad_test

import . "github.com/johnsto/go-randoom/pkg/wad"
import "testing"

func TestIsMapName(t *testing.T) {
	for name, valid := range map[string]bool{
		"":      false,
		"E":     false,
		"M":     false,
		"E1":    false,
		"E1M":   false,
		"E1M ":   false,
		"E0M1":  true,
		"E1M0":  true,
		"E1M1":  true,
		"E1M2":  true,
		"E4M9":  true,
		"E9M9":  true,
		"MAP": false,
		"MAP9": false,
		"MAP9 ": false,
		"MAP9x": false,
		"MAP00": true,
		"MAP01": true,
		"MAP99": true,
	} {
		if IsMapName(name) != valid {
			t.Errorf("IsMapName(\"%s\"): expected %v, got %v", name, valid, !valid)
		}
	}
}

func TestSplitMapName(t *testing.T) {
	for name, em := range map[string][]int{
		"E0M0": {0, 0},
		"E1M1": {1, 1},
		"E4M9": {4,9},
		"E9M9": {9,9},
		"MAP00": {0, 0},
		"MAP01": {0, 1},
		"MAP10": {0, 10},
		"MAP31": {0, 31},
		"MAP99": {0, 99},
	} {
		if e, m, err := SplitMapName(name); err != nil {
			t.Errorf("SplitMapName(\"%s\"): expected error %s", name, err)
		} else if e != em[0] || m != em[1] {
			t.Errorf("SplitMapName(\"%s\"): expected %d.%d, got %d.%d", name, em[0], em[1], e, m)
		}
	}

	for _, name := range []string{"", " ", "E1M", "MAP", " MAP", "MAPaa"} {
		if _, _, err := SplitMapName(name); err == nil {
			t.Errorf("SplitMapName(\"%s\"): expected error, got nil", name)
		}
	}
}
