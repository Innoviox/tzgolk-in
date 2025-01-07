package model

import (
	"fmt"
	"strings"
)

// // https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
// func remove(s []int, i int) []int {
//     s[i] = s[len(s)-1]
//     return s[:len(s)-1]
// }

type Resource int

const (
	Wood Resource = iota
	Stone
	Gold
	Skull
)

type Position struct {
	wheel_id int
	position_num int 
}

func (p *Position) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "%d %d", p.wheel_id, p.position_num)
	return br.String()
}

func MakeEmptyRetrievalMove() Move {
	return Move {
		placing: false,
		workers: make([]int, 0),
		positions: make([]*Position, 0),
		corn: 0,
	}
}

func MakeEmptyPlacementMove() Move {
	return Move {
		placing: true,
		workers: make([]int, 0),
		positions: make([]*Position, 0),
		corn: 0,
	}
}