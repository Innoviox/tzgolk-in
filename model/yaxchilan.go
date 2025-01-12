package model

// todo research
func YaxchilanWrapper(f func (*Game, *Player)) Options {
	return func (g *Game) []Option {
		return []Option{
			func (p *Player) {
				f(g, p)
			},
		}
	}
}

func Yaxchilan0(g *Game) []Option {
	return make([]Option, 0)
}

func Yaxchilan1(g *Game) []Option {
	return []Option{
		func (p *Player) {
			p.resources[Wood] += 1
		},
	}
}

func Yaxchilan2(g *Game) []Option {
	return []Option {
		func (p *Player) {
			p.resources[Stone] += 1
			p.corn += 1
		},
	}
}

func Yaxchilan3(g *Game) []Option {
	return []Option {
		func (p *Player) {
			p.resources[Gold] += 1
			p.corn += 1
		},
	}
}

func Yaxchilan4(g *Game) []Option {
	return []Option {
		func (p *Player) {
			p.resources[Skull] += 1
		},
	}
}

func Yaxchilan5(g *Game) []Option {
	return []Option {
		func (p *Player) {
			p.resources[Gold] += 1
			p.resources[Stone] += 1
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