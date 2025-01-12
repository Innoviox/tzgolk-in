package wheels

import (
	"fmt"
	"strings"
)

type Option func()
type Options func(*Game, *Player) []Option

type Position struct {
	wheel_id int
	corn int
	GetOptions Options
	pData *PalenqueData
}

type PalenqueData struct {
	cornTiles int
	woodTiles int
}

func MakePData() *PalenqueData {
	// todo based on player count
	return &PalenqueData{
		cornTiles: 4,
		woodTiles: 4,
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