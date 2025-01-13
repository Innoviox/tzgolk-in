package wheels

import (
	"fmt"
	. "tzgolkin/model"
)

func Chichen0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

type ChichenSpot struct {
	Temple int
	Points int
	Block bool
	Position int
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


func ChichenX(n int, canForesight bool) Options {
	spot := ChichenSpots()[n]
	return func(g *Game, p *Player) []Option {
		ChichenHelper := func () []Option {
			options := make([]Option, 0)
		
			if spot.Block {
				// if blocK: generate option for gaining each block
				for i := 0; i < 3; i++ {
					options = append(options, Option{
						Execute: func(g *Game, p *Player) {
							g.Temples.Step(p, spot.Temple, 1)
							p.Points += spot.Points
							p.Resources[i] += 1
							p.Resources[Skull] -= 1
			
							g.Calendar.Wheels[4].Positions[spot.Position].CData.Full = true
						},
						Description: fmt.Sprintf("%s temple, %d points, 1 %sT", string(ResourceDebug[i]), spot.Points, string(TempleDebug[spot.Temple])),
					})
				}
			} else {
				// just generate option for points
				options = append(options, Option{
					Execute: func(g *Game, p *Player) {
						g.Temples.Step(p, spot.Temple, 1)
						p.Points += spot.Points
						p.Resources[Skull] -= 1
			
						g.Calendar.Wheels[4].Positions[spot.Position].CData.Full = true
					},
					Description: fmt.Sprintf("%s temple, %d points", string(TempleDebug[spot.Temple]), spot.Points),
				})
			}
		
			return options
		}

		options := make([]Option, 0)

		options = append(options, Option{
			Execute: func(g *Game, p *Player) {

			},
			Description: "skip",
		})

		if canForesight && g.Research.Foresight(p.Color) {
			if n < 8 {
				options = append(options, ChichenX(n + 1, false)(g, p)...)
			} else {
				// mimic mirror
				for i := 0; i < 8; i++ {
					options = append(options, ChichenX(i, false)(g, p)...)
				}
			}
		}

		if g.Calendar.Wheels[4].Positions[spot.Position].CData.Full || p.Resources[Skull] == 0 {
			return options 
		}

		if g.Research.Devout(p.Color) {
			// for each block
			for i := 0; i < 3; i++ {
				if p.Resources[i] > 0 {
					for _, o := range ChichenHelper() {
						// add "spend block for temple" to each option
						options = append(options, g.Temples.GainTempleStep(p, Option {
							Execute: func(g *Game, p *Player) {
								p.Resources[i] -= 1
								o.Execute(g, p)
							},
							Description: fmt.Sprintf("%s, [theo] pay 1 %s", o.Description, string(ResourceDebug[i])),
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
		options = append(options, ChichenX(i, true))
	}

	return options
}

func MakeChichen() *Wheel {
	positions := make([]*Position, 0)

	options := Chichen()

	for i := 0; i < len(options); i++ {
		positions = append(positions, &Position{
			Wheel_id: 4,
			Corn: i,
			GetOptions: options[i],
			CData: MakeCData(),
		})
	}

	positions = append(positions, &Position {
		Wheel_id: 4, 
		Corn: 10,
		GetOptions: Flatten(options),
	})

	return &Wheel {
		Id: 4,
		Size: len(positions),
		Occupied: make(map[int]int),
		Positions: positions,
		Name: "Chichen Itza",
	}
}