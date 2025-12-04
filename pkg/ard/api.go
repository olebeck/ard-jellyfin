package ard

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

var client http.Client

func GetProgram(day time.Time) (*ArdProgram, error) {
	slog.Debug("GetProgram", "day", day.Format(time.DateOnly))
	res, err := client.Get("https://programm-api.ard.de/program/api/program?day=" + day.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	var program ArdProgram
	if err = d.Decode(&program); err != nil {
		return nil, err
	}
	return &program, nil
}

type ChannelInfo struct {
	ID   string
	Name string
}

func GetChannels() ([]ChannelInfo, error) {
	slog.Debug("GetChannels")
	res, err := client.Get("https://programm-api.ard.de/nownext/api/now")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	var body map[string]any
	err = d.Decode(&body)
	if err != nil {
		return nil, err
	}
	var channels []ChannelInfo
	for _, channel := range body {
		channel, ok := channel.(map[string]any)
		if !ok {
			continue
		}
		id, _ := channel["id"].(string)
		name, _ := channel["name"].(string)
		channels = append(channels, ChannelInfo{
			ID:   id,
			Name: name,
		})
	}
	return channels, nil
}

func GetChannelPage(channelId, Crid string) (*Page, error) {
	slog.Debug("GetChannelPage", "channel", channelId)
	res, err := client.Get(fmt.Sprintf("https://api.ardmediathek.de/page-gateway/pages/%s/item/%s?embedded=true&mcV6=true", channelId, Crid))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	var page Page
	if err = d.Decode(&page); err != nil {
		return nil, err
	}
	return &page, nil
}
