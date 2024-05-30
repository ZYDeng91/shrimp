package main

import (
	"github.com/jfreymuth/oggvorbis"
	"io"
	"strings"
)

type decoder struct {
	r   *oggvorbis.Reader
	src io.Reader
	err error
}

func NewDecoder(src io.Reader) (*decoder, error) {
	r, err := oggvorbis.NewReader(src)
	if err != nil {
		return nil, err
	}
	return &decoder{r, src, nil}, nil
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
