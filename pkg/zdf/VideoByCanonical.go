package zdf

import (
	"log/slog"
	"time"
)

type Image struct {
	AltText string         `json:"altText"`
	Caption any            `json:"caption"`
	List    map[string]any `json:"list"`
}
type ContentOwner struct {
	Title   string `json:"title"`
	Details any    `json:"details"`
}
type StreamingOptions struct {
	Ad  bool   `json:"ad"`
	Ut  bool   `json:"ut"`
	Dgs bool   `json:"dgs"`
	Ov  bool   `json:"ov"`
	Ks  bool   `json:"ks"`
	Fsk string `json:"fsk"`
}
type EpisodeInfo struct {
	EpisodeNumber          any  `json:"episodeNumber"`
	SeasonNumber           any  `json:"seasonNumber"`
	HideEpisodeInformation bool `json:"hideEpisodeInformation"`
}
type GenreInfo struct {
	Original    any `json:"original"`
	Transformed any `json:"transformed"`
}
type Moods struct {
	Nodes []any `json:"nodes"`
}
type VisualDimension struct {
	Moods Moods `json:"moods"`
}

type Seo struct {
	Title string `json:"title"`
}
type Vod struct {
	Visible     any    `json:"visible"`
	VisibleFrom any    `json:"visibleFrom"`
	VisibleTo   any    `json:"visibleTo"`
	Fsk         string `json:"fsk"`
}
type External struct {
	StreamAnchorSourceURL         any `json:"streamAnchorSourceUrl"`
	StreamAnchorSourceURLTemplate any `json:"streamAnchorSourceUrlTemplate"`
}
type Content struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Length  string `json:"length"`
	NolC0   string `json:"nol_c0"`
	NolC2   string `json:"nol_c2"`
	NolC5   string `json:"nol_c5"`
	NolC7   string `json:"nol_c7"`
	NolC9   string `json:"nol_c9"`
	Assetid string `json:"assetid"`
	NolC10  string `json:"nol_c10"`
	NolC12  string `json:"nol_c12"`
	NolC13  string `json:"nol_c13"`
	NolC14  string `json:"nol_c14"`
	NolC15  string `json:"nol_c15"`
	NolC16  string `json:"nol_c16"`
	NolC18  string `json:"nol_c18"`
	NolC19  string `json:"nol_c19"`
	Program string `json:"program"`
}
type Nielsen struct {
	Content Content `json:"content"`
}
type Config struct {
	PlayerTrackingRateInSeconds int `json:"playerTrackingRateInSeconds"`
}

type Video struct {
	AvExternalID           string `json:"av_external_id"`
	AvContentID            string `json:"av_content_id"`
	AvSection              any    `json:"av_section"`
	AvContentfamily        any    `json:"av_contentfamily"`
	AvTitle                string `json:"av_title"`
	AvSmartCollectionID    any    `json:"av_smart_collection_id"`
	AvShow                 any    `json:"av_show"`
	AvContentType          string `json:"av_content_type"`
	AvType                 string `json:"av_type"`
	AvContentDuration      any    `json:"av_content_duration"`
	AvIschildrencontent    bool   `json:"av_ischildrencontent"`
	AvBroadcaster          string `json:"av_broadcaster"`
	AvBroadcastdetail      any    `json:"av_broadcastdetail"`
	AvEpisodeNo            any    `json:"av_episode_no"`
	AvGenre                any    `json:"av_genre"`
	AvPublicationTimestamp any    `json:"av_publication_timestamp"`
}
type VideoByCanonical struct {
	ID                 string             `json:"id"`
	Canonical          string             `json:"canonical"`
	RecoModel          string             `json:"recoModel"`
	ContentType        string             `json:"contentType"`
	CurrentMediaType   string             `json:"currentMediaType"`
	Title              string             `json:"title"`
	SharingURL         string             `json:"sharingUrl"`
	LeadParagraph      string             `json:"leadParagraph"`
	EditorialDate      time.Time          `json:"editorialDate"`
	ProductionYear     int                `json:"productionYear"`
	Teaser             Teaser             `json:"teaser"`
	ContentOwner       ContentOwner       `json:"contentOwner"`
	StreamingOptions   StreamingOptions   `json:"streamingOptions"`
	EpisodeInfo        EpisodeInfo        `json:"episodeInfo"`
	StructuralMetadata StructuralMetadata `json:"structuralMetadata"`
	SmartCollection    any                `json:"smartCollection"`
	Seo                Seo                `json:"seo"`
	Availability       Availability       `json:"availability"`
	Subtitle           any                `json:"subtitle"`
	WebURL             string             `json:"webUrl"`
	EmbeddingPossible  bool               `json:"embeddingPossible"`
	PublicationDate    time.Time          `json:"publicationDate"`
	External           External           `json:"external"`
	ScheduledMedia     any                `json:"scheduledMedia"`
	CurrentMedia       CurrentMedia       `json:"currentMedia"`
	Tracking           Tracking           `json:"tracking"`
	NextEditorialVideo any                `json:"nextEditorialVideo"`
}

