package zdf

import (
	"fmt"
	"log/slog"
	"time"
)

type Next struct {
	Title        string    `json:"title"`
	AirtimeEnd   time.Time `json:"airtimeEnd"`
	AirtimeBegin time.Time `json:"airtimeBegin"`
	Typename     string    `json:"__typename"`
}

type Now struct {
	Title                 string    `json:"title"`
	Subheadline           any       `json:"subheadline"`
	Text                  string    `json:"text"`
	AirtimeBegin          time.Time `json:"airtimeBegin"`
	EffectiveAirtimeBegin time.Time `json:"effectiveAirtimeBegin"`
	AirtimeEnd            time.Time `json:"airtimeEnd"`
	EffectiveAirtimeEnd   time.Time `json:"effectiveAirtimeEnd"`
	TvService             string    `json:"tvService"`
	Video                 any       `json:"video"`
	Image                 any       `json:"image"`
	Typename              string    `json:"__typename"`
}

type Logo struct {
	AltText  string            `json:"altText"`
	Layouts  map[string]string `json:"layouts"`
	Typename string            `json:"__typename"`
}

type Broadcaster struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Next     Next   `json:"next"`
	Now      Now    `json:"now"`
	Logo     Logo   `json:"logo"`
	Typename string `json:"__typename"`
}

type Broadcasts struct {
	Title                 string    `json:"title"`
	Subheadline           string    `json:"subheadline"`
	Text                  string    `json:"text"`
	AirtimeBegin          time.Time `json:"airtimeBegin"`
	EffectiveAirtimeBegin time.Time `json:"effectiveAirtimeBegin"`
	AirtimeEnd            time.Time `json:"airtimeEnd"`
	EffectiveAirtimeEnd   time.Time `json:"effectiveAirtimeEnd"`
	TvService             string    `json:"tvService"`
	Video                 any       `json:"video"`
	Image                 any       `json:"image"`
	Typename              string    `json:"__typename"`
}

type Epg struct {
	Broadcaster Broadcaster  `json:"broadcaster"`
	Broadcasts  []Broadcasts `json:"broadcasts"`
	Typename    string       `json:"__typename"`
}

type GetEpgData struct {
	Data struct {
		Epg []Epg `json:"epg"`
	} `json:"data"`
}

func GetEpg(apiToken string, from, to time.Time) ([]Epg, error) {
	slog.Debug("GetEpg", "from", from, "to", to)
	var query = map[string]any{
		"operationName": "getEpg",
		"variables": map[string]any{
			"filter": map[string]any{
				"broadcasterIds": []string{"ZDF", "ZDFinfo", "ZDFneo", "3sat", "KI.KA", "PHOENIX", "arte"},
				"from":           from.Format(time.RFC3339),
				"to":             to.Format(time.RFC3339),
			},
		},
		"extensions": map[string]any{
			"clientLibrary":  map[string]any{"name": "@apollo/client", "version": "4.0.9"},
			"persistedQuery": map[string]any{"version": 1, "sha256Hash": "e9b840c60dab34c0cc2ab81673b4e34a680684b01502635a217bce7e91acf9a8"},
		},
	}
	var body GetEpgData
	err := doGraphql(apiToken, query, &body)
	if err != nil {
		return nil, fmt.Errorf("getEpg: %w", err)
	}
	return body.Data.Epg, nil
}
