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

func (d *decoder) GetVendor() string {
	return d.r.CommentHeader().Vendor
}

func (d *decoder) GetHeader() string {
	comments := d.r.CommentHeader().Comments
	res := make([]byte, 0)
	ok := false
	for _, item := range(comments) {
		if item[:6] == "title=" {
			res = append(res, item[6:]...)
			ok = true
			break
		}
	}

	if !ok {
		res = append(res, "unknown"...)
	}

	res = append(res, " - "...)

	ok = false
	for _, item := range(comments) {
		if item[:7] == "artist=" {
			res = append(res, item[7:]...)
			ok = true
			break
		}
	}
	if !ok {
		res = append(res, "unknown"...)
	}

	return string(res)
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
