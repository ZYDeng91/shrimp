# Shrimp

[![Go Report Card](https://goreportcard.com/badge/github.com/zydeng91/shrimp)](https://goreportcard.com/report/github.com/zydeng91/shrimp)
[![Build Status](https://github.com/zydeng91/shrimp/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/zydeng91/shrimp/actions/workflows/go.yml?query=branch%3Amaster)

A fussless command line music player.

## Why?

Shrimp started as a weekend project, with the objective to create a bare-bones music player to suit my usecase (i.e. streams), while adhering to the KISS principle.

I wouldn't consider myself revinventing the wheel if most wheels are heavy and bloated.

### Imaginary QnA

> I need my features

No you don't. Even if you do, they should be handled by external programs (eg. EQ).

Shrimp is mainly intended to play radio streams, many conventional music player features would not apply. Does your FM radio have those fancy features?

> mpv is better

~~Not in terms of resource usage. Playing audio with a full-fledged video player is an overkill. That being said, the difference is negligible in practice, especially with modern hardware.~~

True.

> How do I find radio stations, is there a curated list?

We don't shill stations here, not even my own. If you are lost, try http://dir.xiph.org/codecs/Vorbis.

> Why shrimp, if not cRustacean

It's the best euphemism I can think of when you pin together "Simple Internet Music Player". I like Rust, though.

## Usage

`make build` to produce the executable.

`./shrimp <Stream-URL>`

Simply pass your streaming endpoint URL as an argument.

`./shrimp -f -s <file-location>` to play a local file, remove the `-s` flag to loop ad infinitum.

Cross-platform functionality should be supported by [Oto](https://github.com/ebitengine/oto/tree/v1.0.1).

## Supporting Codec(s)
[Ogg Vorbis](https://github.com/jfreymuth/oggvorbis)

## TODO:
- ~~Display metadata when playing~~
- Take keyboard input when playing (for basic commands like pause and quit)

Most solutions require importing extra libs, we'll see if it's worth. (no interaction = no seek)

- Add Opus codec support
- ~~Add option to play local files~~
- Investigate the mystery of go's compiled binary size

It sux to be go I guess, disabling debug symbols is the most I can do

- Add test cases & mess with github actions
- Add playlist support (will require resampling)

### Meme
![Not a mascot](https://zydeng.com/assets/img/shrimp.png)
