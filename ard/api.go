package ard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client http.Client

func GetProgram(day time.Time) (*ArdProgram, error) {
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

func GetChannelPage(ch *Channel) (*Page, error) {
	res, err := client.Get(fmt.Sprintf("https://api.ardmediathek.de/page-gateway/pages/%s/item/%s?embedded=true&mcV6=true", ch.ID, ch.Crid))
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
