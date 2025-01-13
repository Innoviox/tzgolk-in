package wheels

import (
	"fmt"
	. "tzgolkin/model"
)

func Tikal0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

func Tikal1(g *Game, p *Player) []Option {
	return g.Research.GetOptions(g, p, 1, false)
}

func Tikal2(g *Game, p *Player) []Option {
	return g.GetBuildingOptions(p, -1, true)
}

func Tikal3(g *Game, p *Player) []Option {
	return g.Research.GetOptions(g, p, 2, false)
}

func Tikal4(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	for _, o := range g.GetBuildingOptions(p, -1, true) {
		options = append(options, o)

		for _, o2 := range g.GetBuildingOptions(p, o.BuildingNum, false) {
			options = append(options, Option{
				Execute: func(g *Game, p *Player) {
					o.Execute(g, p)
					o2.Execute(g, p)
				},
				Description: fmt.Sprintf("%s, %s [no res]", o.Description, o2.Description),
			})
		}
	}

	options = append(options, g.GetMonumentOptions(p)...)

	return options
}

func Tikal5(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	for i := 0; i < 3; i++ {
		if p.Resources[i] > 0 {
			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					if (j == k) {
						continue
					}
					options = append(options, Option{
						Execute: func(g *Game, p *Player) {
							p.Resources[i] -= 1
							g.Temples.Step(p, j, 1)
							g.Temples.Step(p, k, 1)
						},
						Description: fmt.Sprintf("pay 1 %s, 1 %sT, 1 %sT", string(ResourceDebug[i]), string(TempleDebug[j]), string(TempleDebug[k])),
					})
				}
			}
		}
	}

	return options
}

func Tikal() []Options {
	return []Options{ Tikal0, Tikal1, Tikal2, Tikal3, Tikal4, Tikal5, }
}

func MakeTikal() *Wheel {
	return MakeWheel(Tikal(), 2, "Tikal")
}