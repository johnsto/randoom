package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/johnsto/go-randoom/pkg/idgames"
	wadio "github.com/johnsto/go-randoom/pkg/wad"
	"github.com/pkg/errors"
)

var (
	verbose  bool
	gameFlag mapFlag = make(mapFlag)

	idFlag       = flag.Int("id", 0, "fetch entry with idgames ID")
	pathFlag     = flag.String("path", "", "idgames path to search (e.g. \"levels/doom/g-i/\")")
	keep         = flag.Bool("keep", false, "keep downloaded WAD file")
	mirrorUrl    = flag.String("mirror", idgames.Mirrors[0].URL, "HTTP mirror url")
	apiUrl       = flag.String("api", idgames.DefaultApiUrl, "idgames API url")
	seed         = flag.Int64("seed", time.Now().Unix(), "random seed")
	noIwadFlag   = flag.Bool("noiwad", false, "don't pass IWAD")
	noLaunchFlag = flag.Bool("nolaunch", false, "don't launch game")
)

func init() {
	flag.Var(&gameFlag, "game", "game name [doom|doom2|heretic|hexen] (e.g. -game doom)"+
		" with optional IWAD (e.g. -game doom2=DOOM2.WAD)")
	flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")
}

func main() {
	flag.Parse()

	// Seed RNG
	if seed != nil {
		rand.Seed(*seed)
	}

	fmt.Println("* Establishing uplink...")

	// Fetch file from idgames
	var file *idgames.File
	var game, iwad string
	var iwads map[wadio.Game]string
	var err error

	if *noIwadFlag == false {
		// Find IWADs
		iwads, err = FindIWADs()
		if err != nil && verbose {
			fmt.Println("> IWAD search failed: %s", err)
		} else if verbose {
			fmt.Printf("> selected %d IWADs:\n", len(iwads))
			for game, iwad := range iwads {
				fmt.Printf(">  %s: %s\n", game, iwad)
			}
		}
	}

	if *idFlag != 0 {
		file, err = idgames.Fetch(*idFlag, &idgames.Config{
			ApiUrl: *apiUrl,
		})

		// FIXME: improve error handling for invalid IDs.
		if err != nil {
			panic(errors.Wrapf(err, "couldn't fetch ID %d", *idFlag))
		}

		if gameFlag.Num() > 0 {
			// Determine game based on file path.
			for game, iwad = range gameFlag {
				if strings.Contains(file.Directory, "/"+game+"/") {
					break
				}
			}
		}
	} else {
		path := *pathFlag
		if path == "" {
			// No idgames path given, so pick a path at random, before
			// we pick a random level from the path.
			game, iwad = gameFlag.Random()

			if game == "" && iwads != nil {
				// Pick random game/IWAD from those found.
				for g, w := range iwads {
					game, iwad = string(g), w
					break
				}
			} else if iwad == "" && iwads != nil {
				// Pick an appropriate IWAD for the chosen game
				iwad = iwads[wadio.Game(game)]
			}

			// Still no game? Abandon...
			if game == "" {
				fmt.Println("Error: No game specified! Use the -game flag to specify the " +
					"game to search for e.g. \"-game doom2\".")
				os.Exit(1)
			}

			if verbose {
				fmt.Printf("> picked %s (IWAD: %s)\n", game, iwad)
			}

			path, err = GetLevelsPath(Game(game), Mission)
			if err != nil {
				panic(errors.Wrapf(err, "couldn't get level path"))
			}

			if verbose {
				fmt.Printf("> picked path %s\n", path)
			}
		}

		// Tidy up path for idgames API
		if strings.HasPrefix(path, "/") {
			path = path[1:]
		}

		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}

		// Fetch random entry from archive.
		file, err = GetRandomLevel(path, &idgames.Config{
			ApiUrl: *apiUrl,
		})

		if err != nil {
			panic(errors.Wrapf(err, "couldn't get random level path"))
		}

		if file == nil {
			fmt.Println("Error: failed to find level")
			os.Exit(1)
		}
	}

	if verbose {
		fmt.Printf("> Got file #%d - %s (%d bytes)\n",
			file.ID, file.Filename, file.Size)
		fmt.Printf("> File URL: %s\n", file.URL)
	}

	url := file.GetMirrorURL(*mirrorUrl)
	if verbose {
		fmt.Printf("> Got mirror URL %s\n", *mirrorUrl)
	}

	// Download archive fom mirror.
	fmt.Println("* Retrieving data...")
	archive, err := FetchArchive(url)
	if err != nil {
		panic(errors.Wrapf(err, "couldn't get archive from %s", url))
	}

	// Check that there's a playable WAD in this archive...
	wadName, wad, mapName := archive.GetFirstPlayableWad()
	if wadName == "" {
		panic(errors.Errorf("couldn't find playable WAD in archive"))
	} else if mapName == "" {
		panic(errors.Errorf("couldn't find playable level in WAD"))
	}

	if verbose {
		fmt.Printf("> Found map name %s in %s\n", mapName, wadName)
	}

	// Determine first playable level in the WAD
	mission, episode, err := wadio.SplitMapName(mapName)
	if err != nil {
		panic(errors.Wrapf(err, "couldn't parse map name"))
	}

	// Save wad to a local file
	wadName = SanitizeFilename(wadName)
	if verbose {
		fmt.Printf("> saving WAD to %s\n", wadName)
	}
	if err := SaveWad(wadName, wad); err != nil {
		panic(errors.Wrapf(err, "couldn't save WAD to %s wadName", wadName))
	}

	if *noLaunchFlag == false {
		wadName, err = filepath.Abs(wadName)
		if err != nil && verbose {
			fmt.Printf("> couldn't get absolute WAD path, using relative: %s\n", err)
		}

		// Launch game!
		args := flag.Args()
		params := []Param{
			FileParam(wadName),
			WarpLevelParam(mission, episode),
		}

		if *noIwadFlag == false && iwad != "" {
			// Add IWAD parameter
			params = append(params, IWADParam(iwad))
		}

		var exe string
		var exeFlags []string
		if len(args) == 0 {
			// No executable given, so find one.
			exe, err = FindExecutable(Game(game))
			if err != nil {
				fmt.Println("Couldn't find DOOM executable!", err)
				os.Exit(1)
			}
			if verbose {
				fmt.Printf("> found executable %s\n", exe)
			}
		} else {
			// Use executable and arguments as given
			exe = args[0]
			exeFlags = args[1:]
		}

		err = LaunchGame(exe, exeFlags, params...)

		if err != nil {
			fmt.Printf("Error: game failed to launch (%s)\n", err)
			os.Exit(1)
		}
	}

	if !*keep {
		// Remove downloaded file
		if err := os.Remove(wadName); err != nil {
			panic(errors.Wrapf(err, "couldn't remove level"))
		}
		if verbose {
			fmt.Printf("> Removed WAD file %s\n", wadName)
		}
	}

	fmt.Println()

	PrintGoodbye(file)
}

// PrintGoodbye prints the exit message including WAD name and author.
func PrintGoodbye(f *idgames.File) {
	fmt.Printf("%s (%s) by %s (%s)\n", f.Title, f.Date.Format("02/Jan/2006"), f.Author, f.Email)
	fmt.Printf("ID: #%d\n", f.ID)
	if f.Votes > 1 {
		fmt.Printf("Rating: %.1f/5 by %d users\n", f.Rating, f.Votes)
	} else if f.Votes == 1 {
		fmt.Printf("Rating: %.1f/5 by %d user\n", f.Rating, f.Votes)
	} else {
		fmt.Printf("Rating: unrated - be the first to leave a review!\n")
	}
	fmt.Printf("More info: %s\n", f.URL)
}
