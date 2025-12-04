package ard

import (
	"net/url"
	"strings"
	"time"
)

type Links struct {
	Self struct {
		Type  string `json:"type"`
		Title string `json:"title"`
		Href  string `json:"href"`
	} `json:"self"`
}

type ChannelTrackingPiano struct {
	WidgetType  string `json:"widget_type"`
	WidgetTitle string `json:"widget_title"`
	WidgetID    string `json:"widget_id"`
}

type Image struct {
	Title        string      `json:"title"`
	Text         interface{} `json:"text"`
	Alt          string      `json:"alt"`
	Src          string      `json:"src"`
	ProducerName string      `json:"producerName"`
}

type Images struct {
	Aspect16X9 Image `json:"aspect16x9"`
}

type ChannelReference struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	MainChannelID string `json:"main_channel_id"`
}

type TrackingPianoTimeSlot struct {
	WidgetSection       string `json:"widget_section"`
	TeaserTitle         string `json:"teaser_title"`
	TeaserRecommended   bool   `json:"teaser_recommended"`
	TeaserInstitutionID string `json:"teaser_institution_id"`
	TeaserInstitution   string `json:"teaser_institution"`
	TeaserID            string `json:"teaser_id"`
	TeaserContentType   string `json:"teaser_content_type"`
}

type ChannelTimeSlot struct {
	ID                    string                `json:"id"`
	Links                 Links                 `json:"links"`
	Type                  string                `json:"type"`
	Title                 string                `json:"title"`
	Duration              int                   `json:"duration"`
	Images                *Images               `json:"images,omitempty"`
	Channel               ChannelReference      `json:"channel"`
	TrackingPiano         TrackingPianoTimeSlot `json:"trackingPiano"`
	CreationDate          time.Time             `json:"creationDate"`
	HeightUnits           int                   `json:"heightUnits"`
	BeginNet              time.Time             `json:"beginNet"`
	BinaryFeatures        []string              `json:"binaryFeatures,omitempty"`
	MaturityContentRating string                `json:"maturityContentRating,omitempty"`
	BroadcastEnd          time.Time             `json:"broadcastEnd"`
	BroadcastedOn         time.Time             `json:"broadcastedOn"`
	CoreSubline           string                `json:"coreSubline"`
	CoreTitle             string                `json:"coreTitle"`
	LastMod               time.Time             `json:"lastMod"`
	NumericID             string                `json:"numericId"`
	Subline               string                `json:"subline,omitempty"`
	Synopsis              string                `json:"synopsis"`
}

func (c *ChannelTimeSlot) Image() string {
	if c.Images == nil {
		return ""
	}
	src := c.Images.Aspect16X9.Src
	sp := strings.Split(src, "?")
	v, _ := url.ParseQuery(sp[1])
	return sp[0] + "?ch=" + v.Get("ch") + "&w=1920"
}

type PublicationService struct {
	Name    string `json:"name"`
	Partner string `json:"partner"`
}

type LocalChannel struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Crid         string `json:"crid"`
	LocalDefault bool   `json:"localDefault,omitempty"`
}

type Channel struct {
	ID                 string               `json:"id"`
	TrackingPiano      ChannelTrackingPiano `json:"trackingPiano"`
	TimeSlots          [][]ChannelTimeSlot  `json:"timeSlots"`
	PublicationService PublicationService   `json:"publicationService"`
	Crid               string               `json:"crid"`
	LocalChannelList   []LocalChannel       `json:"localChannelList,omitempty"`
}

type ProgramTrackingPiano struct {
	PageTitle         string `json:"page_title"`
	PageInstitutionID string `json:"page_institution_id"`
	PageInstitution   string `json:"page_institution"`
	PageID            string `json:"page_id"`
	PageChapter2      string `json:"page_chapter2"`
	PageChapter1      string `json:"page_chapter1"`
}

type ProgramTimeSlot struct {
	Title       string `json:"title"`
	HeightUnits int    `json:"heightUnits"`
	EndDate     string `json:"endDate"`
	StartDate   string `json:"startDate"`
}

type ArdProgram struct {
	Links         Links                `json:"links"`
	Channels      []*Channel           `json:"channels"`
	TrackingPiano ProgramTrackingPiano `json:"trackingPiano"`
	TimeSlots     []ProgramTimeSlot    `json:"timeSlots"`
	CreationDate  time.Time            `json:"creationDate"`
}
