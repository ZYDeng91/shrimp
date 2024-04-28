package main

import (
	"testing"
)

func TestPlay(t *testing.T) {
	src, err := NewSource("testdata/good-1-48k.ogg", true)
	if err != nil {
		t.Fatalf(`error creating source: %v`, err)
	}
	d, err := NewDecoder(src)
	if err != nil {
		t.Fatalf(`error creating decoder: %v`, err)
	}
	if d.r.Length() != 137091 {
		t.Fatalf(`wrong sample size: expected 137091, got %d`, d.r.Length())
	}
	if d.r.SampleRate() != 48000 {
		t.Fatalf(`wrong sample rate: expected 48000, got %d`, d.r.SampleRate())
	}
	if d.r.Channels() != 2 {
		t.Fatalf(`wrong channels: expected stereo(2), got %d`, d.r.Channels())
	}
	player, err := NewPlayer(48000, 2, 2048)
	if err != nil {
		t.Fatalf(`error creating player: %v`, err)
	}
	player.Play(d)
	<-player.done
}
