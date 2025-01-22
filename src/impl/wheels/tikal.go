package wheels

import (
	"fmt"
	. "tzgolkin/engine"
)

func Tikal0(g *Game, p *Player) []*Delta {
	return make([]*Delta, 0)
}

func Tikal1(g *Game, p *Player) []*Delta {
	return SkipWrapper(g.Research.GetOptions(g, p, 1, false))
}

func Tikal2(g *Game, p *Player) []*Delta {
	return SkipWrapper(g.GetBuildingOptions(p, -1, true))
}

func Tikal3(g *Game, p *Player) []*Delta {
	return SkipWrapper(g.Research.GetOptions(g, p, 2, false))
}

func Tikal4(g *Game, p *Player) []*Delta {
	options := make([]*Delta, 0)

	for _, o := range g.GetBuildingOptions(p, -1, true) {
		options = append(options, o)

		o1 := o//.Clone()

		// g2 := g.Clone()
		g.AddDelta(o1, 1, false)

		for _, o2 := range g.GetBuildingOptions(p, o.BuildingNum, false) {
			d := Combine(o, o2)
			d.Description = fmt.Sprintf("%s, %s [no res]", o.Description, o2.Description)
			options = append(options, d)
		}

		g.AddDelta(o1, -1, true)
		// if !g.Exact(g2) {
		// 	fmt.Println("PLATO ERROR %")
		// 	fmt.Println(o1)
		// 	fmt.Println([]int{}[1])
		// }
	}

	options = append(options, g.GetMonumentOptions(p)...)

	return SkipWrapper(options)
}

func Tikal5(g *Game, p *Player) []*Delta {
	options := make([]*Delta, 0)

	for i := 0; i < 3; i++ {
		if p.Resources[i] > 0 {
			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					if (j == k) {
						continue
					}

					r := [4]int{}
					r[i] -= 1

					d := PlayerDeltaWrapper(p.Color, PlayerDelta{
						Resources: r,
					})

					d.Add(g.Temples.Step(p, j, 1), true)
					d.Add(g.Temples.Step(p, k, 1), true)

					d.Description = fmt.Sprintf("pay 1 %s, 1 %sT, 1 %sT", string(ResourceDebug[i]), string(TempleDebug[j]), string(TempleDebug[k]))

					options = append(options, d)
				}
			}
		}
	}

	return SkipWrapper(options)
}

func Tikal() []Options {
	return []Options{ Tikal0, Tikal1, Tikal2, Tikal3, Tikal4, Tikal5, }
}

func MakeTikal() *Wheel {
	return MakeWheel(Tikal(), 2, "Tikal")
}