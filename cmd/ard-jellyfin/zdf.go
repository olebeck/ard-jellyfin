package main

import (
	"github.com/olebeck/ard-jellyfin/zdf"
)

func addZdf(out *XmlTvOutput, m3u8 *M3U8Channels) error {
	players, err := zdf.GetLiveTvPlayers()
	if err != nil {
		return err
	}

	for _, player := range players {
		var url string
		for _, quality := range player.Ptmd.PriorityList[0].Formitaeten[0].Qualities {
			if quality.Quality == "veryhigh" {
				url = quality.Audio.Tracks[0].URI
				break
			}
		}
		m3u8.AddChannel(M3U8Channel{
			ID:          player.ID,
			Name:        player.Title(),
			Image:       player.Content.Image(),
			PlaylistURL: url,
		})

		if out.AddChannel(player.ID, player.Title(), player.Content.Image()) {
			// todo epg
		}

	}

	return nil
}
