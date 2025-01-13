package model

import (
	"fmt"
	"strings"
)

type Position struct {
	Wheel_id int
	Corn int
	GetOptions Options
	PData *PalenqueData
	CData *ChichenData
}

type SpecificPosition struct {
	Wheel_id int
	Corn int
	Execute Option
	FirstPlayer bool
}

type PalenqueData struct {
	CornTiles int
	WoodTiles int
}

type ChichenData struct {
	Full bool
}

func MakePData(hasWood bool) *PalenqueData {
	// todo based on player count
	if (hasWood) {
		return &PalenqueData{
			CornTiles: 4,
			WoodTiles: 4,
		}
	}

	return &PalenqueData{
		CornTiles: 4,
		WoodTiles: 0,
	}
}

func (pd *PalenqueData) HasCornShowing() bool {
	return pd.CornTiles > pd.WoodTiles
}

func MakeCData() *ChichenData {
	return &ChichenData {
		Full: false,
	}
}

func (p *SpecificPosition) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "%d %d", p.Wheel_id, p.Corn)
	return br.String()
}

// func PlayerOption(f func (*Player)) Option {
// 	return func(g *Game, c Color) {
// 		p := g.GetPlayerByColor(c)
// 		f(p)
// 	}
// }

// func SimpleOption(f func (*Player)) Options {
// 	return func(g *Game) []Option {
// 		return []Option{
// 			PlayerOption(f),
// 		}
// 	}
// }