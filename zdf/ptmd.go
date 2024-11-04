package zdf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Ptmd struct {
	Attributes struct {
		DvrDurationDash struct {
			Profile string `json:"profile"`
			Value   string `json:"value"`
		} `json:"dvrDurationDash"`
		DvrDurationHls struct {
			Profile string `json:"profile"`
			Value   string `json:"value"`
		} `json:"dvrDurationHls"`
		Encryption struct {
			Profile string `json:"profile"`
			Value   string `json:"value"`
		} `json:"encryption"`
		Fsk struct {
			Profile string `json:"profile"`
			Value   string `json:"value"`
		} `json:"fsk"`
		GeoLocation struct {
			Profile string `json:"profile"`
			Value   string `json:"value"`
		} `json:"geoLocation"`
		Profile string `json:"profile"`
		Start   struct {
			Profile string    `json:"profile"`
			Value   time.Time `json:"value"`
		} `json:"start"`
		State struct {
			Profile string `json:"profile"`
			Value   string `json:"value"`
		} `json:"state"`
		Station struct {
			Profile string `json:"profile"`
			Value   string `json:"value"`
		} `json:"station"`
		Stop struct {
			Profile string    `json:"profile"`
			Value   time.Time `json:"value"`
		} `json:"stop"`
		StreamSymbol struct {
			Profile string `json:"profile"`
			Value   string `json:"value"`
		} `json:"streamSymbol"`
		Type struct {
			Profile string `json:"profile"`
			Value   string `json:"value"`
		} `json:"type"`
	} `json:"attributes"`
	Captions        []any  `json:"captions"`
	Description     string `json:"description"`
	DocumentVersion int    `json:"documentVersion"`
	ID              string `json:"id"`
	Mandant         string `json:"mandant"`
	PlayerID        string `json:"playerId"`
	PriorityList    []struct {
		Formitaeten []struct {
			Facets     []string `json:"facets"`
			IsAdaptive bool     `json:"isAdaptive"`
			MimeType   string   `json:"mimeType"`
			Profile    string   `json:"profile"`
			Qualities  []struct {
				Audio struct {
					Profile string `json:"profile"`
					Tracks  []struct {
						Cdn      string `json:"cdn"`
						Class    string `json:"class"`
						Language string `json:"language"`
						Profile  string `json:"profile"`
						URI      string `json:"uri"`
					} `json:"tracks"`
				} `json:"audio"`
				MimeCodec string `json:"mimeCodec"`
				Profile   string `json:"profile"`
				Quality   string `json:"quality"`
			} `json:"qualities"`
			Type string `json:"type"`
		} `json:"formitaeten"`
		Profile string `json:"profile"`
	} `json:"priorityList"`
	Profile string `json:"profile"`
	Self    string `json:"self"`
	Version string `json:"version"`
}

func getPtmd(token, id string) (*Ptmd, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.zdf.de/tmd/2/ngplayer_2_5/live/ptmd/%s", id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Api-Auth", "Bearer "+token)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	var ptmd Ptmd
	err = d.Decode(&ptmd)
	if err != nil {
		return nil, err
	}
	return &ptmd, nil
}
