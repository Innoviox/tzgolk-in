package model

import (
	"fmt"
	"strings"
)

type Player struct {
	Resources [4]int
	Corn int 
	Color Color

	Points int
	CornTiles int
	WoodTiles int

	FreeWorkers int
	WorkerDeduction int

	LightSide bool

	Buildings []Building
	Monuments []Monument
}


func (p *Player) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "----Player %s------------\n", p.Color.String())
	fmt.Fprintf(&br, "| Resources: ")
	for i := 0; i < 4; i++ {
		fmt.Fprintf(&br, "%d%s ", p.Resources[i], string(ResourceDebug[i]))
	}
	fmt.Fprintf(&br, "\n| Corn: %d\n", p.Corn)
	fmt.Fprintf(&br, "| Points: %d\n", p.Points)
	fmt.Fprintf(&br, "| Corn Tiles: %d\n", p.CornTiles)
	fmt.Fprintf(&br, "| Wood Tiles: %d\n", p.WoodTiles)
	fmt.Fprintf(&br, "| Free Workers: %d\n", p.FreeWorkers)
	fmt.Fprintf(&br, "| Worker Deduction: %d\n", p.WorkerDeduction)
	fmt.Fprintf(&br, "| Light Side: %t\n", p.LightSide)
	fmt.Fprintf(&br, "------------------------\n")


	return br.String()
}

func (p *Player) Clone() *Player {
	return &Player {
		Resources: p.Resources,
		Corn: p.Corn,
		Color: p.Color,
		Points: p.Points,
		CornTiles: p.CornTiles,
		WoodTiles: p.WoodTiles,
		FreeWorkers: p.FreeWorkers,
		WorkerDeduction: p.WorkerDeduction,
		LightSide: p.LightSide,
		Buildings: p.Buildings,
		Monuments: p.Monuments,
	}
}