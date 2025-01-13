package model

import (
	"fmt"
)

type Building struct {
	Id int
	Cost [4]int
	GetEffects Options
	Color Color
}


func (b *Building) GetCosts(game *Game, player *Player) [][4]int {
	options := make([][4]int, 0)

	if player.CanPay(b.Cost) {
		options = append(options, b.Cost)
	}

	if game.Research.Builder(player.Color) {
		for i := 0; i < 4; i++ {
			if b.Cost[i] > 0 {
				cost := b.Cost
				cost[i] -= 1
				options = append(options, cost)
			}
		}
	}

	return options
}

func (b *Building) CornCost(game *Game, player *Player) int {
	cost := 0

	for i := 0; i < 4; i++ {
		cost += b.Cost[i] * 2
	}

	// todo does research interact with this
	if game.Research.Builder(player.Color) {
		cost -= 2
	}

	return cost
}

func (g *Game) GetBuildingOptions(p *Player, exclude int, useResearch bool) []Option {
	options := make([]Option, 0)

	for _, b := range g.CurrentBuildings {
		if b.Id == exclude {
			continue
		}

		costs := b.GetCosts(g, p)
		for _, cost := range costs {
			for _, effect := range b.GetEffects(g, p) {
				options = append(options, Option{
					Execute: func(g *Game, p *Player) {
						for i := 0; i < 4; i++ {
							p.Resources[i] -= cost[i]
						}

						effect.Execute(g, p)

						if useResearch {
							g.Research.Built(p)
						}

						p.Buildings = append(p.Buildings, b)

						// g.RemoveBuilding(b)
					},
					Description: fmt.Sprintf("[build %d] pay %s, %s +%s", b.Id, CostString(cost), effect.Description, g.Research.BuiltString(p)),
					BuildingNum: b.Id,
				})
			}
		}
	}

	return options
}

func (g *Game) GetMonumentOptions(p *Player) []Option {
	options := make([]Option, 0)

	for _, m := range g.CurrentMonuments {
		if p.CanPay(m.Cost) {
			options = append(options, Option{
				Execute: func(g *Game, p *Player) {
					for i := 0; i < 4; i++ {
						p.Resources[i] -= m.Cost[i]
					}

					p.Monuments = append(p.Monuments, m)

					// g.RemoveMonument(m)
				},
				Description: fmt.Sprintf("[build %d] pay %s, get monument %d", m.Id, CostString(m.Cost), m.Id),
			})
		}
	}

	return options
}
