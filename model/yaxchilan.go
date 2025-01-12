package model

func Yaxchilan0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

func Yaxchilan1(g *Game, p *Player) {
	return []Option {
		func() {
			p.resources[Wood] += 1 + g.research.ResourceBonus(p.color, Wood)
		}
	}
}

func Yaxchilan2(g *Game) []Option {
	return []Option {
		func (p *Player) {
			p.resources[Stone] += 1 + g.research.ResourceBonus(p.color, Stone)
			p.corn += 1
		},
	}
}

func Yaxchilan3(g *Game) []Option {
	return []Option {
		func (p *Player) {
			p.resources[Gold] += 1 + g.research.ResourceBonus(p.color, Gold)
			p.corn += 1
		},
	}
}

func Yaxchilan4(g *Game) []Option {
	return []Option {
		func (p *Player) {
			p.resources[Skull] += 1 + g.research.ResourceBonus(p.color, Skull)
		},
	}
}

func Yaxchilan5(g *Game) []Option {
	return []Option {
		func (p *Player) {
			p.resources[Gold] += 1 + g.research.ResourceBonus(p.color, Gold)
			p.resources[Stone] += 1 + g.research.ResourceBonus(p.color, Stone)
			p.corn += 2
		},
	}
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
	return MakeWheel(Yaxchilan(), 2, "Yaxchilan")
}