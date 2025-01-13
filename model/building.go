package model

import (
	"fmt"
)

type Building struct {
	id int
	cost [4]int
	GetEffects Options
	color Color
}

func MakeBuilding(id int, cost [4]int, getEffects Options, color Color) Building {
	return Building {
		id: id,
		cost: cost,
		GetEffects: getEffects,
		color: color,
	}
}

func (b *Building) CanBuild(player *Player) bool {
	for i := 0; i < 4; i++ {
		if player.resources[i] < b.cost[i] {
			return false
		}
	}

	return true
}

func (b *Building) GetCosts(game *Game, player *Player) [][4]int {
	options := make([][4]int, 0)

	if b.CanBuild(player) {
		options = append(options, b.cost)
	}

	if game.research.Builder(player.color) {
		for i := 0; i < 4; i++ {
			if b.cost[i] > 0 {
				cost := b.cost
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
		cost += b.cost[i] * 2
	}

	// todo does research interact with this
	if game.research.Builder(player.color) {
		cost -= 2
	}

	return cost
}

func (g *Game) GetBuildingOptions(p *Player, exclude int, useResearch bool) []Option {
	options := make([]Option, 0)

	for _, b := range g.currentBuildings {
		if b.id == exclude {
			continue
		}

		costs := b.GetCosts(g, p)
		for _, cost := range costs {
			for _, effect := range b.GetEffects(g, p) {
				options = append(options, Option{
					Execute: func(g *Game, p *Player) {
						for i := 0; i < 4; i++ {
							p.resources[i] -= cost[i]
						}

						effect.Execute(g, p)

						if useResearch {
							g.research.Built(p)
						}

						p.buildings = append(p.buildings, b)

						// g.RemoveBuilding(b)
					},
					description: fmt.Sprintf("[build %d] pay %s, %s +%s", b.id, CostString(cost), effect.description, g.research.BuiltString(p)),
					buildingNum: b.id,
				})
			}
		}
	}

	return options
}

func (g *Game) GetMonumentOptions(p *Player) []Option {
	options := make([]Option, 0)

	for _, m := range g.currentMonuments {
		if m.CanBuild(p) {
			options = append(options, Option{
				Execute: func(g *Game, p *Player) {
					for i := 0; i < 4; i++ {
						p.resources[i] -= m.cost[i]
					}

					p.monuments = append(p.monuments, m)

					// g.RemoveMonument(m)
				},
				description: fmt.Sprintf("[build %d] pay %s, get monument %d", m.id, CostString(m.cost), m.id),
			})
		}
	}

	return options
}
