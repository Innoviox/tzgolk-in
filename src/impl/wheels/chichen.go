package wheels

import (
	"fmt"
	"strings"
	. "tzgolkin/engine"
)

func Chichen0(g *Game, p *Player) []*Delta {
	return make([]*Delta, 0)
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
	return func(g *Game, p *Player) []*Delta {
		ChichenHelperHelper := func (r int) *Delta {
			d := g.Temples.Step(p, spot.Temple, 1)
			pd := PlayerDelta{}
			pd.Points = spot.Points
			pd.Resources[Skull] = -1
			if r != -1 {
				pd.Resources[r] = 1
			}
			f := WheelDelta{PositionDeltas: map[int]PositionDelta{spot.Position: PositionDelta{CData: ChichenDataDelta{
				Full: 1,
			}}}}

			d.PlayerDeltas = map[Color]PlayerDelta{}
			d.PlayerDeltas[p.Color] = pd

			d.CalendarDelta = CalendarDelta{WheelDeltas: map[int]WheelDelta{4: f}}

			return d
		}

		ChichenHelper := func () []*Delta {
			options := make([]*Delta, 0)

			if spot.Block {
				// if blocK: generate option for gaining each block
				for i := 0; i < 3; i++ {
					options = append(options, ChichenHelperHelper(i))
				}
			} else {
				// just generate option for points
				options = append(options, ChichenHelperHelper(-1))
			}
		
			return options
		}

		options := make([]*Delta, 0)

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
			return Skip() 
		}

		if g.Research.Devout(p.Color) {
			// for each block
			for i := 0; i < 3; i++ {
				if p.Resources[i] > 0 {
					for _, o := range ChichenHelper() {
						// add "spend block for temple" to each option
						r := [4]int{0, 0, 0, 0}
						r[i] = -1 

						d := &Delta{PlayerDeltas: map[Color]PlayerDelta{p.Color: PlayerDelta{Resources: r}}}
						d.Description = fmt.Sprintf("[theo] pay 1 %s", string(ResourceDebug[i]))

						d.Add(o)
						options = append(options, g.Temples.GainTempleStep(p, d, 1)...)
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

func ChichenString(wheel *Wheel, workers []*Worker) string {
	var br strings.Builder

	fmt.Fprintf(&br, "| %-12s: ", wheel.Name)

	out := make([]string, wheel.Size)

	for k, v := range wheel.Occupied {
		if v >= 0 {
			out[k] = workers[v].Color.String()
		}
	}

	for k, o := range out {
		if len(o) > 0 {
			fmt.Fprintf(&br, "  %s", o)
		} else if wheel.Positions[k].CData != nil && wheel.Positions[k].CData.Full {
			fmt.Fprintf(&br, "  X")
		} else {
			fmt.Fprintf(&br, "%3d", k)
		}
	}
	fmt.Fprintf(&br, "\n")

	return br.String()
}

func MakeChichen() *Wheel {
	positions := make([]*Position, 0)
	occupied := make(map[int]int)

	options := Chichen()

	for i := 0; i < len(options); i++ {
		positions = append(positions, &Position{
			Wheel_id: 4,
			Corn: i,
			GetOptions: options[i],
			CData: MakeCData(),
		})
		occupied[i] = -1
	}

	positions = append(positions, &Position {
		Wheel_id: 4, 
		Corn: 10,
		GetOptions: Flatten(options),
	})
	occupied[10] = -1

	return &Wheel {
		Id: 4,
		Size: len(positions),
		Occupied: occupied,
		Positions: positions,
		Name: "Chichen Itza",
		String: ChichenString,
	}
}