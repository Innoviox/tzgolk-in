package model

import (
	"fmt"
	"strings"
)

type Position struct {
	Wheel_id int
	Corn int
	GetOptions Options
	PData *PalenqueData
	CData *ChichenData
}

func (p *Position) Clone() *Position {
	var newPData *PalenqueData
	if p.PData != nil {
		newPData = p.PData.Clone()
	}

	var newCData *ChichenData
	if p.CData != nil {
		newCData = p.CData.Clone()
	}

	return &Position {
		Wheel_id: p.Wheel_id,
		Corn: p.Corn,
		GetOptions: p.GetOptions,
		PData: newPData,
		CData: newCData,
	}
}

func (p *Position) Copy(other *Position) {
	var newPData *PalenqueData
	if other.PData != nil {
		newPData = other.PData.Clone()
	}

	var newCData *ChichenData
	if other.CData != nil {
		newCData = other.CData.Clone()
	}

	p.Wheel_id = other.Wheel_id
	p.Corn = other.Corn
	p.GetOptions = other.GetOptions
	p.PData = newPData
	p.CData = newCData
}

type SpecificPosition struct {
	Wheel_id int
	Corn int
	Execute Option
	FirstPlayer bool
}

func (p *SpecificPosition) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "%d %d", p.Wheel_id, p.Corn)
	return br.String()
}

type PalenqueData struct {
	CornTiles int
	WoodTiles int
}

func MakePData(hasWood bool) *PalenqueData {
	// todo based on player count
	if (hasWood) {
		return &PalenqueData{
			CornTiles: 4,
			WoodTiles: 4,
		}
	}

	return &PalenqueData{
		CornTiles: 4,
		WoodTiles: 0,
	}
}

func (pd *PalenqueData) Clone() *PalenqueData {
	return &PalenqueData {
		CornTiles: pd.CornTiles,
		WoodTiles: pd.WoodTiles,
	}
}

func (pd *PalenqueData) HasCornShowing() bool {
	return pd.CornTiles > pd.WoodTiles
}


type ChichenData struct {
	Full bool
}

func MakeCData() *ChichenData {
	return &ChichenData {
		Full: false,
	}
}

func (cd *ChichenData) Clone() *ChichenData {
	return &ChichenData {
		Full: cd.Full,
	}
}