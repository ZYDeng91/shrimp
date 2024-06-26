package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	single := flag.Bool("s", false, "single playback (disable autoplay)")
	isFile := flag.Bool("f", false, "specify input as local file")
	quiet := flag.Bool("q", false, "suppress non-fatal outputs")
	flag.Parse()
	tail := flag.Args()
	if len(tail) == 0 {
		log.Fatal("An argument is required")
	}
	src, err := NewSource(tail[0], *isFile)
	if err != nil {
		log.Fatal(err)
	}

	d, err := NewDecoder(src)
	if err != nil {
		log.Fatal(err)
	}
	if !*quiet {
		fmt.Println("Vendor: ", d.GetVendor())
		fmt.Println("Now Playing: ", d.GetHeader())
	}

	// I had lags for any bufSize above 1000 (1s)
	// also oto won't play until the buffer is full - probably should keep it low/default for now
	bufSize := 0

	player, err := NewPlayer(d.r.SampleRate(), d.r.Channels(), bufSize)
	if err != nil {
		log.Fatal(err)
	}

	for {
		player.Play(d)
		<-player.done
		if *single {
			break
		}
		d.Reset()
		if d.err != nil {
			log.Fatal(d.err)
		}
		if !*quiet {
			// control sequence black magic to update stdout inline
			// cursor up 1 line + clear line + carriage return
			fmt.Print("\033[1A\033[2K\r")
			fmt.Println("Now Playing: ", d.GetHeader())
		}
	}
	log.Print("Done playing, program exited.")
}
