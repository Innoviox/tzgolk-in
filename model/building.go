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