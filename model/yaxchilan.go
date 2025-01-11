package model

// todo research
func YaxchilanWrapper(f func (*Player)) Options {
	return func (g *Game) []Option {
		return []Option{
			func (p *Player) {
				f(p)
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
	positions := make([]*Position, 0)

	options := Yaxchilan()

	for i := 0; i < len(options); i++ {
		positions = append(positions, &Position{
			wheel_id: 2,
			corn: i,
			GetOptions: options[i],
		})
	}

	for i := 6; i < 8; i++ {
		positions = append(positions, &Position{
			wheel_id: 2,
			corn: i,
			GetOptions: flatten(options),
		})
	}

	return &Wheel{
		id: 2,
		size: len(positions),
		occupied: make([]int, 0),
		workers: make([]int, 0),
		positions: positions, 
		rotation: 0,
		name: "Yaxchilan",
	}
}