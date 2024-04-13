package main 

import (
	"os"
	"log"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("An argument is required")
		os.Exit(0)
	}
	url := os.Args[1]
	src, err := URLSource(url)
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
		//fmt.Println("retrying")
		d, err = NewDecoder(src)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Print("Done playing, program exited.")
}
