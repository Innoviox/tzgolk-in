package model

func Palenque0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

func Palenque1(g *Game, p *Player) []Option {
	return []Option{
		func () {
			p.corn += 3 + g.research.CornBonus(p.color, Blue)
		},
	}
}

func Palenque2(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	if g.calendar.wheels[0].positions[2].pData.cornTiles > 0 {
		options = append(options, func() {
			p.corn += 4 + g.research.CornBonus(p.color, Green)
			p.cornTiles += 1
			g.calendar.wheels[0].positions[2].pData.cornTiles -= 1
		})
	} else if g.research.Irrigation(p.color) {
		options = append(options, func() {
			p.corn += 4 + g.research.CornBonus(p.color, Green)
		})
	}

	return options
}

func Jungle(corn int, wood int, position int) Options {
	return func (g *Game, p *Player) []Option {
		options := make([]Option, 0)

		if g.calendar.wheels[0].positions[position].pData.woodTiles > 0 {
			options = append(options, func () {
				p.resources[Wood] += wood + g.research.ResourceBonus(p.color, Wood)
				p.woodTiles += 1
				g.calendar.wheels[0].positions[3].pData.woodTiles -= 1
			})
	
			for i := 0; i < 3; i++ {
				if g.temples.CanStep(p.color, i, -1) {
					options = append(options, func () {
						p.corn += corn + g.research.CornBonus(p.color, Green)
						p.cornTiles += 1
						g.calendar.wheels[0].positions[3].pData.woodTiles -= 1
						g.calendar.wheels[0].positions[3].pData.cornTiles -= 1

						g.temples.Step(p.color, i, -1)
					})
				}
			}
		} 
		
		if g.calendar.wheels[0].positions[position].HasCornShowing() {
			options = append(options, func () {
				p.corn += corn + g.research.CornBonus(p.color, Green)
				p.cornTiles += 1
				g.calendar.wheels[0].positions[3].pData.cornTiles -= 1
			})
		}

		if g.research.Irrigation(p.color) {
			options = append(options, func() {
				p.corn += corn + g.research.CornBonus(p.color, Green)
			})
		}

		return options
	}
}

func Palenque() []Options {
	return []Options{
		Palenque0,
		Palenque1,
		Palenque2,
		Jungle(5, 2, 3),
		Jungle(7, 3, 4),
		Jungle(9, 4, 5),
	}
}

func MakePalenque() Wheel {
	return MakeWheel(Palenque(), 0, "Palenque")
}