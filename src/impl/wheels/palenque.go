package wheels

import (
	"fmt"
	"strings"
	. "tzgolkin/model"
)

func Palenque0(g *Game, p *Player) []Option {
	return make([]Option, 0)
}

func Palenque1(g *Game, p *Player) []Option {
	return []Option{Option{
		Execute: func (g *Game, p *Player) {
			p.Corn += 3 + g.Research.CornBonus(p.Color, Blue)
		},
		Description: fmt.Sprintf("3 + %d Corn", g.Research.CornBonus(p.Color, Blue)),
	}}
}

func Palenque2(g *Game, p *Player) []Option {
	options := make([]Option, 0)

	if g.Calendar.Wheels[0].Positions[2].PData.CornTiles > 0 {
		options = append(options, Option{
			Execute: func(g *Game, p *Player) {
				p.Corn += 4 + g.Research.CornBonus(p.Color, Green)
				p.CornTiles += 1
				g.Calendar.Wheels[0].Positions[2].PData.CornTiles -= 1
			},
			Description: fmt.Sprintf("4 + %d Corn", g.Research.CornBonus(p.Color, Green)),
		})
	} else if g.Research.Irrigation(p.Color) {
		options = append(options, Option{
			Execute: func(g *Game, p *Player) {
				p.Corn += 4 + g.Research.CornBonus(p.Color, Green)
			},
			Description: fmt.Sprintf("4 + %d Corn (irrigation)", g.Research.CornBonus(p.Color, Green)),
		})
	}

	return options
}

func Jungle(Corn int, wood int, position int) Options {
	return func (g *Game, p *Player) []Option {
		options := make([]Option, 0)

		if g.Calendar.Wheels[0].Positions[position].PData.WoodTiles > 0 {
			options = append(options, Option{
				Execute: func (g *Game, p *Player) {
					p.Resources[Wood] += wood + g.Research.ResourceBonus(p.Color, Wood)
					p.WoodTiles += 1
					g.Calendar.Wheels[0].Positions[3].PData.WoodTiles -= 1
				},
				Description: fmt.Sprintf("%d + %d wood", wood, g.Research.ResourceBonus(p.Color, Wood)),
			})
	
			options = append(options, g.Temples.GainTempleStep(p, Option{
				Execute: func (g *Game, p *Player) {
					p.Corn += Corn + g.Research.CornBonus(p.Color, Green)
					p.CornTiles += 1
					g.Calendar.Wheels[0].Positions[3].PData.WoodTiles -= 1
					g.Calendar.Wheels[0].Positions[3].PData.CornTiles -= 1
				},
				Description: fmt.Sprintf("%d + %d Corn, anger", Corn, g.Research.CornBonus(p.Color, Green)),
			}, -1)...)
		} 
		
		if g.Calendar.Wheels[0].Positions[position].PData.HasCornShowing() {
			options = append(options, Option{
				Execute: func (g *Game, p *Player) {
					p.Corn += Corn + g.Research.CornBonus(p.Color, Green)
					p.CornTiles += 1
					g.Calendar.Wheels[0].Positions[3].PData.CornTiles -= 1
				},
				Description: fmt.Sprintf("%d + %d Corn", Corn, g.Research.CornBonus(p.Color, Green)),
			})
		}

		if g.Research.Irrigation(p.Color) {
			options = append(options, Option{
				Execute: func(g *Game, p *Player) {
					p.Corn += Corn + g.Research.CornBonus(p.Color, Green)
				},
				Description: fmt.Sprintf("%d + %d Corn (irrigation)", Corn, g.Research.CornBonus(p.Color, Green)),
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

func PalenqueString(wheel *Wheel, workers []*Worker) string {
	var br strings.Builder

	fmt.Fprintf(&br, "|                ")

	for k := 0; k < wheel.Size; k++ {
		if wheel.Positions[k].PData != nil {
			// ct := "⁰¹²³⁴"[wheel.Positions[k].PData.CornTiles]
			ct := []rune{'₀', '₁', '₂', '₃', '₄'}[wheel.Positions[k].PData.CornTiles]
			wt := []rune{'₀', '₁', '₂', '₃', '₄'}[wheel.Positions[k].PData.WoodTiles]
			// wt := "₀₁₂₃₄"[wheel.Positions[k].PData.WoodTiles]
			// fmt.Fprintf(&br, "%q%d%q", ct, k, wt)
			// char := '⁰'
			// fmt.Fprintf(&br, " %q ", char)
			br.WriteRune(rune(wt))
			br.WriteRune(rune(ct))
			fmt.Fprintf(&br, " ")
			
			// br.WriteRune(rune(wt))
		} else {
			fmt.Fprintf(&br, "   ")
		}
	}
	fmt.Fprintf(&br, "\n")
		
	fmt.Fprintf(&br, "| %-12s: ", wheel.Name)

	out := make([]string, wheel.Size)

	for k, v := range wheel.Occupied {
		out[k] = workers[v].Color.String()
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
		Occupied: make(map[int]int),
		Positions: positions, 
		Name: "Palenque",
		String: PalenqueString,
	}
}