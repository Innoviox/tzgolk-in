package wheels

import (
	"fmt"
	. "tzgolkin/engine"
	. "tzgolkin/delta"
)

func Uxmal0(g *Game, p *Player) []*Delta {
	return make([]*Delta, 0)
}

func Uxmal1(g *Game, p *Player) []*Delta {
	options := make([]*Delta, 0)

	if p.Corn >= 3 {
		d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
			Corn: -3,
		})
		d.Description = "pay 3 Corn"
		options = append(options, g.Temples.GainTempleStep(p, d, 1)...)
	}

	return SkipWrapper(options)
}

func Uxmal2(g *Game, p *Player) []*Delta {
	// Corn := TotalCorn(p)

	// CornOptions := GenerateCornExchanges(Corn, []CornOption{CornOption{
	// 	Corn: p.Corn,
	// 	resources: p.Resources,
	// }})

	options := make([]*Delta, 0)
	// for _, o := range CornOptions {
	// 	options = append(options, Option{
	// 		Execute: func(g *Game, p *Player) {
	// 			p.Corn = o.Corn
	// 			p.Resources = o.Resources
	// 		},
	// 		Description: fmt.Sprintf("exchange to %d Corn, %v", o.Corn, o.Resources),
	// 	})
	// }

	return SkipWrapper(options)
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

func Uxmal3(g *Game, p *Player) []*Delta {
	return []*Delta{g.UnlockWorker(p.Color)}
}

func Uxmal4(g *Game, p *Player) []*Delta {
	options := make([]*Delta, 0)

	for k, v := range g.CurrentBuildings {
		if v != 1 {
			continue
		}
		b := g.Buildings[k]

		cost := b.CornCost(g, p) 
		if p.Corn >= cost {
			for _, effect := range b.GetEffects(g, p) {
				d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
					Corn: -cost,
					Buildings: map[int]int{
						b.Id: 1,
					},
				})

				d.Add(effect)
				d.Add(g.Research.Built(p))

				d.Buildings = map[int]int{
					b.Id: 1,
				}

				d.Description = fmt.Sprintf("[build %d] pay %d Corn, %s", b.Id, cost, effect.Description)
				options = append(options, d)
			}
		}
	}

	return SkipWrapper(options)
}

func Uxmal5(g *Game, p *Player) []*Delta {
	options := make([]*Delta, 0)

	if p.Corn == 0 {
		return Skip()
	}

	allOptions := make([]*Delta, 0)
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
		options = append(options, Combine(option, PlayerDeltaWrapper(int(p.Color), PlayerDelta{
			Corn: -1,
		})))
	}

	return options
}

func Uxmal() []Options {
	return []Options{ Uxmal0, Uxmal1, Uxmal2, Uxmal3, Uxmal4, Uxmal5 }
}

func MakeUxmal() *Wheel {
	return MakeWheel(Uxmal(), 3, "Uxmal")
}