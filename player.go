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
	p.otoplayer = p.otoctx.NewPlayer(d)
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
