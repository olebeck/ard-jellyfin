package zdf

import "strings"

type Player struct {
	ID       string
	APIToken string
	Content  *PlayerContent
	Ptmd     *Ptmd
}

func (p *Player) Title() string {
	return strings.Replace(p.Content.Title, " Livestream", "", 1)
}
