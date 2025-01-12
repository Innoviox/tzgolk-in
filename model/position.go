package model

import (
	"fmt"
	"strings"
)

type Position struct {
	wheel_id int
	corn int
	GetOptions Options
	pData *PalenqueData
	cData *ChichenData
}

type PalenqueData struct {
	cornTiles int
	woodTiles int
}

type ChichenData struct {
	full bool
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

func MakeCData() *ChichenData {
	return &ChichenData {
		full: false,
	}
}

func (p *Position) String() string {
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