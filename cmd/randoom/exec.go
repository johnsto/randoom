package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	wadio "github.com/johnsto/randoom/pkg/wad"
	"github.com/pkg/errors"
)

type Games []Game

func (gs Games) Includes(game Game) bool {
	for _, g := range gs {
		if g == game {
			return true
		}
	}
	return false
}

// Executables is a list of known Doom executables, with the most preferred ones first.
// The list has been ordered to prefer 'vanilla' experiences over fancier ones.
// This is used by FindExecutable to find an appropriate Doom executable to launch.
var Executables = map[string]Games{
	"chocolate-doom": {Doom, Doom2},
	"chocolate-heretic": {Heretic},
	"chocolate-hexen": {Hexen},
	"doomlegacy": {Doom, Doom2, Heretic},
	"doomsday":   {Doom, Doom2, Heretic, Hexen},
	"gzdoom":     {Doom, Doom2, Heretic, Hexen, Strife},
	//"vavoom": {Doom, Doom2, Heretic, Hexen},
	"zdoom":       {Doom, Doom2, Heretic, Hexen},
	"prboom":      {Doom, Doom2},
	"odamex":      {Doom, Doom2},
	"crispy-doom": {Doom, Doom2},
}

// FindIWADs returns a map of games to their respective IWAD paths. Only one IWAD per
// game is returned, and the IWAD is typically the latest retail release found.
func FindIWADs() (map[Game]string, error) {
	versions, err := wadio.FindIWADs(wadio.GetPaths())
	if err != nil {
		return nil, err
	}

	weights := make(map[Game]int)
	iwads := make(map[Game]string)

	for path, version := range versions {
		game := Game(version.Game)
		if _, ok := iwads[game]; !ok || weights[game] < version.Weight {
			iwads[game] = path
			weights[game] = version.Weight
		}
	}

	return iwads, nil
}

// LaunchGame launches the game.
func LaunchGame(command string, args []string, params ...Param) error {
	fmt.Println("* Launching payload...")

	for _, param := range params {
		args = append(args, param()...)
	}

	if verbose {
		fmt.Printf("> Launching %s %s\n", command, strings.Join(args, " "))
	}

	cmd := exec.Command(command, args...)

	// FIXME: some ports don't play along with PrefixWriter, giving "broken pipe" errors,
	// so just disable it for now.
	//exe := filepath.Base(command)
	cmd.Stdout = os.Stdout //NewPrefixWriter(os.Stdout, []byte{'\n'}, []byte("["+exe+"] "))
	cmd.Stderr = os.Stderr //NewPrefixWriter(os.Stderr, []byte{'\n'}, []byte("["+exe+"] "))

	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "couldn't run game")
	}

	return nil
}

type Param func() []string

// FileParam constructs a -file parameter slice.
func FileParam(name string) Param {
	return func() []string {
		return []string{"-file", name}
	}
}

// WarpLevelParam constructs a -warp parameter slice for the given episode and mission.
func WarpLevelParam(episode, mission int) Param {
	return func() []string {
		if episode > 0 {
			return []string{"-warp", strconv.Itoa(episode), strconv.Itoa(mission)}
		} else {
			return []string{"-warp", strconv.Itoa(mission)}
		}
	}
}

func IWADParam(path string) Param {
	return func() []string {
		return []string{"-iwad", path}
	}
}

// FindExecutable returns the path of the Doom executable to use.
func FindExecutable(game Game) (string, error) {
	for exe, games := range Executables {
		if !games.Includes(game) {
			continue
		}
		path, err := exec.LookPath(exe)
		if err == nil {
			return path, nil
		}
		if err == exec.ErrNotFound {
			continue
		}
	}
	return "", exec.ErrNotFound
}
