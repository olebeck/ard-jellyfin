package xmltv

import (
	"encoding/xml"
	"time"
)

type TV struct {
	XMLName           xml.Name    `xml:"tv"`
	SourceInfoURL     string      `xml:"source-info-url,attr,omitempty"`
	SourceInfoName    string      `xml:"source-info-name,attr,omitempty"`
	GeneratorInfoName string      `xml:"generator-info-name,attr,omitempty"`
	Channels          []Channel   `xml:"channel"`
	Programmes        []Programme `xml:"programme"`
}

type Channel struct {
	ID           string        `xml:"id,attr"`
	DisplayNames []DisplayName `xml:"display-name"`
	Icons        []Icon        `xml:"icon"`
	URLs         []string      `xml:"url"`
}

type DisplayName struct {
	Lang  string `xml:"lang,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Icon struct {
	Src    string `xml:"src,attr,omitempty"`
	Width  string `xml:"width,attr,omitempty"`
	Height string `xml:"height,attr,omitempty"`
}

type Image struct {
	Type   string `xml:"type,attr"`
	Size   int    `xml:"size,attr"`
	Orient string `xml:"orient,attr"`
	System string `xml:"system,attr"`
	Value  string `xml:",chardata"`
}

type Programme struct {
	Start        XMLTVTime     `xml:"start,attr"`
	Stop         XMLTVTime     `xml:"stop,attr,omitempty"`
	ChannelID    string        `xml:"channel,attr"`
	Titles       []Title       `xml:"title"`
	Subtitles    []Subtitle    `xml:"sub-title"`
	Descriptions []Description `xml:"desc"`
	Categories   []Category    `xml:"category"`
	Rating       *Rating       `xml:"rating,omitempty"`
	EpisodeNum   []EpisodeNum  `xml:"episode-num"`
	Icon         Icon          `xml:"icon"`
	Live         *struct{}     `xml:"live"`
}

type Title struct {
	Lang  string `xml:"lang,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Subtitle struct {
	Lang  string `xml:"lang,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Description struct {
	Lang  string `xml:"lang,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Category struct {
	Lang  string `xml:"lang,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Rating struct {
	System string `xml:"system,attr,omitempty"`
	Value  string `xml:"value"`
	Icons  []Icon `xml:"icon"`
}

type EpisodeNum struct {
	System string `xml:"system,attr,omitempty"`
	Value  string `xml:",chardata"`
}

type XMLTVTime struct {
	time.Time
}

func (t *XMLTVTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: t.Time.Format("20060102150405 -0700"),
	}, nil
}

func (t *XMLTVTime) UnmarshalXMLAttr(attr xml.Attr) error {
	parsedTime, err := time.Parse("20060102150405 -0700", attr.Value)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}
