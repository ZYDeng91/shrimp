package main 

import (
	"log"
	"flag"
)

func main() {
	single := flag.Bool("s", false, "single playback (disable autoplay)")
	isFile := flag.Bool("f", false, "specify input as local file")
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

	bufSize := 2048

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
	}

	log.Print("Done playing, program exited.")
}
