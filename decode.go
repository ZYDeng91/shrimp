package main

import (
	"github.com/jfreymuth/oggvorbis"
	"io"
	"strings"
)

type decoder struct {
	r   	*oggvorbis.Reader
	samples [][2]float64
	src 	io.Reader
	err 	error
}

func NewDecoder(src io.Reader) (*decoder, error) {
	r, err := oggvorbis.NewReader(src)
	if err != nil {
		return nil, err
	}
	return &decoder{r, nil, src, nil}, nil
}

func (d *decoder) GetVendor() string {
	return d.r.CommentHeader().Vendor
}

func (d *decoder) GetHeader() string {
	comments := d.r.CommentHeader().Comments
	patterns := []string{"title=", "artist="}
	res := make([]string, len(patterns))
	for _, item := range comments {
		for i, pattern := range patterns {
			if res[i] == "" && item[:len(pattern)] == pattern {
				res[i] = item[len(pattern):]
			}
		}
	}
	for i, item := range res {
		if item == "" {
			res[i] = "unknown"
		}
	}
	return strings.Join(res, " - ")
}

func (d *decoder) ReadDecoded(samples [][2]float64) (n int, ok bool) {
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
		// EOF is passed as well
		if err != nil {
			d.err = err
			break
		}
	}
	return n, ok
}

func (d *decoder) Reset() {
	d.r, d.err = oggvorbis.NewReader(d.src)
}

// implements io.Reader
func (d *decoder) Read(buf []byte) (int, error) {
	// should report an error if buf size is odd
	// hardcoded convert ratio

	// correspondence: len(buf) should be len(samples) * 2(byte depth) * channels
	ratio := 4 // for some reason mono channel plays as usual
	ns := len(buf) / ratio
	if len(d.samples) < ns {
		d.samples = make([][2]float64, ns)
	}

	n, ok := d.ReadDecoded(d.samples)
	if !ok {
		return 0, d.err
	}
	Convert(d.samples, buf)
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
