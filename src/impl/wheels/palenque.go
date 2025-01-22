package wheels

import (
	"fmt"
	"strings"
	. "tzgolkin/engine"
	. "tzgolkin/delta"
)

func Palenque0(g *Game, p *Player) []*Delta {
	return make([]*Delta, 0)
}

func Palenque1(g *Game, p *Player) []*Delta {
	d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
		Corn: 3 + g.Research.CornBonus(p.Color, Blue),
	})
	d.Description = fmt.Sprintf("3 + %d Corn", g.Research.CornBonus(p.Color, Blue))

	return []*Delta{d}
}

func Palenque2(g *Game, p *Player) []*Delta {
	options := make([]*Delta, 0)

	if g.Calendar.Wheels[0].Positions[2].PData.CornTiles > 0 {
		d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
			Corn: 4 + g.Research.CornBonus(p.Color, Blue),
			CornTiles: 1,
		})

		tmpDelta := GetDelta()
		tmpDelta.CalendarDelta.WheelDeltas = WheelDeltaMapPool.Get().(map[int]WheelDelta)
		positionDeltaMap := PositionDeltaMapPool.Get().(map[int]PositionDelta)
		positionDeltaMap[2] = PositionDelta{
			PData: PalenqueDataDelta{
				CornTiles: -1,
			},
		}
		wheelDelta := WheelDelta{
			PositionDeltas: positionDeltaMap,
		}
		tmpDelta.CalendarDelta.WheelDeltas[0] = wheelDelta
		d.Add(tmpDelta)
		PutDelta(tmpDelta)

		d.Description = fmt.Sprintf("4 + %d Corn", g.Research.CornBonus(p.Color, Blue))
		options = append(options, d)
	} else if g.Research.Irrigation(p.Color) {
		d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
			Corn: 4 + g.Research.CornBonus(p.Color, Blue),
		})
		d.Description = fmt.Sprintf("4 + %d Corn (irrigation)", g.Research.CornBonus(p.Color, Green))

		options = append(options, d)
	} else {
		options = Skip()
	}

	return options
}

func Jungle(Corn int, wood int, position int) Options {
	return func (g *Game, p *Player) []*Delta {
		options := make([]*Delta, 0)

		if g.Calendar.Wheels[0].Positions[position].PData.WoodTiles > 0 {
			d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
				Resources: [4]int{wood + g.Research.ResourceBonus(p.Color, Wood), 0, 0, 0},
				WoodTiles: 1,
			})

			tmpDelta := GetDelta()
			tmpDelta.CalendarDelta.WheelDeltas = WheelDeltaMapPool.Get().(map[int]WheelDelta)
			positionDeltaMap := PositionDeltaMapPool.Get().(map[int]PositionDelta)
			positionDeltaMap[position] = PositionDelta{
				PData: PalenqueDataDelta{
					WoodTiles: -1,
				},
			}
			wheelDelta := WheelDelta{
				PositionDeltas: positionDeltaMap,
			}
			tmpDelta.CalendarDelta.WheelDeltas[0] = wheelDelta
			d.Add(tmpDelta)
			PutDelta(tmpDelta)
	

			d.Description = fmt.Sprintf("%d + %d wood", wood, g.Research.ResourceBonus(p.Color, Wood))
			options = append(options, d)

			d2 := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
				Corn: Corn + g.Research.CornBonus(p.Color, Green),
				CornTiles: 1,
			})

			tmpDelta2 := GetDelta()
			tmpDelta2.CalendarDelta.WheelDeltas = WheelDeltaMapPool.Get().(map[int]WheelDelta)
			positionDeltaMap2 := PositionDeltaMapPool.Get().(map[int]PositionDelta)
			positionDeltaMap2[position] = PositionDelta{
				PData: PalenqueDataDelta{
					WoodTiles: -1,
					CornTiles: -1,
				},
			}
			wheelDelta2 := WheelDelta{
				PositionDeltas: positionDeltaMap,
			}
			tmpDelta2.CalendarDelta.WheelDeltas[0] = wheelDelta2
			d2.Add(tmpDelta2)
			PutDelta(tmpDelta2)

			d2.Description = fmt.Sprintf("%d + %d Corn", Corn, g.Research.CornBonus(p.Color, Green))
			options = append(options, g.Temples.GainTempleStep(p, d2, -1)...)
		} 
		
		if g.Calendar.Wheels[0].Positions[position].PData.HasCornShowing() {
			d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
				Corn: Corn + g.Research.CornBonus(p.Color, Green),
				CornTiles: 1,
			})

			tmpDelta := GetDelta()
			tmpDelta.CalendarDelta.WheelDeltas = WheelDeltaMapPool.Get().(map[int]WheelDelta)
			positionDeltaMap := PositionDeltaMapPool.Get().(map[int]PositionDelta)
			positionDeltaMap[position] = PositionDelta{
				PData: PalenqueDataDelta{
					CornTiles: -1,
				},
			}
			wheelDelta := WheelDelta{
				PositionDeltas: positionDeltaMap,
			}
			tmpDelta.CalendarDelta.WheelDeltas[0] = wheelDelta
			d.Add(tmpDelta)
			PutDelta(tmpDelta)

			d.Description = fmt.Sprintf("%d + %d Corn", Corn, g.Research.CornBonus(p.Color, Green))
			options = append(options, d)
		}

		if g.Research.Irrigation(p.Color) {
			d := PlayerDeltaWrapper(int(p.Color), PlayerDelta{
				Corn: Corn + g.Research.CornBonus(p.Color, Green),
			})
			d.Description = fmt.Sprintf("%d + %d Corn (irrigation)", Corn, g.Research.CornBonus(p.Color, Green))
			options = append(options, d)
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

func PalenqueString(wheel *Wheel, workers []*Worker) string {
	var br strings.Builder

	fmt.Fprintf(&br, "|                ")

	for k := 0; k < wheel.Size; k++ {
		if wheel.Positions[k].PData != nil {
			fmt.Fprintf(&br, "%d%d ", wheel.Positions[k].PData.WoodTiles, wheel.Positions[k].PData.CornTiles)
		} else {
			fmt.Fprintf(&br, "   ")
		}
	}
	fmt.Fprintf(&br, "\n")
		
	fmt.Fprintf(&br, "| %-12s: ", wheel.Name)

	out := make([]string, wheel.Size)

	for k, v := range wheel.Occupied {
		if v >= 0 {
			out[v] = workers[k].Color.String()
		}
	}

	for k, o := range out {
		if len(o) > 0 {
			fmt.Fprintf(&br, "  %s", o)
		} else {
			fmt.Fprintf(&br, "%3d", k)
		}
	}
	fmt.Fprintf(&br, "\n")

	return br.String()
}

func MakePalenque() *Wheel {
	positions := make([]*Position, 0)

	options := Palenque()

	for i := 0; i < 2; i++ {
		positions = append(positions, &Position{
			Wheel_id: 0,
			Corn: i,
			GetOptions: options[i],
			// PData: MakePData(i > 2),
		})
	}

	for i := 2; i < len(options); i++ {
		positions = append(positions, &Position{
			Wheel_id: 0,
			Corn: i,
			GetOptions: options[i],
			PData: MakePData(i > 2),
		})
	}

	for i := 6; i < 8; i++ {
		positions = append(positions, &Position{
			Wheel_id: 0,
			Corn: i,
			GetOptions: Flatten(options),
		})
	}

	return &Wheel{
		Id: 0,
		Size: len(positions),
		Occupied: MakeOccupied(24),
		Positions: positions, 
		Name: "Palenque",
		String: PalenqueString,
	}
}