package main

import (
	"encoding/xml"
	"os"
	"time"

	"github.com/olebeck/ard-jellyfin/xmltv"
)

type XmlTvOutput struct {
	tv         xmltv.TV
	channelMap map[string]struct{}
}

func (x *XmlTvOutput) Create(outputPath string) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	e := xml.NewEncoder(f)
	e.Indent("", "  ")
	return e.Encode(x.tv)
}

func (x *XmlTvOutput) AddChannel(id, name, image string) bool {
	if _, ok := x.channelMap[id]; ok {
		return false
	}
	x.tv.Channels = append(x.tv.Channels, xmltv.Channel{
		ID:           id,
		DisplayNames: []xmltv.DisplayName{{Value: name}},
		Icons:        []xmltv.Icon{{Src: image}},
	})
	x.channelMap[id] = struct{}{}
	return true
}

func (x *XmlTvOutput) AddProgramme(channelID, title, desc, subline string, start, end time.Time, icon string) {
	x.tv.Programmes = append(x.tv.Programmes, xmltv.Programme{
		ChannelID:    channelID,
		Titles:       []xmltv.Title{{Value: title}},
		Descriptions: []xmltv.Description{{Value: desc}},
		Subtitles:    []xmltv.Subtitle{{Value: subline}},
		Start:        xmltv.XMLTVTime{Time: start},
		Stop:         xmltv.XMLTVTime{Time: end},
		Live:         &struct{}{},
		Icon:         xmltv.Icon{Src: icon},
	})
}
