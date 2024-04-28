package main

import (
	"testing"
)

func TestDecode(t *testing.T) {
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
}

func TestConvert(t *testing.T) {
	samples := [][2]float64{{0.0025265244767069817, 0.005967817734926939}, {0.00829835794866085, 0.008000102825462818}}
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	Convert(samples, buf)
	buf_expected := []byte{82, 0, 195, 0, 15, 1, 6, 1}
	if string(buf) != string(buf_expected) {
		t.Fatalf(`wrong convert result`)
	}
}
