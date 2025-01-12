package model

// import (
// 	"tzgolkin/model/wheels"
// )

type Building struct {
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