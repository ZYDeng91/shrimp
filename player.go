package main 

import (
	"io"
	"log"
	"github.com/hajimehoshi/oto"
)

type player struct {
	samples [][2]float64
	buf []byte
	player *oto.Player
	done chan bool
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
			for i := range p.samples {
				for c := range p.samples[i] {
					val := p.samples[i][c]
					if val < -1 {
						val = -1
					}
					if val > +1 {
						val = +1
					}
					valInt16 := int16(val * (1<<15 - 1))
					low := byte(valInt16)
					high := byte(valInt16 >> 8)
					p.buf[i*4+c*2+0] = low
					p.buf[i*4+c*2+1] = high
				}
			}
			p.player.Write(p.buf)
		}
	}()
}
