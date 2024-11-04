package main

import (
	"fmt"
	"time"

	"github.com/olebeck/ard-jellyfin/ard"
)

func addArd(out *XmlTvOutput, m3u8 *M3U8Channels) error {
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

	for _, ch := range program.Channels {
		page := channelPages[ch.ID]
		m3u8.AddChannel(M3U8Channel{
			ID:          ch.ID,
			Name:        ch.PublicationService.Name,
			Image:       page.Image(),
			PlaylistURL: page.PlaylistURL(),
		})
	}

	for _, ch := range program.Channels {
		page := channelPages[ch.ID]
		if out.AddChannel(ch.ID, ch.PublicationService.Name, page.Image()) {
			for _, slots := range ch.TimeSlots {
				for _, slot := range slots {
					out.AddProgramme(ch.ID, slot.Title, slot.Synopsis, slot.Subline, slot.BroadcastedOn, slot.BroadcastEnd, slot.Image())
				}
			}
		}

	}

	return nil
}
