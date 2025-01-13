package model

import (
	"fmt"
)

func Yaxchilan0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

func Yaxchilan1(g *Game, p *Player) []Option {
	return []Option {Option{
		Execute: func() {
			p.resources[Wood] += 1 + g.research.ResourceBonus(p.color, Wood)
		},
		description: fmt.Sprintf("1 + %d wood", g.research.ResourceBonus(p.color, Wood)),
	}}
}

func Yaxchilan2(g *Game, p *Player) []Option {
	return []Option {Option{
		Execute: func () {
			p.resources[Stone] += 1 + g.research.ResourceBonus(p.color, Stone)
			p.corn += 1
		},
		description: fmt.Sprintf("1 + %d stone, 1 corn", g.research.ResourceBonus(p.color, Stone)),
	}}
}

func Yaxchilan3(g *Game, p *Player) []Option {
	return []Option {Option{
		Execute: func () {
			p.resources[Gold] += 1 + g.research.ResourceBonus(p.color, Gold)
			p.corn += 1
		},
		description: fmt.Sprintf("1 + %d gold, 1 corn", g.research.ResourceBonus(p.color, Gold)),
	}}
}

func Yaxchilan4(g *Game, p *Player) []Option {
	return []Option {Option{
		Execute: func () {
			p.resources[Skull] += 1 + g.research.ResourceBonus(p.color, Skull)
		},
		description: fmt.Sprintf("1 + %d skull", g.research.ResourceBonus(p.color, Skull)),
	}}
}

func Yaxchilan5(g *Game, p *Player) []Option {
	return []Option {Option{
		Execute: func () {
			p.resources[Gold] += 1 + g.research.ResourceBonus(p.color, Gold)
			p.resources[Stone] += 1 + g.research.ResourceBonus(p.color, Stone)
			p.corn += 2
		},
		description: fmt.Sprintf("1 + %d gold, 1 + %d stone, 2 corn", g.research.ResourceBonus(p.color, Gold), g.research.ResourceBonus(p.color, Stone)),
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