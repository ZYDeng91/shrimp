package main

import (
	"github.com/hajimehoshi/oto"
	"io"
	"log"
)

type player struct {
	samples   [][2]float64
	buf       []byte
	otoplayer *oto.Player
	done      chan bool
}

func NewPlayer(sampleRate, channels, bufSize int) (*player, error) {
	samples := make([][2]float64, bufSize)
	buf := make([]byte, bufSize*4)
	ctx, err := oto.NewContext(sampleRate, channels, 2, bufSize*4)
	if err != nil {
		return nil, err
	}
	otoplayer := ctx.NewPlayer()
	done := make(chan bool)
	return &player{samples, buf, otoplayer, done}, nil
}

func (p *player) Play(d *decoder) {
	go func() {
		for {
			_, ok := d.Read(p.samples)
			if !ok {
				if d.err == io.EOF {
					p.done <- true
					break
				} else {
					log.Fatal(d.err)
				}
			}
			Convert(p.samples, p.buf)
			p.otoplayer.Write(p.buf)
		}
	}()
}

// convert float to bytes
// buf is updated inplace
func Convert(samples [][2]float64, buf []byte) {
	for i := range samples {
		for c := range samples[i] {
			val := samples[i][c]
			if val < -1 {
				val = -1
			}
			if val > +1 {
				val = +1
			}
			valInt16 := int16(val * (1<<15 - 1))
			low := byte(valInt16)
			high := byte(valInt16 >> 8)
			buf[i*4+c*2+0] = low
			buf[i*4+c*2+1] = high
		}
	}
}
