package main

import (
	"fmt"
	"io"
	"os"
)

type M3U8Channel struct {
	ID          string
	Name        string
	Image       string
	PlaylistURL string
}

func (m *M3U8Channel) WriteTo(w io.Writer) (int64, error) {
	fmt.Fprintf(w, `#EXTINF:-1,tvg-id="%s",tvg-name="%s",tvg-logo="%s"`, m.ID, m.Name, m.Image)
	w.Write([]byte{'\n'})
	io.WriteString(w, m.PlaylistURL)
	w.Write([]byte{'\n', '\n'})
	return 0, nil
}

type M3U8Channels struct {
	Channels   []M3U8Channel
	channelMap map[string]struct{}
}

func (m *M3U8Channels) HaveChannel(id string) bool {
	_, ok := m.channelMap[id]
	return ok
}

func (m *M3U8Channels) AddChannel(channel M3U8Channel) {
	if _, ok := m.channelMap[channel.ID]; ok {
		return
	}
	m.channelMap[channel.ID] = struct{}{}
	m.Channels = append(m.Channels, channel)
}

func (m *M3U8Channels) WriteTo(w io.Writer) (int64, error) {
	io.WriteString(w, "#EXTM3U\n")
	io.WriteString(w, "#EXT-X-VERSION:3\n")
	for _, channel := range m.Channels {
		channel.WriteTo(w)
	}
	return 0, nil
}

func (m *M3U8Channels) Create(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = m.WriteTo(f); err != nil {
		return err
	}
	return nil
}
