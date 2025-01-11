package model

// todo research
func Palenque0(g *Game) []Option {
	return make([]Option, 0)
}

func Palenque1(g *Game) []Option {
	return []Option{
		func (p *Player) {
			p.corn += 3
		},
	}
}

func Palenque2(g *Game) []Option {
	options := make([]Option, 0)

	if g.calendar.wheels[0].positions[2].pData.cornTiles > 0 {
		options = append(options, func(p *Player) {
			p.corn += 4
			p.cornTiles += 1
			g.calendar.wheels[0].positions[2].pData.cornTiles -= 1
		})
	}

	return options
}

func Palenque3(g *Game) []Option {
	options := make([]Option, 0)

	if g.calendar.wheels[0].positions[3].pData.woodTiles > 0 {
		options = append(options, func (p *Player) {
			p.resources[Wood] += 2
			p.woodTiles += 1
			g.calendar.wheels[0].positions[3].pData.woodTiles -= 1
		})

		// todo anger the gods
		options = append(options, func (p *Player) {
			p.corn += 5
			p.cornTiles += 1
			g.calendar.wheels[0].positions[3].pData.woodTiles -= 1
			g.calendar.wheels[0].positions[3].pData.cornTiles -= 1
		})
	} else if g.calendar.wheels[0].positions[3].pData.cornTiles > 0 {
		options = append(options, func (p *Player) {
			p.corn += 5
			p.cornTiles += 1
			g.calendar.wheels[0].positions[3].pData.cornTiles -= 1
		})
	}
}

func MakePalenque() Wheel {
	positions := make([]*Position, 0)
	positions = append(positions, &Position{
		wheel_id: 1,
		corn: 0,
		Execute: Palenque0,
		decisions: 0, 
	})


	return Wheel {
		id: 1,
		size: len(positions),
		occupied: make([]int, 0),
		workers: make([]int, 0),
		positions: positions,
		rotation: 0,
		name: "Palenque",
	}
}