package engine

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

	Buildings map[int]int
	Monuments map[int]int
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

func (p *Player) Copy(other *Player) {
	p.Resources = [4]int{other.Resources[0], other.Resources[1], other.Resources[2], other.Resources[3]}
	
	p.Corn = other.Corn
	p.Color = other.Color
	p.Points = other.Points
	p.CornTiles = other.CornTiles
	p.WoodTiles = other.WoodTiles
	p.FreeWorkers = other.FreeWorkers
	p.WorkerDeduction = other.WorkerDeduction
	p.LightSide = other.LightSide

	p.Buildings = CopyMap(other.Buildings)
	p.Monuments = CopyMap(other.Monuments)
}

func (p *Player) AddDelta(delta PlayerDelta, mul int) {
	for i := 0; i < 4; i++ {
		p.Resources[i] += delta.Resources[i] * mul
	}

	p.Corn += delta.Corn * mul
	p.Points += delta.Points * mul
	p.CornTiles += delta.CornTiles * mul
	p.WoodTiles += delta.WoodTiles * mul
	p.FreeWorkers += delta.FreeWorkers * mul
	p.WorkerDeduction += delta.WorkerDeduction * mul

	p.LightSide = Bool(delta.LightSide, mul, p.LightSide)
	for k, v := range delta.Buildings {
		p.Buildings[k] += v * mul
	}
	for k, v := range delta.Monuments {
		p.Monuments[k] += v * mul
	}
	// todo buildings & monuments
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
	fmt.Fprintf(&br, "| Buildings: %d\n", CountValues(p.Buildings, 1))
	fmt.Fprintf(&br, "| Monuments: %d\n", CountValues(p.Monuments, 1))
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
	points := float64(p.Points) * 2
	points += float64(p.TotalCorn())
	points += float64(p.Resources[Skull]) * 3

	for k, v := range p.Monuments {
		if v == 1 {
			points += float64(g.Monuments[k].GetPoints(g, p))
		}
	}

	for _, w := range g.Workers {
		if w.Color == p.Color {
			if (w.Available) {
				points += 0.1
			} else if w.Wheel_id != -1 {
				points += float64(w.Position) / 10
			}
		}
	}

	return points
}