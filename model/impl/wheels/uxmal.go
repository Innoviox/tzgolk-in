package wheels

import (
	"fmt"
	. "tzgolkin/model"
)

func Uxmal0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

func Uxmal1(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	if p.Corn > 3 {
		options = append(options, g.Temples.GainTempleStep(p, Option{
			Execute: func(g *Game, p *Player) {
				p.Corn -= 3
			},
			Description: "pay 3 Corn",
		}, 1)...,)
	}

	return options
}

func Uxmal2(g *Game, p *Player) []Option {
	// Corn := TotalCorn(p)

	// CornOptions := GenerateCornExchanges(Corn, []CornOption{CornOption{
	// 	Corn: p.Corn,
	// 	resources: p.Resources,
	// }})

	options := make([]Option, 0)
	// for _, o := range CornOptions {
	// 	options = append(options, Option{
	// 		Execute: func(g *Game, p *Player) {
	// 			p.Corn = o.Corn
	// 			p.Resources = o.Resources
	// 		},
	// 		Description: fmt.Sprintf("exchange to %d Corn, %v", o.Corn, o.Resources),
	// 	})
	// }

	return options
}

type CornOption struct {
	Corn int
	Resources [4]int
}

func GenerateCornExchanges(Corn int, base []CornOption) []CornOption {
	if (Corn < 2) {
		return base
	}

	options := make([]CornOption, 0)
	
	if (Corn >= 2) {
		for _, o := range GenerateCornExchanges(Corn - 2, base) {
			options = append(options, CornOption{
				Corn: o.Corn - 2,
				Resources: [4]int{ o.Resources[0] + 1, o.Resources[1], o.Resources[2], o.Resources[3] },
			})
		}
	}

	if (Corn >= 3) {
		for _, o := range GenerateCornExchanges(Corn - 3, base) {
			options = append(options, CornOption{
				Corn: o.Corn - 3,
				Resources: [4]int{ o.Resources[0], o.Resources[1] + 1, o.Resources[2], o.Resources[3] },
			})
		}
	}

	if (Corn >= 4) {
		for _, o := range GenerateCornExchanges(Corn - 4, base) {
			options = append(options, CornOption{
				Corn: o.Corn - 4,
				Resources: [4]int{ o.Resources[0], o.Resources[1], o.Resources[2] + 1, o.Resources[3] },
			})
		}
	}

	return options
}

func Uxmal3(g *Game, p *Player) []Option {
	return []Option{Option{
		Execute: func(g *Game, p *Player) {
			g.UnlockWorker(p.Color)
		},
		Description: "unlock worker",
	}}
}

func Uxmal4(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	for _, b := range g.CurrentBuildings {
		cost := b.CornCost(g, p) 
		if p.Corn >= cost {
			for _, effect := range b.GetEffects(g, p) {
				options = append(options, Option{
					Execute: func(g *Game, p *Player) {
						p.Corn -= cost
						effect.Execute(g, p)

						g.Research.Built(p)
						
						p.Buildings = append(p.Buildings, b)
						// g.RemoveBuilding(b)
					},
					Description: fmt.Sprintf("[build %d] pay %d Corn, %s", b.Id, cost, effect.Description),
				})
			}
		}
	}

	return options
}

func Uxmal5(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	if p.Corn == 0 {
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
			Execute: func(g *Game, p *Player) {
				p.Corn -= 1
				option.Execute(g, p)
			},
			Description: fmt.Sprintf("[mirror] pay 1 Corn, %s", option.Description),
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