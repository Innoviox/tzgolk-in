package model 

import (
	"fmt"
	"strings"
)

type Move struct {
	placing bool
	workers []int

	positions []*Position

	corn int
}

func (m *Move) Retrieve(worker int) Move {
	return Move {
		placing: m.placing,
		workers: append(m.workers, worker),
		positions: m.positions,
		corn: m.corn,
	}
}

func (m *Move) Place(worker int, position *Position) Move {
	return Move {
		placing: m.placing,
		workers: append(m.workers, worker),
		positions: append(m.positions, position),
		corn: m.corn,
	}
}

func (m *Move) String() string {
	var br strings.Builder

	if m.placing {
		fmt.Fprintf(&br, "%s ", "Place")
	} else {
		fmt.Fprintf(&br, "%s ", "Retrieve")
	}

	for i := 0; i < len(m.workers); i++ {
		fmt.Fprintf(&br, "(%d, %s) ", m.workers[i], m.positions[i].String())
	}


	return br.String()
}