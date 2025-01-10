package model

import (
	"fmt"
	"strings"
)

type Position struct {
	wheel_id int
	corn int
	Execute func(*Game, Color, int) // todo color type
	decisions int
}

func (p *Position) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "%d %d", p.wheel_id, p.corn)
	return br.String()
}