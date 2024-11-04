package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/olebeck/ard-jellyfin/ard"
)

func writeXmlTv(output string, program *ard.ArdProgram, channelPages map[string]*ard.Page) error {
	tv, err := Ard2XmlTV(program, channelPages)
	if err != nil {
		return err
	}
	f, err := os.Create(path.Join(output, "ard.xml"))
	if err != nil {
		return err
	}
	defer f.Close()
	e := xml.NewEncoder(f)
	e.Indent("", "  ")
	return e.Encode(tv)
}

func writeM3U8(output string, program *ard.ArdProgram, channelPages map[string]*ard.Page) error {
	channels, err := Ard2M3U8(program, channelPages)
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(output, "ard.m3u8"))
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = channels.WriteTo(f); err != nil {
		return err
	}
	return nil
}

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
	program, err := ard.GetProgram(time.Now())
	if err != nil {
		return fmt.Errorf("ard.GetProgram: %w", err)
	}

	var channelPages = make(map[string]*ard.Page)
	for _, ch := range program.Channels {
		page, err := ard.GetChannelPage(ch)
		if err != nil {
			return fmt.Errorf("ard.GetChannelPage: %w", err)
		}
		channelPages[ch.ID] = page
	}

	if err = writeXmlTv(output, program, channelPages); err != nil {
		return fmt.Errorf("writeXmlTv: %w", err)
	}

	if err := writeM3U8(output, program, channelPages); err != nil {
		return fmt.Errorf("writeM3U8: %w", err)
	}
	return nil
}
