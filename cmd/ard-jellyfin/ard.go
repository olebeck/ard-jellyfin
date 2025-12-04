package main

import (
	"fmt"
	"time"

	"github.com/olebeck/ard-jellyfin/pkg/ard"
)

func addArd(epg *XmlTvOutput, m3u8 *M3U8Channels) error {

	// dw isnt in the epg program so special case for it
	dwPage, err := ard.GetChannelPage("daserste", "Y3JpZDovL2RldXRzY2hld2VsbGUuZGUvTGl2ZXN0cmVhbS1EZXV0c2NoZVdlbGxl")
	if err != nil {
		return fmt.Errorf("ard.GetChannelPage: %w", err)
	}
	m3u8.AddChannel(M3U8Channel{
		ID:          "deutsche-welle",
		Name:        "Deutsche Welle (DW)",
		Image:       dwPage.Image(),
		PlaylistURL: dwPage.PlaylistURL(),
	})
	epg.AddChannel("deutsche-welle", "Deutsche Welle (DW)", dwPage.Image())

	// fetch channel pages
	program, err := ard.GetProgram(time.Now())
	if err != nil {
		return fmt.Errorf("ard.GetProgram: %w", err)
	}
	var channelPages = make(map[string]*ard.Page)
	for _, ch := range program.Channels {
		page, err := ard.GetChannelPage(ch.ID, ch.Crid)
		if err != nil {
			return fmt.Errorf("ard.GetChannelPage: %w", err)
		}
		channelPages[ch.ID] = page
	}

	now := time.Now()
	for day := -1; day < 7; day++ {
		program, err := ard.GetProgram(now.Add(time.Duration(day) * 24 * time.Hour))
		if err != nil {
			return fmt.Errorf("ard.GetProgram: %w", err)
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

		for _, channel := range program.Channels {
			page := channelPages[channel.ID]
			epg.AddChannel(channel.ID, channel.PublicationService.Name, page.Image())
			for _, slots := range channel.TimeSlots {
				for _, slot := range slots {
					epg.AddProgramme(
						channel.ID,
						slot.Title,
						slot.Synopsis,
						slot.Subline,
						slot.BroadcastedOn,
						slot.BroadcastEnd,
						slot.Image(),
					)
				}
			}
		}
	}

	return nil
}
