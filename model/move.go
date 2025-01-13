package model 

import (
	"fmt"
	"strings"
	"github.com/innoviox/tzgolkin/model/types"
)

func (m *Move) Retrieve(worker int, position *SpecificPosition, cornBack int) Move {
	return Move {
		placing: m.placing,
		workers: append(m.workers, worker),
		positions: append(m.positions, position),
		corn: m.corn + cornBack,
	}
}

func (m *Move) Place(worker int, position *SpecificPosition) Move {
	return Move {
		placing: m.placing,
		workers: append(m.workers, worker),
		positions: append(m.positions, position),
		corn: m.corn + len(m.workers) + position.corn,
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