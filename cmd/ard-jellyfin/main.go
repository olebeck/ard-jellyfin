package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/olebeck/ard-jellyfin/xmltv"
)

func main() {
	output := flag.String("output", "output", "the output folder")
	listen := flag.String("listen", "", "address to listen with an http server on, if you dont want to serve the files using an external server/directly pass the paths to jellyfin")
	timerSeconds := flag.Int("timer", 3600, "how long to wait between updates in seconds, 0 if you want it to only run once")
	flag.Parse()

	fmt.Printf("output folder: %s\n", *output)

	if *listen != "" {
		ln, err := net.Listen("tcp", *listen)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("http Listening on: %s\n", ln.Addr())
		go http.Serve(ln, http.FileServer(http.FS(os.DirFS(*output))))
	}

	os.Mkdir(*output, 0777)

	timerValue := *timerSeconds
	timerDuration := time.Duration(max(timerValue, 1)) * time.Second
	t := time.NewTicker(timerDuration)
	for {
		err := run(*output)
		if err != nil {
			log.Printf("Error: %s\n", err)
			if timerValue == 0 {
				os.Exit(1)
			}
		}
		if timerValue == 0 {
			break
		}
		fmt.Printf("Waiting %s\n", timerDuration)
		<-t.C
	}
}

func run(output string) error {
	fmt.Printf("Running at %s\n", time.Now().Truncate(time.Second))

	tvOut := &XmlTvOutput{
		tv: xmltv.TV{
			SourceInfoName:    "ard-jellyfin",
			GeneratorInfoName: "ard-jellyfin",
		},
		channelMap: make(map[string]struct{}),
	}
	m3u8 := &M3U8Channels{
		channelMap: make(map[string]struct{}),
	}

	if err := addArd(tvOut, m3u8); err != nil {
		return fmt.Errorf("addArd: %w", err)
	}

	if err := addZdf(tvOut, m3u8); err != nil {
		return fmt.Errorf("addZdf: %w", err)
	}

	if err := tvOut.Create(path.Join(output, "ard.xml")); err != nil {
		return fmt.Errorf("tvOut.Create: %w", err)
	}

	if err := m3u8.Create(path.Join(output, "ard.m3u8")); err != nil {
		return fmt.Errorf("m3u8.Create: %w", err)
	}
	return nil
}
