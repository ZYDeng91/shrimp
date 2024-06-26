# Shrimp

[![Go Report Card](https://goreportcard.com/badge/github.com/zydeng91/shrimp)](https://goreportcard.com/report/github.com/zydeng91/shrimp)
[![Build Status](https://github.com/zydeng91/shrimp/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/zydeng91/shrimp/actions/workflows/go.yml?query=branch%3Amaster)

A fussless command line music player.

## Why?

Shrimp started as a weekend project, with the objective to create a bare-bones music player to suit my personal usecase (icecast stream), while adhering to the KISS principle.

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

### Dependencies

Dependencies come from one of the few imports, [Oto](https://github.com/ebitengine/oto#prerequisite).

tl;dr on Linux you'll need ALSA's development library. `libasound2-dev` on Debian based distros.

### Build

`make build` to produce the executable.

### Run

`./shrimp <Stream-URL>` to play a stream over http(s).

`./shrimp -f -s <file-location>` to play a local file, drop the `-s` flag to loop ad infinitum.

### Platforms

Cross-platform functionality is supported by [Oto](https://github.com/ebitengine/oto#platforms)'s drivers.

i.e. Windows, macOS, Linux, BSD, Android, iOS, Wasm, Nintendo Switch, Xbox

### Supporting Codec(s)
[Ogg Vorbis](https://github.com/jfreymuth/oggvorbis)

## Meme
<img alt="not a mascot" src="https://zydeng.com/assets/img/shrimp.png" width=40%>
