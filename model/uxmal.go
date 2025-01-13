package model

import (
	"fmt"
)

func Uxmal0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

func Uxmal1(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	if p.corn > 3 {
		options = append(options, g.temples.GainTempleStep(p, Option{
			Execute: func() {
				p.corn -= 3
			},
			description: "pay 3 corn",
		}, 1)...,)
	}

	return options
}

func Uxmal2(g *Game, p *Player) []Option {
	// todo
	return make([]Option, 0)
}

func Uxmal3(g *Game, p *Player) []Option {
	return []Option{Option{
		Execute: func() {
			g.UnlockWorker(p.color)
		},
		description: "unlock worker",
	}}
}

func Uxmal4(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	for _, b := range g.currentBuildings {
		cost := b.CornCost(g, p) 
		if p.corn >= cost {
			for _, effect := range b.GetEffects(g, p) {
				options = append(options, Option{
					Execute: func() {
						p.corn -= cost
						effect.Execute()

						g.research.Built(p)
						// todo building colors?
					},
					description: fmt.Sprintf("[build] pay %d corn, %s", cost, effect.description),
				})
			}
		}
	}

	return options
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
		options = append(options, Option{
			Execute: func() {
				p.corn -= 1
				option.Execute()
			},
			description: fmt.Sprintf("[mirror] pay 1 corn, %s", option.description),
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