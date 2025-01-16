package wheels

import (
	"fmt"
	. "tzgolkin/engine"
)

func Yaxchilan0(g *Game, p *Player) []*Delta {
	return make([]*Delta, 0)
}

func Yaxchilan1(g *Game, p *Player) []*Delta {
	return []*Delta {Option{
		Execute: func(g *Game, p *Player) {
			p.Resources[Wood] += 1 + g.Research.ResourceBonus(p.Color, Wood)
		},
		Description: fmt.Sprintf("1 + %d wood", g.Research.ResourceBonus(p.Color, Wood)),
	}}
}

func Yaxchilan2(g *Game, p *Player) []*Delta {
	return []*Delta {Option{
		Execute: func (g *Game, p *Player) {
			p.Resources[Stone] += 1 + g.Research.ResourceBonus(p.Color, Stone)
			p.Corn += 1
		},
		Description: fmt.Sprintf("1 + %d stone, 1 Corn", g.Research.ResourceBonus(p.Color, Stone)),
	}}
}

func Yaxchilan3(g *Game, p *Player) []*Delta {
	return []*Delta {Option{
		Execute: func (g *Game, p *Player) {
			p.Resources[Gold] += 1 + g.Research.ResourceBonus(p.Color, Gold)
			p.Corn += 1
		},
		Description: fmt.Sprintf("1 + %d gold, 1 Corn", g.Research.ResourceBonus(p.Color, Gold)),
	}}
}

func Yaxchilan4(g *Game, p *Player) []*Delta {
	return []*Delta {Option{
		Execute: func (g *Game, p *Player) {
			p.Resources[Skull] += 1 + g.Research.ResourceBonus(p.Color, Skull)
		},
		Description: fmt.Sprintf("1 + %d skull", g.Research.ResourceBonus(p.Color, Skull)),
	}}
}

func Yaxchilan5(g *Game, p *Player) []*Delta {
	return []*Delta {Option{
		Execute: func (g *Game, p *Player) {
			p.Resources[Gold] += 1 + g.Research.ResourceBonus(p.Color, Gold)
			p.Resources[Stone] += 1 + g.Research.ResourceBonus(p.Color, Stone)
			p.Corn += 2
		},
		Description: fmt.Sprintf("1 + %d gold, 1 + %d stone, 2 Corn", g.Research.ResourceBonus(p.Color, Gold), g.Research.ResourceBonus(p.Color, Stone)),
	}}
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