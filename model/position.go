package model

import (
	"fmt"
	"strings"
)

type Position struct {
	wheel_id int
	position_num int 
	corn int
	Execute func(*Game, string) // todo color type
}

func (p *Position) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "%d %d", p.wheel_id, p.position_num)
	return br.String()
}