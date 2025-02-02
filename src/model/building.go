package model

import (
	"fmt"
	"strings"
)

type Building struct {
	Id int
	Cost [4]int
	GetEffects Options
	Color Color
}

func (b *Building) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "------\n")
	fmt.Fprintf(&br, "|")
	n := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < b.Cost[i]; j++ {
			fmt.Fprintf(&br, "%s", string(ResourceDebug[i]))
			n++
		}
	}
	for ; n < 4; n++ {
		fmt.Fprintf(&br, " ")
	}
	fmt.Fprintf(&br, "|\n")
	fmt.Fprintf(&br, "|%-4d|\n", b.Id)
	fmt.Fprintf(&br, "------\n")

	return br.String()
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
				if player.CanPay(cost) {
					options = append(options, cost)
				}
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