package zdf

import (
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var client http.Client

func findElements(n *html.Node, attrName string) []*html.Node {
	var nodes []*html.Node
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == attrName {
					nodes = append(nodes, n)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return nodes
}

func GetLiveTvPlayers() ([]Player, error) {
	res, err := client.Get("https://www.zdf.de/live-tv")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	nodes, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	playerElements := findElements(nodes, "data-zdfplayer-jsb")
	var players []Player

	type Jsb struct {
		Config                string `json:"config"`
		ConfigTiviAppIos      string `json:"configTiviAppIos"`
		ConfigTiviAppAndroid  string `json:"configTiviAppAndroid"`
		ConfigHeuteAppIos     string `json:"configHeuteAppIos"`
		ConfigHeuteAppAndroid string `json:"configHeuteAppAndroid"`
		Content               string `json:"content"`
		EmbedContent          string `json:"embed_content"`
		Autoplay              bool   `json:"autoplay"`
		APIToken              string `json:"apiToken"`
		Preload               bool   `json:"preload"`
		Primary               bool   `json:"primary"`
		IsLivestream          bool   `json:"isLivestream"`
		Fsk16ErrorDisabled    bool   `json:"fsk16ErrorDisabled"`
		AspectRatio           string `json:"aspectRatio"`
	}
	for _, node := range playerElements {
		var id string
		var jsb Jsb
		for _, attr := range node.Attr {
			switch attr.Key {
			case "data-zdfplayer-id":
				id = attr.Val
			case "data-zdfplayer-jsb":
				err = json.Unmarshal([]byte(attr.Val), &jsb)
				if err != nil {
					return nil, err
				}
			}
		}

		content, err := getPlayerContent(jsb.APIToken, jsb.Content)
		if err != nil {
			return nil, err
		}

		ptmd, err := getPtmd(jsb.APIToken, content.MainVideoContent.HTTPZdfDeRelsTarget.Streams.Default.ExtID)
		if err != nil {
			return nil, err
		}

		players = append(players, Player{
			ID:       strings.Split(id, "-")[0],
			APIToken: jsb.APIToken,
			Content:  content,
			Ptmd:     ptmd,
		})
	}
	return players, nil
}
