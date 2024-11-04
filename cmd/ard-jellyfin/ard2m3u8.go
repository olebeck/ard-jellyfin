package main

import (
	"github.com/olebeck/ard-jellyfin/ard"
)

func Ard2M3U8(program *ard.ArdProgram, channelPages map[string]*ard.Page) (*M3U8Channels, error) {
	channels := M3U8Channels{}
	for _, ch := range program.Channels {
		page := channelPages[ch.ID]
		channels.Channels = append(channels.Channels, M3U8Channel{
			ID:          ch.ID,
			Name:        ch.PublicationService.Name,
			Image:       page.Image(),
			PlaylistURL: page.PlaylistURL(),
		})
	}
	return &channels, nil
}
