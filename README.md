# RanDOOM

*Rekindling the days of weekends spent downloading random Doom WADs from BBS's in the 1990's*

## Usage

Install RanDOOM using `go get`, e.g.

    go get github.com/johnsto/randoom/cmd/randoom

You can then launch it with `randoom`.

## FAQ

### What does it actually do?

RanDOOM is really just a downloader and launcher for WAD files found in the [idgames repository](https://legacy.doomworld.com/idgames/), with a distinct focus on giving a endless\* supply of DOOMage. It does this by finding a downloading a random WAD file each time it's launched.

\* not really endless.

### Does it generate levels?

No, it only fetches existing user-made levels. If you want a random level generator, see [ObHack](http://www.samiam.org/ObHack/).

### What games does it support?

Doom, Doom 2, Heretic, Hexen and Strife, although I've only actually tested it on the first two. On launch, RanDOOM will search common paths for IWAD files for each of these games, and launch the appropriate game (again, randomly!)

### What source ports does it support?

It has built-in support for Chocolate Doom, Doom Legacy, Doomsday, GZDoom, ZDoom, PrBoom, Odamex and Crispy Doom - if you have any of these installed, it will automatically launch the most appropriate one (provided it can be found.)

If you're using another port, or want more control, you can instead pass your chosen executable (and any custom arguments), e.g.:

    randoom -game doom vavoom -nomusic -fast -nostartup

### How do I set the skill?

You can pass your chosen skill level (and executable) via the command line, e.g.

    randoom -game doom2 doomsday -skill 4
    
### What does "Couldn't find DOOM executable! executable file not found in %PATH%" mean?

That means RanDOOM couldn't find the Doom executable! On Linux, this means it's not in your path - try running `which doom` or `which doom2` to verify, and then fix accordingly.

On Windows, things are trickier. I suggest installing [Chocolate Doom](https://www.chocolate-doom.org) as it gives you the vanilla Doom experience and automatically loads in IWADs from Steam. Try extracting it into (e.g.) `c:\ChocolateDoom`. Then you can run randoom like this:

    randoom -game doom c:\ChocolateDoom\chocolate-doom.exe -nomusic -fast -nostartup

## TODO

- [x] Test on Linux.
- [x] Test on Windows.
- [ ] Test on MacOS.
- [ ] Support DeHacked files.
- [ ] Support multi-PWAD archives.
- [ ] Support finding IWADs in local Steam installs.
- [ ] Support FreeDOOM IWAD files.
- [ ] Tidy up command line parsing.
- [ ] Improve release packaging.
