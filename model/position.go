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
			cornTiles: 4,
			woodTiles: 4,
		}
	}

	return &PalenqueData{
		cornTiles: 4,
		woodTiles: 0,
	}
}

func (pd *PalenqueData) HasCornShowing() bool {
	return pd.cornTiles > pd.woodTiles
}

func MakeCData() *ChichenData {
	return &ChichenData {
		full: false,
	}
}

func (p *SpecificPosition) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "%d %d", p.wheel_id, p.corn)
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