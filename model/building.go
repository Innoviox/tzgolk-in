package model

import (
	"fmt"
)

type Building struct {
	id int
	cost [4]int
	GetEffects Options
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

func (g *Game) GetBuildingOptions(p *Player, exclude int) []Option {
	options := make([]Option, 0)

	for _, b := range g.currentBuildings {
		if b.id == exclude {
			continue
		}

		costs := b.GetCosts(g, p)
		for _, cost := range costs {
			for _, effect := range b.GetEffects(g, p) {
				options = append(options, Option{
					Execute: func() {
						for i := 0; i < 4; i++ {
							p.resources[i] -= cost[i]
						}

						effect.Execute()

						g.research.Built(p)
						// todo building colors?
					},
					description: fmt.Sprintf("[build] pay %s, %s", CostString(cost), effect.description),
				})
			}
		}
	}

	return options
}