package main

import (
	"github.com/ebitengine/oto/v3"
	"time"
)

type player struct {
	otoctx    *oto.Context
	otoplayer *oto.Player
	done      chan bool
}

// a decoder wrapper to implement io.Reader
type converter struct {
	d       *decoder
	samples [][2]float64
}

func NewPlayer(sampleRate, channels, bufSize int) (*player, error) {
	ctx, ready, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   sampleRate,
		ChannelCount: channels,
		Format:       oto.FormatSignedInt16LE,
		BufferSize:   time.Duration(bufSize) * time.Millisecond,
	})
	if err != nil {
		return nil, err
	}
	<-ready
	done := make(chan bool)
	return &player{ctx, nil, done}, nil
}

func (p *player) Play(d *decoder) {
	c := NewConverter(d)
	p.otoplayer = p.otoctx.NewPlayer(c)
	p.otoplayer.Play()

	go func() {
		for p.otoplayer.IsPlaying() {
			time.Sleep(100 * time.Millisecond)
		}
		// IsPlaying does not wait for hardware to play buffer
		// sleep for a BufferSize to compensate?

		// quick fix: close resource to prevent audio overlap
		p.otoplayer.Close()
		p.done <- true
	}()
}

func NewConverter(d *decoder) *converter {
	return &converter{d, nil}
}

func (c *converter) Read(buf []byte) (int, error) {
	// should report an error if buf size is odd
	// hardcoded convert ratio

	// correspondence: len(buf) should be len(samples) * 2(bit depth) * channels
	ratio := 4 // for some reason mono channel plays as usual
	ns := len(buf) / ratio
	if len(c.samples) < ns {
		c.samples = make([][2]float64, ns)
	}

	n, ok := c.d.Read(c.samples)
	if !ok {
		return 0, c.d.err
	}
	Convert(c.samples, buf)
	return ratio * n, nil
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
