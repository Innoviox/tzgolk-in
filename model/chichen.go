package model

import (
	"fmt"
)

func Chichen0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

type ChichenSpot struct {
	temple int
	points int
	block bool
	position int
}

func ChichenSpots() []ChichenSpot {
	return []ChichenSpot{
		{0, 4, false, 1},
		{0, 5, false, 2},
		{0, 6, false, 3},
		{2, 7, false, 4},
		{2, 8, false, 5},
		{2, 8, true, 6},
		{1, 10, false, 7},
		{1, 11, true, 8},
		{1, 13, true, 9},
	}
}


func ChichenX(n int) Options {
	spot := ChichenSpots()[n]
	return func(g *Game, p *Player) []Option {
		ChichenHelper := func () []Option {
			options := make([]Option, 0)
		
			if spot.block {
				// if blocK: generate option for gaining each block
				for i := 0; i < 3; i++ {
					options = append(options, Option{
						Execute: func() {
							g.temples.Step(p, spot.temple, 1)
							p.points += spot.points
							p.resources[i] += 1
							p.resources[Skull] -= 1
			
							g.calendar.wheels[4].positions[spot.position].cData.full = true
						},
						description: fmt.Sprintf("%s temple, %d points, 1 %sT", string(ResourceDebug[i]), spot.points, string(TempleDebug[spot.temple])),
					})
				}
			} else {
				// just generate option for points
				options = append(options, Option{
					Execute: func() {
						g.temples.Step(p, spot.temple, 1)
						p.points += spot.points
						p.resources[Skull] -= 1
			
						g.calendar.wheels[4].positions[spot.position].cData.full = true
					},
					description: fmt.Sprintf("%s temple, %d points", string(TempleDebug[spot.temple]), spot.points),
				})
			}
		
			return options
		}

		options := make([]Option, 0)

		if g.research.Foresight(p.color) {
			if n < 8 {
				options = append(options, ChichenX(n + 1)(g, p)...)
			} else {
				// mimic mirror
				for i := 0; i < 8; i++ {
					options = append(options, ChichenX(i)(g, p)...)
				}
			}
		}

		if g.calendar.wheels[4].positions[spot.position].cData.full || p.resources[Skull] == 0 {
			return options 
		}

		if g.research.Devout(p.color) {
			// for each block
			for i := 0; i < 3; i++ {
				if p.resources[i] > 0 {
					for _, o := range ChichenHelper() {
						// add "spend block for temple" to each option
						options = append(options, g.temples.GainTempleStep(p, Option {
							Execute: func() {
								p.resources[i] -= 1
								o.Execute()
							},
							description: fmt.Sprintf("%s, [theo] pay 1 %s", o.description, string(ResourceDebug[i])),
						}, 1)...)
					}
				}
			}
		} else {
			// just use each option
			options = append(options, ChichenHelper()...)
		}

		return options
	}
}
func Chichen() []Options {
	options := make([]Options, 0)

	options = append(options, Chichen0)

	for i := range ChichenSpots() {
		options = append(options, ChichenX(i))
	}

	return options
}

func MakeChichen() *Wheel {
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
		corn: 10,
		GetOptions: flatten(options),
	})

	return &Wheel {
		id: 4,
		size: len(positions),
		occupied: make([]int, 0),
		workers: make([]int, 0),
		positions: positions,
		name: "Chichen Itza",
	}
}