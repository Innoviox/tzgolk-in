package model

func Tikal0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

func Tikal1(g *Game, p *Player) []Option {
	return g.research.GetOptions(p, 1)
}

func Tikal2(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	for _, b := range g.currentBuildings {
		costs := b.GetCosts(g, p)
		for _, cost := range costs {
			for _, effect := range b.GetEffects(g, p) {
				options = append(options, func() {
					for i := 0; i < 4; i++ {
						p.resources[i] -= cost[i]
					}

					effect()

					if g.research.HasLevel(p.color, Construction, 1) {
						p.corn += 1
					}
					if g.research.HasLevel(p.color, Construction, 2) {
						p.points += 2
					}
					// todo building colors?
				})
			}
		}
	}

	return options
}

func Tikal3(g *Game, p *Player) []Option {
	return g.research.GetOptions(p, 2)
}

func Tikal4(g *Game, p *Player) []Option {

}

func Tikal5(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	for i := 0; i < 3; i++ {
		if p.resources[i] > 0 {
			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					if (j == k) {
						continue
					}
					options = append(options, func() {
						p.resources[i] -= 1
						g.temples.Step(p.color, j)
						g.temples.Step(p.color, k)
					})
				}
			}
		}
	}
}