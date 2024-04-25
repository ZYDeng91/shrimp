# Shrimp

Shrimp is a fuss-less, feature-less command line internet radio music player.

## Why?

The whole thing is put together within 2 days as a weekend project, with the objective to create a minimalist music player to suit my usecase (i.e. streams).

### Imaginary QnA

> This is missing so many features

For now, Shrimp is intented to play radio streams, many conventional music player features do not apply (e.g. playlists). Does your FM radio have fancy features like that?

> mpv is better

Not in terms of resource usage. Playing audio with a full-fledged video player is an overkill. That being said, the difference is negligible in practice, especially with modern hardware.

> How do I find radio stations, is there a curated list?

We don't shill stations here, not even my own. If you are lost, try http://dir.xiph.org/codecs/Vorbis.

> Can I save my favorite stations

This can be and should be handled by users. Just feed me the URL.

> Why shrimp, if not cRustacean

It's the best euphemism I can think of when you pin together "Simple Internet Music Player". I like Rust, though.

## Usage

`make build` to produce the executable.

`./shrimp <Stream-URL>`

Simply pass your streaming endpoint URL as an argument.

`./shrimp -f -s <file-location>` to play a local file, remove the `-s` flag to loop ad infinitum.

[Oto](https://github.com/ebitengine/oto)'s drivers should provide cross-platform support.

## Supporting Codec(s)
Ogg Vorbis (https://github.com/jfreymuth/oggvorbis)

## TODO:
- ~~Display metadata when playing~~
- Take keyboard input when playing (for basic commands like pause and quit)
- Add Opus codec support
- ~~Add option to play local files~~
- Investigate the mystery of go's compiled binary size
it sux to be go i guess, disabling debug symbols is the most I can do 
