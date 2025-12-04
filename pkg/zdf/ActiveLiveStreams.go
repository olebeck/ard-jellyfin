package zdf

import (
	"log/slog"
	"time"
)

type CurrentMediaNodes struct {
	PtmdTemplate   string    `json:"ptmdTemplate"`
	LiveMediaType  string    `json:"liveMediaType"`
	PlannedStart   any       `json:"plannedStart"`
	PlannedStop    any       `json:"plannedStop"`
	EditorialStart any       `json:"editorialStart"`
	EditorialStop  any       `json:"editorialStop"`
	Start          time.Time `json:"start"`
	Stop           time.Time `json:"stop"`
	State          string    `json:"state"`
	TvService      string    `json:"tvService"`
	GeoLocation    string    `json:"geoLocation"`
	Typename       string    `json:"__typename"`
}
type CurrentMedia struct {
	Nodes    []CurrentMediaNodes `json:"nodes"`
	Typename string              `json:"__typename"`
}
type Teaser struct {
	Title    string `json:"title"`
	Image    Image  `json:"image"`
	Typename string `json:"__typename"`
}
type PublicationFormInfo struct {
	Original    string `json:"original"`
	Transformed string `json:"transformed"`
	Typename    string `json:"__typename"`
}
type StructuralMetadata struct {
	IsChildrenContent   bool                `json:"isChildrenContent"`
	GenreInfo           GenreInfo           `json:"genreInfo"`
	PublicationFormInfo PublicationFormInfo `json:"publicationFormInfo"`
	VisualDimension     VisualDimension     `json:"visualDimension"`
}
type SmartCollection struct {
	Title              string             `json:"title"`
	ID                 string             `json:"id"`
	Canonical          string             `json:"canonical"`
	StructuralMetadata StructuralMetadata `json:"structuralMetadata"`
	Typename           string             `json:"__typename"`
}
type Availability struct {
	FskBlocked bool   `json:"fskBlocked"`
	Vod        Vod    `json:"vod"`
	Typename   string `json:"__typename"`
}
type Click struct {
	PageExternalID          string `json:"page_external_id"`
	OnsiteadType            string `json:"onsitead_type"`
	OnsiteadAdvertiser      any    `json:"onsitead_advertiser"`
	TargetTitle             string `json:"target_title"`
	TargetType              string `json:"target_type"`
	TargetSection           string `json:"target_section"`
	TargetGenre             string `json:"target_genre"`
	TargetIsnewscontent     bool   `json:"target_isnewscontent"`
	TargetIssportcontent    bool   `json:"target_issportcontent"`
	TargetIschildrencontent bool   `json:"target_ischildrencontent"`
}
type Zdf struct {
	View     string `json:"view"`
	Click    string `json:"click"`
	Config   Config `json:"config"`
	Pause    string `json:"pause"`
	Play     string `json:"play"`
	Autoplay string `json:"autoplay"`
}
type Piano struct {
	Click Click `json:"click"`
}
type Tracking struct {
	Piano    Piano  `json:"piano"`
	Zdf      Zdf    `json:"zdf"`
	Typename string `json:"__typename"`
}
type VideosNodes struct {
	ID                 string             `json:"id"`
	Canonical          string             `json:"canonical"`
	CurrentMediaType   string             `json:"currentMediaType"`
	EditorialDate      time.Time          `json:"editorialDate"`
	ScheduledMedia     any                `json:"scheduledMedia"`
	CurrentMedia       CurrentMedia       `json:"currentMedia"`
	Teaser             Teaser             `json:"teaser"`
	StructuralMetadata StructuralMetadata `json:"structuralMetadata"`
	SmartCollection    SmartCollection    `json:"smartCollection"`
	Availability       Availability       `json:"availability"`
	Tracking           Tracking           `json:"tracking"`
	Typename           string             `json:"__typename"`
}
type Videos struct {
	Nodes    []VideosNodes `json:"nodes"`
	Typename string        `json:"__typename"`
}

func GetActiveLiveStreams(appToken string) (*Videos, error) {
	slog.Debug("GetActiveLiveStreams")
	var body struct {
		Data struct {
			Videos Videos `json:"videos"`
		} `json:"data"`
	}
	err := doGraphql(appToken, map[string]any{
		"operationName": "getActiveLiveStreams",
		"variables":     map[string]any{},
		"extensions": map[string]any{
			"clientLibrary":  map[string]any{"name": "@apollo/client", "version": "4.0.9"},
			"persistedQuery": map[string]any{"version": 1, "sha256Hash": "d64a1821302c98e8498b76d51e8eeebc7f5c8c7b7874ac370403b73243e37fc9"},
		},
	}, &body)
	if err != nil {
		return nil, err
	}
	return &body.Data.Videos, nil
}
