package model

func Uxmal0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

func Uxmal1(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	if p.corn > 3 {
		for i := 0; i < 3; i++ {
			options = append(options, func() {
				p.corn -= 3
				g.temples.Step(p.color, i)
			})
		}
	}

	return options
}

func Uxmal2(g *Game, p *Player) []Option {
	// todo
}

func Uxmal3(g *Game, p *Player) []Option {
	return []Option{
		func() {
			g.UnlockWorker(p.color)
		},
	}
}

func Uxmal4(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	for _, b := range g.currentBuildings {
		cost := b.CornCost(g, p) // todo how does research interact with this
		if p.corn >= cost {
			for _, effect := range b.GetEffects(g, p) {
				options = append(options, func() {
					p.corn -= cost
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
}

func Uxmal5(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	if p.corn == 0 {
		return options
	}

	allOptions := make([]Option, 0)
	for _, option := range Yaxchilan() {
		allOptions = append(allOptions, option(g, p)...)
	}
	for _, option := range Tikal() {
		allOptions = append(allOptions, option(g, p)...)
	}
	for _, option := range []Options{ Uxmal1, Uxmal2, Uxmal3, Uxmal4 } {
		allOptions = append(allOptions, option(g, p)...)
	}


	for _, option := range allOptions {
		options = append(options, func() {
			p.corn -= 1
			option()
		})
	}

	return options
}

func Uxmal() []Options {
	return []Options{ Uxmal0, Uxmal1, Uxmal2, Uxmal3, Uxmal4, Uxmal5 }
}

func MakeUxmal() *Wheel {
	return MakeWheel(Uxmal(), 3, "Uxmal")
}