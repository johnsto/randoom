# RanDOOM

*Rekindling the days of weekends spent downloading random Doom WADs from BBS's in the 1990's*

## Usage

Install RanDOOM using `go get`, e.g.

    go get github.com/johnsto/randoom/pkg/randoom

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

    randoom -game doom vavoom -skill 4

### How do I set the skill?

You can pass your chosen skill level (and executable) via the command line, e.g.

    randoom -game doom2 doomsday -skill 2

## TODO

- [x] Test on Linux.
- [ ] Test on Windows.
- [ ] Test on MacOS.
- [ ] Support DeHacked files.
- [ ] Support multi-PWAD archives.
- [ ] Tidy up command line parsing.
- [ ] Improve release packaging.
