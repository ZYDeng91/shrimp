package main

import (
	"io"
	"fmt"
	"net/http"
)

func URLSource(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err	
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid HTTP response: %s", res.Status)
	}
	mime := res.Header["Content-Type"][0]
	if mime != "application/ogg" && mime != "audio/ogg" {
		return nil, fmt.Errorf("Wrong MIME type (%s), are you sure this is an ogg/vorbis stream?", mime)
	}

	return res.Body, nil
}
