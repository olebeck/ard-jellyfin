package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

type Args struct {
	Output string
	Listen string
	Debug  bool
	Timer  int
}

func parseArgs() (*Args, error) {
	var args Args = Args{
		Output: "output",
		Listen: "",
		Debug:  false,
		Timer:  3600,
	}
	if output, ok := os.LookupEnv("OUTPUT_DIR"); ok {
		args.Output = output
	}
	if listen, ok := os.LookupEnv("LISTEN"); ok {
		args.Listen = listen
	}
	if _, ok := os.LookupEnv("DEBUG"); ok {
		args.Debug = true
	}
	if timer, ok := os.LookupEnv("TIMER"); ok {
		timerI, err := strconv.ParseInt(timer, 0, 32)
		if err != nil {
			return nil, err
		}
		args.Timer = int(timerI)
	}
	flag.StringVar(&args.Output, "output", args.Output, "the output folder")
	flag.StringVar(&args.Listen, "listen", args.Listen, "address to listen with an http server on, if you dont want to serve the files using an external server/directly pass the paths to jellyfin")
	flag.BoolVar(&args.Debug, "debug", args.Debug, "enable debug logs")
	flag.IntVar(&args.Timer, "timer", args.Timer, "how long to wait between updates in seconds, 0 if you want it to only run once")
	flag.Parse()
	return &args, nil
}

func main() {
	args, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}
	if args.Debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	absOutput, err := filepath.Abs(args.Output)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output folder: %s\n", absOutput)

	if args.Listen != "" {
		ln, err := net.Listen("tcp", args.Listen)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("http Listening on: %s\n", ln.Addr())
		go http.Serve(ln, http.FileServer(http.FS(os.DirFS(args.Output))))
	}

	os.MkdirAll(args.Output, 0777)

	timerDuration := time.Duration(max(args.Timer, 10)) * time.Second
	t := time.NewTicker(timerDuration)
	for {
		err := run(args.Output)
		if err != nil {
			log.Printf("Error: %s\n", err)
			if args.Timer == 0 {
				os.Exit(1)
			}
		}
		if args.Timer == 0 {
			break
		}
		fmt.Printf("Waiting %s\n", timerDuration)
		<-t.C
	}
}

func run(output string) error {
	fmt.Printf("Running at %s\n", time.Now().Truncate(time.Second))

	epg := NewXmlTvOutput()
	m3u8 := &M3U8Channels{
		channelMap: make(map[string]struct{}),
	}

	if err := addArd(epg, m3u8); err != nil {
		return fmt.Errorf("addArd: %w", err)
	}

	if err := addZdf(epg, m3u8); err != nil {
		return fmt.Errorf("addZdf: %w", err)
	}

	if err := epg.Create(path.Join(output, "ard.xml")); err != nil {
		return fmt.Errorf("epg.Create: %w", err)
	}

	if err := m3u8.Create(path.Join(output, "ard.m3u8")); err != nil {
		return fmt.Errorf("m3u8.Create: %w", err)
	}
	return nil
}
