package main

import (
	"io"
	"os"
	"fmt"
	"net/http"
)

func NewSource(url string, isFile bool) (io.Reader, error) {
	if isFile {
		return fileSource(url)
	} else {
		return URLSource(url)
	}
}

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

func fileSource(file string) (io.Reader, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return f, nil
}
