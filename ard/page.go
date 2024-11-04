package ard

import (
	"net/url"
	"strings"
	"time"
)

type Page struct {
	CoreAssetType    string `json:"coreAssetType"`
	FskRating        string `json:"fskRating"`
	ID               string `json:"id"`
	IsChildContent   bool   `json:"isChildContent"`
	IsFamilyFriendly bool   `json:"isFamilyFriendly"`
	Personalized     bool   `json:"personalized"`
	Links            struct {
		Self struct {
			ID      string `json:"id"`
			URLID   string `json:"urlId"`
			Title   string `json:"title"`
			Href    string `json:"href"`
			Type    string `json:"type"`
			Partner string `json:"partner"`
		} `json:"self"`
	} `json:"links"`
	TargetAudienceAgeMax interface{} `json:"targetAudienceAgeMax"`
	TargetAudienceAgeMin interface{} `json:"targetAudienceAgeMin"`
	Title                string      `json:"title"`
	Tracking             struct {
		AggregationLevelID int `json:"aggregationLevelId"`
		AtiCustomVars      struct {
			ClipTitle             string `json:"clipTitle"`
			MediaType             string `json:"mediaType"`
			Lra                   string `json:"lra"`
			Channel               string `json:"channel"`
			Show                  string `json:"show"`
			ContentTypes          string `json:"contentTypes"`
			MediaDistributionType int    `json:"mediaDistributionType"`
			ContentID             int    `json:"contentId"`
			MetadataID            string `json:"metadataId"`
		} `json:"atiCustomVars"`
		Chapter2      string      `json:"chapter2"`
		Chapter3      interface{} `json:"chapter3"`
		EnvironmentID int         `json:"environmentId"`
		PageTitle     string      `json:"pageTitle"`
		SzmType       string      `json:"szmType"`
	} `json:"tracking"`
	TrackingPiano struct {
		PageChapter1      string `json:"page_chapter1"`
		PageChapter2      string `json:"page_chapter2"`
		PageContentID     string `json:"page_content_id"`
		PageContentTitle  string `json:"page_content_title"`
		PageID            string `json:"page_id"`
		PageInstitution   string `json:"page_institution"`
		PageInstitutionID string `json:"page_institution_id"`
		PagePublisher     string `json:"page_publisher"`
		PagePublisherID   string `json:"page_publisher_id"`
		PageTitle         string `json:"page_title"`
		Site              string `json:"site"`
	} `json:"trackingPiano"`
	Widgets []struct {
		AvailableTo    interface{}   `json:"availableTo"`
		BinaryFeatures []interface{} `json:"binaryFeatures"`
		BroadcastedOn  time.Time     `json:"broadcastedOn"`
		Embeddable     bool          `json:"embeddable"`
		Geoblocked     bool          `json:"geoblocked"`
		ID             string        `json:"id"`
		Image          struct {
			Alt          string `json:"alt"`
			ProducerName string `json:"producerName"`
			Src          string `json:"src"`
			Title        string `json:"title"`
		} `json:"image"`
		IsChildContent        bool   `json:"isChildContent"`
		MaturityContentRating string `json:"maturityContentRating"`
		MediaCollection       struct {
			Embedded struct {
				IsGeoBlocked bool `json:"isGeoBlocked"`
				Live         struct {
					DvrWindowSeconds int `json:"dvrWindowSeconds"`
				} `json:"live"`
				Meta struct {
					BroadcastedOnDateTime time.Time `json:"broadcastedOnDateTime"`
					ClipSourceName        string    `json:"clipSourceName"`
					Images                []struct {
						Alt             string `json:"alt"`
						AspectRatio     string `json:"aspectRatio"`
						ImageSourceName string `json:"imageSourceName"`
						Kind            string `json:"kind"`
						Title           string `json:"title"`
						URL             string `json:"url"`
					} `json:"images"`
					PublicationService struct {
						ID      string `json:"id"`
						Name    string `json:"name"`
						Partner string `json:"partner"`
					} `json:"publicationService"`
					Synopsis string `json:"synopsis"`
					Title    string `json:"title"`
				} `json:"meta"`
				PluginData struct {
					JumpmarksAll struct {
						RefreshSeconds int    `json:"refreshSeconds"`
						URL            string `json:"url"`
					} `json:"jumpmarks@all"`
					MultistreamAll struct {
						FeedURL           string `json:"feedUrl"`
						PollRate          int    `json:"pollRate"`
						NoContentPollRate int    `json:"noContentPollRate"`
					} `json:"multistream@all"`
					RecommendationAll struct {
						IsAutoplay   bool   `json:"isAutoplay"`
						TimerSeconds int    `json:"timerSeconds"`
						URL          string `json:"url"`
					} `json:"recommendation@all"`
					TrackingAgfAll struct {
						AppID    string `json:"appId"`
						ClipData struct {
							Assetid string `json:"assetid"`
							Length  string `json:"length"`
							NolC0   string `json:"nol_c0"`
							NolC10  string `json:"nol_c10"`
							NolC12  string `json:"nol_c12"`
							NolC15  string `json:"nol_c15"`
							NolC16  string `json:"nol_c16"`
							NolC18  string `json:"nol_c18"`
							NolC19  string `json:"nol_c19"`
							NolC2   string `json:"nol_c2"`
							NolC5   string `json:"nol_c5"`
							NolC7   string `json:"nol_c7"`
							NolC9   string `json:"nol_c9"`
							Program string `json:"program"`
							Title   string `json:"title"`
							Type    string `json:"type"`
						} `json:"clipData"`
						Tracker string `json:"tracker"`
					} `json:"trackingAgf@all"`
					TrackingAtiAll struct {
						Config struct {
							Site int `json:"site"`
						} `json:"config"`
						IsEnabled bool `json:"isEnabled"`
						RichMedia struct {
							BroadcastMode   string `json:"broadcastMode"`
							Duration        string `json:"duration"`
							MediaLabel      string `json:"mediaLabel"`
							MediaLevel2     string `json:"mediaLevel2"`
							MediaTheme1     string `json:"mediaTheme1"`
							MediaTheme2     string `json:"mediaTheme2"`
							MediaTheme3     string `json:"mediaTheme3"`
							MediaType       string `json:"mediaType"`
							PlayerID        string `json:"playerId"`
							RefreshDuration struct {
								Num0  int `json:"0"`
								Num1  int `json:"1"`
								Num12 int `json:"12"`
								Num18 int `json:"18"`
							} `json:"refreshDuration"`
						} `json:"richMedia"`
						Tracker string `json:"tracker"`
					} `json:"trackingAti@all"`
					TrackingPianoAll struct {
						AvContent struct {
							AvBroadcastingType string `json:"av_broadcasting_type"`
							AvContent          string `json:"av_content"`
							AvContentDuration  int    `json:"av_content_duration"`
							AvContentID        string `json:"av_content_id"`
							AvContentType      string `json:"av_content_type"`
							AvInstitution      string `json:"av_institution"`
							AvInstitutionID    string `json:"av_institution_id"`
							AvPublisher        string `json:"av_publisher"`
							AvPublisherID      string `json:"av_publisher_id"`
						} `json:"avContent"`
						Config struct {
							Events []string `json:"events"`
							Site   int      `json:"site"`
						} `json:"config"`
						IsEnabled bool `json:"isEnabled"`
					} `json:"trackingPiano@all"`
				} `json:"pluginData"`
				Streams []struct {
					Kind     string `json:"kind"`
					KindName string `json:"kindName"`
					Media    []struct {
						AspectRatio string `json:"aspectRatio"`
						Audios      []struct {
							Kind         string `json:"kind"`
							LanguageCode string `json:"languageCode"`
						} `json:"audios"`
						ForcedLabel                 string        `json:"forcedLabel"`
						HasEmbeddedSubtitles        bool          `json:"hasEmbeddedSubtitles"`
						IsAdaptiveQualitySelectable bool          `json:"isAdaptiveQualitySelectable"`
						IsHighDynamicRange          bool          `json:"isHighDynamicRange"`
						MaxHResolutionPx            int           `json:"maxHResolutionPx"`
						MaxVResolutionPx            int           `json:"maxVResolutionPx"`
						MimeType                    string        `json:"mimeType"`
						Subtitles                   []interface{} `json:"subtitles"`
						URL                         string        `json:"url"`
						VideoCodec                  string        `json:"videoCodec"`
					} `json:"media"`
					VideoLanguageCode string `json:"videoLanguageCode"`
				} `json:"streams"`
				Subtitles []interface{} `json:"subtitles"`
			} `json:"embedded"`
			Href string `json:"href"`
		} `json:"mediaCollection"`
		Pagination         interface{} `json:"pagination"`
		Personalized       bool        `json:"personalized"`
		PlayerConfig       interface{} `json:"playerConfig"`
		PublicationService struct {
			Name string `json:"name"`
			Logo struct {
				Title        string `json:"title"`
				Alt          string `json:"alt"`
				ProducerName string `json:"producerName"`
				Src          string `json:"src"`
				AspectRatio  string `json:"aspectRatio"`
			} `json:"logo"`
			PublisherType string `json:"publisherType"`
			Partner       string `json:"partner"`
			ID            string `json:"id"`
			CoreID        string `json:"coreId"`
		} `json:"publicationService"`
		Links struct {
			Self struct {
				ID      string `json:"id"`
				URLID   string `json:"urlId"`
				Title   string `json:"title"`
				Href    string `json:"href"`
				Type    string `json:"type"`
				Partner string `json:"partner"`
			} `json:"self"`
		} `json:"links"`
		Relates []struct {
			BinaryFeatures        []any     `json:"binaryFeatures"`
			BroadcastedOn         time.Time `json:"broadcastedOn"`
			CoreAssetType         string    `json:"coreAssetType"`
			ID                    string    `json:"id"`
			Images                *Images   `json:"images"`
			IsChildContent        bool      `json:"isChildContent"`
			LongTitle             string    `json:"longTitle"`
			MaturityContentRating string    `json:"maturityContentRating"`
			MediumTitle           string    `json:"mediumTitle"`
			Personalized          bool      `json:"personalized"`
			Playtime              any       `json:"playtime"`
			PublicationService    struct {
				Name string `json:"name"`
				Logo struct {
					Title        string `json:"title"`
					Alt          string `json:"alt"`
					ProducerName string `json:"producerName"`
					Src          string `json:"src"`
					AspectRatio  string `json:"aspectRatio"`
				} `json:"logo"`
				PublisherType string `json:"publisherType"`
				Partner       string `json:"partner"`
				ID            string `json:"id"`
				CoreID        string `json:"coreId"`
			} `json:"publicationService"`
			Links struct {
				Self struct {
					ID      string `json:"id"`
					URLID   string `json:"urlId"`
					Title   string `json:"title"`
					Href    string `json:"href"`
					Type    string `json:"type"`
					Partner string `json:"partner"`
				} `json:"self"`
				Target struct {
					ID      string `json:"id"`
					URLID   string `json:"urlId"`
					Title   string `json:"title"`
					Href    string `json:"href"`
					Type    string `json:"type"`
					Partner string `json:"partner"`
				} `json:"target"`
			} `json:"links"`
			ShortTitle    string      `json:"shortTitle"`
			Show          interface{} `json:"show"`
			Subtitled     bool        `json:"subtitled"`
			TitleVisible  bool        `json:"titleVisible"`
			TrackingPiano struct {
				TeaserContentType   string `json:"teaser_content_type"`
				TeaserID            string `json:"teaser_id"`
				TeaserInstitution   string `json:"teaser_institution"`
				TeaserInstitutionID string `json:"teaser_institution_id"`
				TeaserTitle         string `json:"teaser_title"`
			} `json:"trackingPiano"`
			Type string `json:"type"`
		} `json:"relates"`
		Show          interface{} `json:"show"`
		Synopsis      string      `json:"synopsis"`
		Title         string      `json:"title"`
		TrackingPiano struct {
			TeaserRecommended bool   `json:"teaser_recommended"`
			WidgetID          string `json:"widget_id"`
			WidgetTitle       string `json:"widget_title"`
			WidgetTyp         string `json:"widget_typ"`
		} `json:"trackingPiano"`
		Type string `json:"type"`
	} `json:"widgets"`
}

func (p *Page) PlaylistURL() string {
	return p.Widgets[0].MediaCollection.Embedded.Streams[0].Media[0].URL
}

func (p *Page) Image() string {
	src := p.Widgets[0].Image.Src
	sp := strings.Split(src, "?")
	v, _ := url.ParseQuery(sp[1])
	return sp[0] + "?ch=" + v.Get("ch") + "&w=1920"
}
