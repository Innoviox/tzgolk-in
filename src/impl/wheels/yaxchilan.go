package wheels

import (
	"fmt"
	. "tzgolkin/engine"
	. "tzgolkin/delta"
)

func Yaxchilan0(g *Game, p *Player) []*Delta {
	return make([]*Delta, 0)
}

func Yaxchilan1(g *Game, p *Player) []*Delta {
	d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
		Resources: [4]int{1 + g.Research.ResourceBonus(p.Color, Wood), 0, 0, 0},
	})

	d.Description = fmt.Sprintf("1 + %d wood", g.Research.ResourceBonus(p.Color, Wood))
	return []*Delta{d}
}

func Yaxchilan2(g *Game, p *Player) []*Delta {
	d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
		Resources: [4]int{0, 1 + g.Research.ResourceBonus(p.Color, Stone), 0, 0},
		Corn: 1,
	})

	d.Description = fmt.Sprintf("1 + %d stone, 1 Corn", g.Research.ResourceBonus(p.Color, Stone))
	return []*Delta{d}
}

func Yaxchilan3(g *Game, p *Player) []*Delta {
	d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
		Resources: [4]int{0, 0, 1 + g.Research.ResourceBonus(p.Color, Gold), 0},
		Corn: 1,
	})

	d.Description = fmt.Sprintf("1 + %d gold, 1 Corn", g.Research.ResourceBonus(p.Color, Gold))
	return []*Delta{d}
}

func Yaxchilan4(g *Game, p *Player) []*Delta {
	d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
		Resources: [4]int{0, 0, 0, 1 + g.Research.ResourceBonus(p.Color, Skull)},
		Corn: 1,
	})

	d.Description = fmt.Sprintf("1 + %d skull", g.Research.ResourceBonus(p.Color, Skull))
	return []*Delta{d}
}

func Yaxchilan5(g *Game, p *Player) []*Delta {
	d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
		Resources: [4]int{0, 1 + g.Research.ResourceBonus(p.Color, Stone), 
						  1 + g.Research.ResourceBonus(p.Color, Gold), 0},
		Corn: 2,
	})

	d.Description = fmt.Sprintf("1 + %d gold, 1 + %d stone, 2 Corn", g.Research.ResourceBonus(p.Color, Gold), g.Research.ResourceBonus(p.Color, Stone))
	return []*Delta{d}
}

func Yaxchilan() []Options {
	return []Options{
		Yaxchilan0,
		Yaxchilan1,
		Yaxchilan2,
		Yaxchilan3,
		Yaxchilan4,
		Yaxchilan5,
	}
}

func MakeYaxchilan() *Wheel {
	return MakeWheel(Yaxchilan(), 1, "Yaxchilan")
}