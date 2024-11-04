package zdf

import (
	"encoding/json"
	"net/http"
	"time"
)

type PlayerContent struct {
	ContentType    string    `json:"contentType"`
	OnlineOnly     bool      `json:"onlineOnly"`
	Title          string    `json:"title"`
	EditorialDate  time.Time `json:"editorialDate"`
	LeadParagraph  string    `json:"leadParagraph"`
	Teasertext     string    `json:"teasertext"`
	TeaserImageRef struct {
		Title     string `json:"title"`
		AltText   string `json:"altText"`
		Source    string `json:"source"`
		Profile   string `json:"profile"`
		Self      string `json:"self"`
		Canonical string `json:"canonical"`
		Layouts   struct {
			One280X720  string `json:"1280x720"`
			One920X1080 string `json:"1920x1080"`
			Two400X1350 string `json:"2400x1350"`
			Three84X216 string `json:"384x216"`
			Seven68X432 string `json:"768x432"`
			Original    string `json:"original"`
		} `json:"layouts"`
	} `json:"teaserImageRef"`
	EmbeddingPossible     bool `json:"embeddingPossible"`
	HTTPZdfDeRelsCategory struct {
		Profile   string `json:"profile"`
		Self      string `json:"self"`
		Canonical string `json:"canonical"`
	} `json:"http://zdf.de/rels/category"`
	HTTPZdfDeRelsURI        string `json:"http://zdf.de/rels/uri"`
	HTTPZdfDeRelsSharingURL string `json:"http://zdf.de/rels/sharing-url"`
	HTTPZdfDeRelsBrand      struct {
		Profile   string `json:"profile"`
		Self      string `json:"self"`
		Canonical string `json:"canonical"`
		Title     string `json:"title"`
	} `json:"http://zdf.de/rels/brand"`
	TvService string `json:"tvService"`
	Profile   string `json:"profile"`
	Self      string `json:"self"`
	Canonical string `json:"canonical"`
	Tracking  struct {
		AtInternet struct {
			Page struct {
				Level1       string `json:"level1"`
				Level2       string `json:"level2"`
				Chapter1     string `json:"chapter1"`
				Chapter2     string `json:"chapter2"`
				Chapter3     string `json:"chapter3"`
				Name         string `json:"name"`
				CustomObject struct {
					Domain     string `json:"domain"`
					Level1     string `json:"level1"`
					Level2     string `json:"level2"`
					Chapter1   string `json:"chapter1"`
					Broadcast  string `json:"broadcast"`
					ID         string `json:"id"`
					InhaltsTyp string `json:"inhaltsTyp"`
					Chapter4   string `json:"chapter4"`
				} `json:"customObject"`
			} `json:"page"`
			RichMedia struct {
				MediaType       string `json:"mediaType"`
				MediaLevel2     string `json:"mediaLevel2"`
				MediaLabel      string `json:"mediaLabel"`
				MediaTheme1     string `json:"mediaTheme1"`
				MediaTheme2     string `json:"mediaTheme2"`
				MediaTheme3     string `json:"mediaTheme3"`
				RefreshDuration struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
					Num5 int `json:"5"`
				} `json:"refreshDuration"`
				BroadcastMode string `json:"broadcastMode"`
			} `json:"richMedia"`
			Event struct {
				Level1       string `json:"level1"`
				Level2       string `json:"level2"`
				Chapter1     string `json:"chapter1"`
				CustomObject struct {
					Domain     string `json:"domain"`
					Level1     string `json:"level1"`
					Level2     string `json:"level2"`
					Chapter1   string `json:"chapter1"`
					Broadcast  string `json:"broadcast"`
					ID         string `json:"id"`
					InhaltsTyp string `json:"inhaltsTyp"`
				} `json:"customObject"`
			} `json:"event"`
			CustomVars struct {
				Site struct {
					ID       string `json:"id"`
					Chapter2 string `json:"chapter2"`
					Chapter3 string `json:"chapter3"`
				} `json:"site"`
			} `json:"customVars"`
		} `json:"atInternet"`
		Nielsen struct {
			Content struct {
				Type    string `json:"type"`
				Assetid string `json:"assetid"`
				Program string `json:"program"`
				Title   string `json:"title"`
				Length  string `json:"length"`
				NolC0   string `json:"nol_c0"`
				NolC2   string `json:"nol_c2"`
				NolC5   string `json:"nol_c5"`
				NolC7   string `json:"nol_c7"`
				NolC8   string `json:"nol_c8"`
				NolC9   string `json:"nol_c9"`
				NolC10  string `json:"nol_c10"`
				NolC12  string `json:"nol_c12"`
				NolC14  string `json:"nol_c14"`
				NolC15  string `json:"nol_c15"`
				NolC16  string `json:"nol_c16"`
				NolC18  string `json:"nol_c18"`
				NolC19  string `json:"nol_c19"`
			} `json:"content"`
			AddContent struct {
				NolC10 string `json:"nol_c10"`
				NolC14 string `json:"nol_c14"`
				NolC15 string `json:"nol_c15"`
				NolC19 string `json:"nol_c19"`
			} `json:"addContent"`
		} `json:"nielsen"`
		Szmng struct {
			St string `json:"st"`
			Cp string `json:"cp"`
			Sv string `json:"sv"`
			Co string `json:"co"`
		} `json:"szmng"`
		Zdf struct {
			ConfigV2 struct {
				PlayerTrackingRateInSeconds int `json:"playerTrackingRateInSeconds"`
			} `json:"config-v2"`
			TemplatesV2 struct {
				Click    string `json:"click"`
				View     string `json:"view"`
				Pause    string `json:"pause"`
				Play     string `json:"play"`
				Autoplay string `json:"autoplay"`
			} `json:"templates-v2"`
		} `json:"zdf"`
	} `json:"tracking"`
	HTTPZdfDeRelsNextVideoPersonalized string `json:"http://zdf.de/rels/next-video-personalized"`
	HTTPZdfDeRelsNextVideoTimeout      int    `json:"http://zdf.de/rels/next-video-timeout"`
	CurrentVideoType                   string `json:"currentVideoType"`
	MainVideoContent                   struct {
		HTTPZdfDeRelsTarget struct {
			Profile                          string `json:"profile"`
			Self                             string `json:"self"`
			Canonical                        string `json:"canonical"`
			HTTPZdfDeRelsStreamsPtmdTemplate string `json:"http://zdf.de/rels/streams/ptmd-template"`
			Streams                          struct {
				Default struct {
					Label                            string `json:"label"`
					ExtID                            string `json:"extId"`
					HTTPZdfDeRelsStreamsPtmdTemplate string `json:"http://zdf.de/rels/streams/ptmd-template"`
				} `json:"default"`
			} `json:"streams"`
		} `json:"http://zdf.de/rels/target"`
		Profile string `json:"profile"`
	} `json:"mainVideoContent"`
}

func (p *PlayerContent) Image() string {
	return p.TeaserImageRef.Layouts.Original
}

func getPlayerContent(token, url string) (*PlayerContent, error) {
	req, err := http.NewRequest("GET", url, nil)
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
	var content PlayerContent
	err = d.Decode(&content)
	if err != nil {
		return nil, err
	}
	return &content, nil
}
