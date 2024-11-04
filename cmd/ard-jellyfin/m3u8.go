package main

import (
	"fmt"
	"io"
)

type M3U8Channel struct {
	ID          string
	Name        string
	Image       string
	PlaylistURL string
}

func (m *M3U8Channel) WriteTo(w io.Writer) (int64, error) {
	nn := 0
	n, _ := fmt.Fprintf(w, "#EXTINF:-1,")
	nn += n
	n, _ = fmt.Fprintf(w, "tvg-id=\"%s\",", m.ID)
	nn += n
	n, _ = fmt.Fprintf(w, "tvg-epgid=\"%s\",", m.ID)
	nn += n
	n, _ = fmt.Fprintf(w, "tvg-name=\"%s\",", m.Name)
	nn += n
	n, _ = fmt.Fprintf(w, "tvg-logo=\"%s\"", m.Image)
	nn += n
	n, _ = fmt.Fprintf(w, "\n%s\n", m.PlaylistURL)
	nn += n
	return int64(nn), nil
}

type M3U8Channels struct {
	Channels []M3U8Channel
}

func (m *M3U8Channels) WriteTo(w io.Writer) (int64, error) {
	nn := 0
	n, _ := io.WriteString(w, "#EXTM3U\n")
	nn += n
	n, _ = io.WriteString(w, "#EXT-X-VERSION:3\n")
	nn += n
	for _, channel := range m.Channels {
		n, _ := channel.WriteTo(w)
		nn += int(n)
	}
	return int64(nn), nil
}
