package model

import (
	"fmt"
	"strings"
	"github.com/innoviox/tzgolkin/model/types"
)

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

func (pd *PalenqueData) HasCornShowing() bool {
	return pd.cornTiles > pd.woodTiles
}

func MakeCData() *ChichenData {
	return &ChichenData {
		full: false,
	}
}

func (p *SpecificPosition) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "%d %d", p.wheel_id, p.corn)
	return br.String()
}