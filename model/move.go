package model 

import (
	"fmt"
	"strings"
)

type Move struct {
	Placing bool
	Workers []int

	Positions []*SpecificPosition

	Corn int

	Execute func()
}

// -- MARK -- Basic methods
func (m *Move) String() string {
	var br strings.Builder

	if m.Placing {
		fmt.Fprintf(&br, "%s ", "Place")
	} else {
		fmt.Fprintf(&br, "%s ", "Retrieve")
	}

	for i := 0; i < len(m.Workers); i++ {
		fmt.Fprintf(&br, "(%d, %s) ", m.Workers[i], m.Positions[i].String())
	}


	return br.String()
}

// -- MARK -- Unique methods
func (m *Move) Retrieve(worker int, position *SpecificPosition, CornBack int) Move {
	return Move {
		Placing: m.Placing,
		Workers: append(m.Workers, worker),
		Positions: append(m.Positions, position),
		Corn: m.Corn + CornBack,
	}
}

func (m *Move) Place(worker int, position *SpecificPosition) Move {
	return Move {
		Placing: m.Placing,
		Workers: append(m.Workers, worker),
		Positions: append(m.Positions, position),
		Corn: m.Corn + len(m.Workers) + position.Corn,
	}
}