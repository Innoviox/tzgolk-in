package engine 

import (
	"fmt"
	"strings"
)

type Move struct {
	Placing bool
	Workers []int

	Positions []*SpecificPosition

	Corn int

	Begged int
	Player Color

	Execute func()
}

// -- MARK -- Basic methods
func MakeEmptyRetrievalMove(c Color) Move {
	return Move {
		Placing: false,
		Workers: make([]int, 0),
		Positions: make([]*SpecificPosition, 0),
		Corn: 0,
		Begged: -1, // there's no blue temple
		Player: c,
	}
}

func MakeEmptyPlacementMove(c Color) Move {
	return Move {
		Placing: true,
		Workers: make([]int, 0),
		Positions: make([]*SpecificPosition, 0),
		Corn: 0,
		Begged: -1,
		Player: c,
	}
}

func (m *Move) String() string {
	var br strings.Builder

	if m.Begged != -1 {
		fmt.Fprintf(&br, "[Beg %d] ", m.Begged)
	}

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
		Begged: m.Begged,
		Player: m.Player,
	}
}

func (m *Move) Place(worker int, position *SpecificPosition) Move {
	return Move {
		Placing: m.Placing,
		Workers: append(m.Workers, worker),
		Positions: append(m.Positions, position),
		Corn: m.Corn + len(m.Workers) + position.Corn,
		Begged: m.Begged,
		Player: m.Player,
	}
}

func (m *Move) Beg(temple int) Move {
	return Move {
		Placing: m.Placing,
		Workers: m.Workers,
		Positions: m.Positions,
		Corn: m.Corn,
		Begged: temple,
		Player: m.Player,
	}
}