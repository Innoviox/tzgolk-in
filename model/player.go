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

	fmt.Fprintf(&br, "----Player %s------------\n", p.color.String())
	fmt.Fprintf(&br, "| Resources: ")
	for i := 0; i < 4; i++ {
		fmt.Fprintf(&br, "%d%s ", p.resources[i], string(ResourceDebug[i]))
	}
	fmt.Fprintf(&br, "\n| Corn: %d\n", p.corn)
	fmt.Fprintf(&br, "| Points: %d\n", p.points)
	fmt.Fprintf(&br, "| Corn Tiles: %d\n", p.cornTiles)
	fmt.Fprintf(&br, "| Wood Tiles: %d\n", p.woodTiles)
	fmt.Fprintf(&br, "| Free Workers: %d\n", p.freeWorkers)
	fmt.Fprintf(&br, "| Worker Deduction: %d\n", p.workerDeduction)
	fmt.Fprintf(&br, "| Light Side: %t\n", p.lightSide)
	fmt.Fprintf(&br, "------------------------\n")


	return br.String()
}

func (p *Player) Clone() *Player {
	return &Player {
		resources: p.resources,
		corn: p.corn,
		color: p.color,
		points: p.points,
		cornTiles: p.cornTiles,
		woodTiles: p.woodTiles,
		freeWorkers: p.freeWorkers,
		workerDeduction: p.workerDeduction,
		lightSide: p.lightSide,
		buildings: p.buildings,
		monuments: p.monuments,
	}
}