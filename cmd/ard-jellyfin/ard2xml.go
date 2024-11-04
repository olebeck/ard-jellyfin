package main

import (
	"github.com/olebeck/ard-jellyfin/ard"

	"github.com/olebeck/ard-jellyfin/xmltv"
)

func Ard2XmlTV(program *ard.ArdProgram, channelPages map[string]*ard.Page) (*xmltv.TV, error) {
	tv := xmltv.TV{
		SourceInfoName:    "ard-jellyfin",
		GeneratorInfoName: "ard-jellyfin",
	}

	// Create a map to keep track of channels and avoid duplicates
	channelMap := make(map[string]bool)

	for _, ch := range program.Channels {
		page := channelPages[ch.ID]
		// Add channel to XMLTV structure if not already added
		if !channelMap[ch.ID] {
			xmltvChannel := xmltv.Channel{
				ID: ch.ID,
				DisplayNames: []xmltv.DisplayName{
					{Value: ch.PublicationService.Name},
				},
				Icons: []xmltv.Icon{
					{Src: page.Image()},
				},
			}

			tv.Channels = append(tv.Channels, xmltvChannel)
			channelMap[ch.ID] = true
		}

		// Convert each timeslot in the channel to XMLTV Programme
		for _, slots := range ch.TimeSlots {
			for _, slot := range slots {
				xmltvProgramme := xmltv.Programme{
					ChannelID: ch.ID, // Channel ID that this timeslot belongs to
					Titles: []xmltv.Title{
						{Value: slot.Title},
					},
					Descriptions: []xmltv.Description{
						{Value: slot.Synopsis},
					},
					Subtitles: []xmltv.Subtitle{
						{Value: slot.Subline},
					},
					Start: xmltv.XMLTVTime{Time: slot.BroadcastedOn},
					Stop:  xmltv.XMLTVTime{Time: slot.BroadcastEnd},
					Live:  &struct{}{},
				}

				if slot.Images != nil {
					img := slot.Images.Aspect16X9
					xmltvProgramme.Icon = xmltv.Icon{
						Src: img.Src,
					}
				}

				tv.Programmes = append(tv.Programmes, xmltvProgramme)
			}
		}
	}

	return &tv, nil
}
