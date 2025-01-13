package model

import (
	"fmt"
	"os"
	"strings"
)

type Temple struct {
	Steps int
	PlayerLocations map[Color]int
	Age1Prize int
	Age2Prize int
	Points []int
	Resources map[int]Resource
}

func MakeTemple(steps int, playerLocations map[Color]int, age1Prize int, age2Prize int, points []int, resources map[int]Resource) *Temple {
	return &Temple {
		Steps: steps,
		PlayerLocations: playerLocations,
		Age1Prize: age1Prize,
		Age2Prize: age2Prize,
		Points: points,
		Resources: resources,
	}
}

func MakeTemples(temples []*Temple) *Temples {
	return &Temples{
		Temples: temples,
	}
}

func (t *Temple) Clone() *Temple {
	newLocations := make(map[Color]int)
	for k, v := range t.PlayerLocations {
		newLocations[k] = v
	}

	return &Temple {
		Steps: t.Steps,
		PlayerLocations: newLocations,
		Age1Prize: t.Age1Prize,
		Age2Prize: t.Age2Prize,
		Points: t.Points,
		Resources: t.Resources,
	}
}

type Temples struct {
	Temples []*Temple
}

func (t *Temples) Clone() *Temples {
	var newTemples = make([]*Temple, 0)
	for _, temple := range t.Temples {
		newTemples = append(newTemples, temple.Clone())
	}

	return &Temples {
		Temples: newTemples,
	}
}

func (t *Temples) Step(p *Player, temple int, dir int) {
	t.Temples[temple].PlayerLocations[p.Color] += dir
	if t.Temples[temple].PlayerLocations[p.Color] < 0 {
		t.Temples[temple].PlayerLocations[p.Color] = 0
	}

	if t.Temples[temple].PlayerLocations[p.Color] >= t.Temples[temple].Steps {
		t.Temples[temple].PlayerLocations[p.Color] = t.Temples[temple].Steps - 1
		p.LightSide = true
	}
}

func (t *Temples) CanStep(p *Player, temple int, dir int) bool {
	if dir == -1 {
		return t.Temples[temple].PlayerLocations[p.Color] > 0
	} else if dir == 1 {
		return t.Temples[temple].PlayerLocations[p.Color] < t.Temples[temple].Steps - 1
	} else {
		return false 
	}
}

func (t *Temples) GainTempleStep(p *Player, o Option, dir int) []Option {
	options := make([]Option, 0)

	for i := 0; i < 3; i++ {
		if t.CanStep(p, i, dir) {
			options = append(options, Option{
				Execute: func(g *Game, p *Player) {
					g.Temples.Step(p, i, dir)
					o.Execute(g, p)
				},
				Description: fmt.Sprintf("%s, %s temple %d", o.Description, p.Color.String(), dir),
			})
		}
	}

	return options
}

func (t *Temples) GainResources(p *Player) {
	for i := 0; i < 3; i++ {
		step := t.Temples[i].PlayerLocations[p.Color]
		for k, v := range t.Temples[i].Resources {
			if step >= k {
				p.Resources[v] += 1
				fmt.Fprintf(os.Stdout, "giving %s to %s from temple %d\n", string(ResourceDebug[v]), p.Color.String(), i)
			}
		}
	}
}

func (t *Temples) GainPoints(p *Player, age int) {
	for i := 0; i < 3; i++ {
		j := 0
		j += t.Temples[i].Points[t.Temples[i].PlayerLocations[p.Color]]

		isHighest := t.Temples[i].IsHighest(p)
		if isHighest == 0 {
			if age == 1 {
				j += t.Temples[i].Age1Prize / 2
			} else {
				j += t.Temples[i].Age2Prize / 2
			}
		} else if isHighest == 1 {
			if age == 1 {
				j += t.Temples[i].Age1Prize
			} else {
				j += t.Temples[i].Age2Prize
			}
		}
		fmt.Fprintf(os.Stdout, "giving %d points to %s from temple %d\n", j, p.Color.String(), i)
		p.Points += j
	}
}

func (t *Temple) IsHighest(p *Player) int {
	step := t.PlayerLocations[p.Color]

	highest := 1

	for i := 0; i < 4; i++ {
		if t.PlayerLocations[Color(i)] > step {
			return -1
		} else if Color(i) != p.Color && t.PlayerLocations[Color(i)] == step {
			highest = 0
		}
	}

	return highest
}

func (t *Temples) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "----Temples------------\n")

	mostSteps := 0
	for _, temple := range t.Temples {
		if temple.Steps > mostSteps {
			mostSteps = temple.Steps
		}
	}

	var steps [][]strings.Builder
	for i := 0; i < mostSteps + 1; i++ {
		row := make([]strings.Builder, 0)
		for j := 0; j < 3; j++ {
			row = append(row, strings.Builder{})
		}

		steps = append(steps, row)
	}

	for i, temple := range t.Temples {
		for c, step := range temple.PlayerLocations {
			fmt.Fprintf(&steps[step][i], "%s", c.String())
		}
	}

	for i := mostSteps; i >= 0; i-- {
		for j := 0; j < 3; j++ {
			if i == t.Temples[j].Steps {
				fmt.Fprintf(&br, " %s---------- ", string(TempleDebug[j]))
			} else if i > t.Temples[j].Steps {
				fmt.Fprintf(&br, "             ")
			} else {
				res, ok := t.Temples[j].Resources[i]
				if ok {
					fmt.Fprintf(&br, " |%s|", string(ResourceDebug[int(res)]))
				} else {
					fmt.Fprintf(&br, " | |")
				}

				fmt.Fprintf(&br, "%-4s|", steps[i][j].String())

				pts := t.Temples[j].Points[i]
				fmt.Fprintf(&br, "%2d| ", pts)
			}
		}
		fmt.Fprintf(&br, "\n")
	}

	steps = nil

	return br.String()
}