package wheels

func Chichen0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

func ChichenX(temple int, points int, block boolean, position int) Options {
	return func(g *Game, p *Player) []Option {
		options := make([]Option, 0)

		if g.calendar.wheels[4].positions[position].cData.full {
			return options 
		}
		if block {
			for i := 0; i < 3; i++ {
				options = append(options, func() {
					g.temples.Step(p.color, temple, 1)
					p.points += points
					p.resources[i] += 1

					g.calendar.wheels[4].positions[position].cData.full = true
				})
			}
		} else {
			options = append(options, func() {
				g.temples.Step(p.color, temple, 1)
				p.points += points

				g.calendar.wheels[4].positions[position].cData.full = true
			})
		}

		return options
	}
}

func Chichen() []Options {
	return []Options{
		Chichen0,
		ChichenX(0, 4, false, 1),
		ChichenX(0, 5, false, 2),
		ChichenX(0, 6, false, 3),
		ChichenX(1, 7, false, 4),
		ChichenX(1, 8, false, 5),
		ChichenX(1, 8, true, 6),
		ChichenX(2, 10, false, 7),
		ChichenX(2, 11, true, 8),
		ChichenX(2, 13, true, 9),
	}
}

func MakeChichen() Wheel {
	positions := make([]*Position, 0)

	options := Chichen()

	for i := 0; i < len(options); i++ {
		positions = append(positions, &Position{
			wheel_id: 4,
			corn: i,
			GetOptions: options[i],
			cData: MakeCData(),
		})
	}

	positions = append(positions, &Position {
		wheel_id: 4, 
		corn: i,
		GetOptions: flatten(options),
	})

	return &Wheel {
		id: 4,
		size: len(positions),
		occupied: make([]int, 0),
		workers: make([]int, 0),
		positions: positions,
		rotation: 0,
		name: "Chichen Itza",
	}
}