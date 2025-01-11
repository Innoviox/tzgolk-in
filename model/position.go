package model

import (
	"fmt"
	"strings"
)

type Option func(*Game, Color)

type Position struct {
	wheel_id int
	corn int
	options []Option // todo color type
}

func (p *Position) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "%d %d", p.wheel_id, p.corn)
	return br.String()
}

func PlayerOption(f func (*Player)) Option {
	return func(g *Game, c Color) {
		p := g.GetPlayerByColor(c)
		f(p)
	}
}