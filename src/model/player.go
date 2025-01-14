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

// -- MARK -- Basic methods
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

func (p *Player) String(g *Game) string {
	var br strings.Builder

	nWorkers := 0
	for _, w := range g.Workers {
		if w.Color == p.Color && (w.Available || w.Wheel_id > 0) {
			nWorkers++
		}
	}

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
	fmt.Fprintf(&br, "| Buildings: %d\n", len(p.Buildings))
	fmt.Fprintf(&br, "| Monuments: %d\n", len(p.Monuments))
	fmt.Fprintf(&br, "| Workers: %d\n", nWorkers)
	fmt.Fprintf(&br, "------------------------\n")


	return br.String()
}

// -- MARK -- Unique methods
func (p *Player) CanPay(cost [4]int) bool {
	for i := 0; i < 4; i++ {
		if p.Resources[i] < cost[i] {
			return false
		}
	}
	
	return true
}

func (p *Player) TotalCorn() int {
	Corn := p.Corn
	Corn += 2 * p.Resources[Wood]
	Corn += 3 * p.Resources[Stone]
	Corn += 4 * p.Resources[Gold]
	
	return Corn
}

func (p *Player) Evaluate(g *Game) float64 {
	points := float64(p.Points)
	points += float64(p.TotalCorn()) / 2
	points += float64(p.Resources[Skull]) * 3

	for _, m := range p.Monuments {
		points += float64(m.GetPoints(g, p))
	}

	for _, w := range g.Workers {
		if w.Color == p.Color {
			if (w.Available) {
				points += 1
			} else if w.Wheel_id != -1 {
				points += float64(w.Position) / 2
			}
		}
	}

	return points
}