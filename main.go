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
		fmt.Print("Now Playing: ", d.GetHeader())
	}
	bufSize := 2048

	// for playlists, files may have different sample rate & channels
	// init a new player every time is undesirable
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
		//fmt.Println("retrying")
		d, err = NewDecoder(src)
		if err != nil {
			log.Fatal(err)
		}
		if !*quiet {
			// control sequence black magic to update stdout inline
			fmt.Print("\033[2K\r")
			fmt.Print("Now Playing: ", d.GetHeader())
		}
	}
	fmt.Println()
	log.Print("Done playing, program exited.")
}