var queryVideoByCanonical = `
query VideoByCanonical($canonical: String!, $first: Int) {
    videoByCanonical(canonical: $canonical) {
        id
        canonical
        recoModel
        contentType
        currentMediaType
        title
        sharingUrl
        leadParagraph
        editorialDate
        productionYear
        teaser { 
            image {
                altText
                caption
                list
            }
        }
        contentOwner {
            title
            details
        }
        streamingOptions {
            ad
            ut
            dgs
            ov
            ks
            fsk
        }
        episodeInfo {
            episodeNumber
            seasonNumber
            hideEpisodeInformation
        }
        structuralMetadata {
            isChildrenContent
            genreInfo {
                original
                transformed
            }
            publicationFormInfo {
                original
                transformed
            }
            visualDimension {
                moods(first: $first) {
                    nodes {
                        mood
                    }
                }
            }
        }
        smartCollection {
            id
            canonical  
            title
            collectionType
            sharingUrl
            structuralMetadata {
                contentFamily  
                publicationFormInfo {
                    original
                    transformed
                }
            }
        }
        seo {
            title
        }
        availability {
            fskBlocked
            vod {
                visible
                visibleFrom
                visibleTo
                fsk
            }
        }  
        subtitle
        webUrl
        embeddingPossible
        publicationDate
        external {
            streamAnchorSourceUrl
            streamAnchorSourceUrlTemplate
        }
        scheduledMedia {
            availableFrom
            availableTo
            editorialStart
            editorialStop
        }
        currentMedia {
            nodes {
                ptmdTemplate
                ... on VodMedia {
                    duration
                    aspectRatio
                    visible
                    geoLocation
                    highestVerticalResolution
                    streamAnchorTags {
                        nodes {
                            anchorOffset
                            anchorLabel
                        }
                    }
                    skipIntro {
                        startIntroTimeOffset
                        stopIntroTimeOffset
                        skipButtonDisplayTime
                        skipButtonLabel
                    }
                    vodMediaType
                    label
                    contentType
                }
                ... on LiveMedia {
                    geoLocation
                    tvService
                    title
                    start
                    stop
                    editorialStart
                    editorialStop
                    encryption
                    liveMediaType
                    label
                }
                id
            }
        }
        tracking {
            nielsen
            zdf
            piano(filter: video) 
        }
        nextEditorialVideo {
            id
            canonical
            recoModel
            contentType
            currentMediaType
            title
            sharingUrl
            leadParagraph
            editorialDate
            productionYear
            teaser { 
                image {
                    altText
                    caption
                    list
                }
            }
            contentOwner {
                title
                details
            }
            streamingOptions {
                ad
                ut
                dgs
                ov
                ks
                fsk
            }
            episodeInfo {
                episodeNumber
                seasonNumber
                hideEpisodeInformation
            }
            structuralMetadata {
                isChildrenContent
                genreInfo {
                    original
                    transformed
                }
                publicationFormInfo {
                    original
                    transformed
                }
                visualDimension {
                    moods(first: $first) {
                        nodes {
                            mood
                        }
                    }
                }
            }
            smartCollection {
                id
                canonical  
                title
                collectionType
                sharingUrl
                structuralMetadata {
                    contentFamily  
                    publicationFormInfo {
                        original
                        transformed
                    }
                }
            }
            seo {
                title
            }
            availability {
                fskBlocked
                vod {
                    visible
                    visibleFrom
                    visibleTo
                    fsk
                }
            }
            recoModel
            currentMedia {
                nodes {
                    ptmdTemplate
                    ... on VodMedia {
                        duration
                    }
                }
            }
        }
    }
}`

func GetVideoByCanonical(appToken string, canonical string) (*VideoByCanonical, error) {
	slog.Debug("GetVideoByCanonical", "canonical", canonical)
	var body struct {
		Data struct {
			VideoByCanonical VideoByCanonical `json:"videoByCanonical"`
		} `json:"data"`
	}
	err := doGraphql(appToken, map[string]any{
		"operationName": "VideoByCanonical",
		"query":         queryVideoByCanonical,
		"variables": map[string]any{
			"canonical": canonical,
			"first":     1,
		},
	}, &body)
	if err != nil {
		return nil, err
	}
	return &body.Data.VideoByCanonical, nil
}
