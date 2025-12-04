package main

import (
	"log/slog"
	"slices"
	"strings"
	"time"

	"github.com/olebeck/ard-jellyfin/pkg/zdf"
)

// livestream -> epg id
var channelsMap = map[string]string{
	"zdf-live-beitrag-100":     "ZDF",
	"zdfinfo-live-beitrag-100": "ZDFinfo",
	"zdfneo-live-beitrag-100":  "ZDFneo",
	"3sat-live-beitrag-100":    "3sat",
	"kika-live-beitrag-100":    "KiKA",
	"phoenix-live-beitrag-100": "PHOENIX",
	"arte-livestream-100":      "arte",
}

func addZdf(epg *XmlTvOutput, m3u8 *M3U8Channels) error {
	appToken, err := zdf.GetAppToken()
	if err != nil {
		return err
	}

	channelTitles := make(map[string]string)
	channelImages := make(map[string]string)

	now := time.Now().Add(-12 * time.Hour)
	skipPhoenix := m3u8.HaveChannel("phoenix")
	skipKika := m3u8.HaveChannel("kika")
	for day := -1; day < 7; day++ {
		start := now.Add(time.Duration(day) * 24 * time.Hour)
		epgs, err := zdf.GetEpg(appToken, start, start.Add(24*time.Hour))
		if err != nil {
			return err
		}
		for _, epgEntry := range epgs {
			if skipPhoenix && epgEntry.Broadcaster.ID == "PHOENIX" {
				continue
			}
			if skipKika && epgEntry.Broadcaster.ID == "KI.KA" {
				continue
			}
			if epg.AddChannel(
				epgEntry.Broadcaster.ID,
				epgEntry.Broadcaster.Title,
				epgEntry.Broadcaster.Logo.Layouts["original"],
			) {
				channelImages[epgEntry.Broadcaster.ID] = epgEntry.Broadcaster.Logo.Layouts["original"]
				channelTitles[epgEntry.Broadcaster.ID] = epgEntry.Broadcaster.Title
			}
			for _, broadcast := range epgEntry.Broadcasts {
				desc := broadcast.Text
				desc = strings.ReplaceAll(desc, "<br/>", "\n")
				epg.AddProgramme(
					epgEntry.Broadcaster.ID,
					broadcast.Title,
					desc,
					broadcast.Subheadline,
					broadcast.AirtimeBegin,
					broadcast.AirtimeEnd,
					"",
				)
			}
		}
	}

	activeLivestreams, err := zdf.GetActiveLiveStreams(appToken)
	if err != nil {
		return err
	}

	for _, videoNode := range activeLivestreams.Nodes {
		if len(videoNode.CurrentMedia.Nodes) == 0 {
			continue
		}

		channelId, ok := channelsMap[videoNode.Canonical]
		if !ok {
			// skip not in the EPG
			continue
		}

		if skipPhoenix && channelId == "PHOENIX" {
			continue
		}
		if skipKika && channelId == "KI.KA" {
			continue
		}

		mediaNode := videoNode.CurrentMedia.Nodes[0]
		ptmdSplit := strings.Split(mediaNode.PtmdTemplate, "/")
		ptmdId := ptmdSplit[len(ptmdSplit)-1]
		ptmd, err := zdf.GetPtmd(appToken, ptmdId)
		if err != nil {
			return err
		}
		if ptmd.ID == "" {
			slog.Debug("ptmd not found", "channelId", channelId, "ptmdId", ptmdId)
			continue
		}
		video, err := zdf.GetVideoByCanonical(appToken, videoNode.Canonical)
		if err != nil {
			return err
		}

		imageUrl, ok := channelImages[channelId]
		title := channelTitles[channelId]
		if !ok {
			imageUrl, _ = video.Teaser.Image.List["original"].(string)
			title, _ = strings.CutSuffix(video.Title, " Livestream")
		}

		qualityIndex := slices.IndexFunc(ptmd.PriorityList[0].Formitaeten[0].Qualities, func(quality zdf.Qualities) bool {
			return quality.Quality == "veryhigh"
		})
		if qualityIndex == -1 {
			qualityIndex = 0
		}
		quality := ptmd.PriorityList[0].Formitaeten[0].Qualities[qualityIndex]
		trackIndex := slices.IndexFunc(quality.Audio.Tracks, func(track zdf.Tracks) bool {
			return track.Language == "deu"
		})
		if trackIndex == -1 {
			trackIndex = 0
		}
		track := quality.Audio.Tracks[0]
		playlistUrl := track.URI

		m3u8.AddChannel(M3U8Channel{
			ID:          channelId,
			Name:        title,
			Image:       imageUrl,
			PlaylistURL: playlistUrl,
		})
	}

	return nil
}
