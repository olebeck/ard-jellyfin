package zdf

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Ptmd struct {
	Attributes      Attributes     `json:"attributes"`
	Captions        []any          `json:"captions"`
	Description     string         `json:"description"`
	DocumentVersion int            `json:"documentVersion"`
	ID              string         `json:"id"`
	Mandant         string         `json:"mandant"`
	PlayerID        string         `json:"playerId"`
	PriorityList    []PriorityList `json:"priorityList"`
	Profile         string         `json:"profile"`
	Self            string         `json:"self"`
	Version         string         `json:"version"`
}
type DvrDurationDash struct {
	Profile string `json:"profile"`
	Value   string `json:"value"`
}
type DvrDurationHls struct {
	Profile string `json:"profile"`
	Value   string `json:"value"`
}
type Encryption struct {
	Profile string `json:"profile"`
	Value   string `json:"value"`
}
type Fsk struct {
	Profile string `json:"profile"`
	Value   string `json:"value"`
}
type GeoLocation struct {
	Profile string `json:"profile"`
	Value   string `json:"value"`
}
type Start struct {
	Profile string    `json:"profile"`
	Value   time.Time `json:"value"`
}
type State struct {
	Profile string `json:"profile"`
	Value   string `json:"value"`
}
type Station struct {
	Profile string `json:"profile"`
	Value   string `json:"value"`
}
type Stop struct {
	Profile string    `json:"profile"`
	Value   time.Time `json:"value"`
}
type StreamSymbol struct {
	Profile string `json:"profile"`
	Value   string `json:"value"`
}
type Type struct {
	Profile string `json:"profile"`
	Value   string `json:"value"`
}
type Attributes struct {
	DvrDurationDash DvrDurationDash `json:"dvrDurationDash"`
	DvrDurationHls  DvrDurationHls  `json:"dvrDurationHls"`
	Encryption      Encryption      `json:"encryption"`
	Fsk             Fsk             `json:"fsk"`
	GeoLocation     GeoLocation     `json:"geoLocation"`
	Profile         string          `json:"profile"`
	Start           Start           `json:"start"`
	State           State           `json:"state"`
	Station         Station         `json:"station"`
	Stop            Stop            `json:"stop"`
	StreamSymbol    StreamSymbol    `json:"streamSymbol"`
	Type            Type            `json:"type"`
}
type Tracks struct {
	Cdn      string `json:"cdn"`
	Class    string `json:"class"`
	Language string `json:"language"`
	Profile  string `json:"profile"`
	URI      string `json:"uri"`
}
type Audio struct {
	Profile string   `json:"profile"`
	Tracks  []Tracks `json:"tracks"`
}
type Qualities struct {
	Audio     Audio  `json:"audio"`
	MimeCodec string `json:"mimeCodec"`
	Profile   string `json:"profile"`
	Quality   string `json:"quality"`
}
type Formitaeten struct {
	Facets     []string    `json:"facets"`
	IsAdaptive bool        `json:"isAdaptive"`
	MimeType   string      `json:"mimeType"`
	Profile    string      `json:"profile"`
	Qualities  []Qualities `json:"qualities"`
	Type       string      `json:"type"`
}
type PriorityList struct {
	Formitaeten []Formitaeten `json:"formitaeten"`
	Profile     string        `json:"profile"`
}

func GetPtmd(appToken, id string) (*Ptmd, error) {
	slog.Debug("GetPtmd", "id", id)
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.zdf.de/tmd/2/ngplayer_2_5/live/ptmd/%s", id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Api-Auth", "Bearer "+appToken)
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
