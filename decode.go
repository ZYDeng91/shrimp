package main 

import (
	"io"
	"github.com/pkg/errors"
	"github.com/jfreymuth/oggvorbis"
)

type decoder struct {
	r *oggvorbis.Reader
	src io.Reader
	err error
}

func NewDecoder(src io.Reader) (*decoder, error){
	r, err := oggvorbis.NewReader(src)
	if err != nil {
		return nil, err
	}
	return &decoder{r, src, nil}, nil
}

func (d *decoder) Read(samples [][2]float64) (n int, ok bool) {
	if d.err != nil {
		return 0, false
	}
	var tmp [2]float32
	for i := range samples {
		dn, err := d.r.Read(tmp[:])
		if dn == 2 {
			samples[i][0], samples[i][1] = float64(tmp[0]), float64(tmp[1])
			n++
			ok = true
		}
		if err == io.EOF {
			d.err = err
			break
		}
		if err != nil {
			d.err = errors.Wrap(err, "ogg/vorbis")
			break
		}
	}
	return n, ok
}
