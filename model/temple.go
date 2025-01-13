package model

import (
	"fmt"
	"os"
	"strings"
)

type Temple struct {
	steps int
	playerLocations map[Color]int
	age1Prize int
	age2Prize int
	points []int
	resources map[int]Resource
}

func (t *Temple) Clone() *Temple {
	newLocations := make(map[Color]int)
	for k, v := range t.playerLocations {
		newLocations[k] = v
	}

	return &Temple {
		steps: t.steps,
		playerLocations: newLocations,
		age1Prize: t.age1Prize,
		age2Prize: t.age2Prize,
		points: t.points,
		resources: t.resources,
	}
}

type Temples struct {
	temples []*Temple
}

func (t *Temples) Clone() *Temples {
	var newTemples = make([]*Temple, 0)
	for _, temple := range t.temples {
		newTemples = append(newTemples, temple.Clone())
	}

	return &Temples {
		temples: newTemples,
	}
}

// todo real temple names
func Brown() *Temple {
	return &Temple {
		steps: 7,
		playerLocations: map[Color]int{
			Red: 1,
			Green: 1,
			Blue: 1,
			Yellow: 1,
		},
		age1Prize: 6,
		age2Prize: 2,
		points: []int{-1, 0, 2, 4, 6, 7, 8},
		resources: map[int]Resource {
			2: Stone,
			4: Stone,
		},
	}
}

func YellowT() *Temple {
	return &Temple {
		steps: 9,
		playerLocations: map[Color]int{
			Red: 1,
			Green: 1,
			Blue: 1,
			Yellow: 1,
		},
		age1Prize: 2,
		age2Prize: 6,
		points: []int{-2, 0, 1, 2, 4, 6, 9, 12, 13},
		resources: map[int]Resource {
			3: Gold,
			5: Gold,
		},
	}
}

func GreenT() *Temple {
	return &Temple {
		steps: 8,
		playerLocations: map[Color]int {
			Red: 1,
			Green: 1,
			Blue: 1,
			Yellow: 1,
		},
		age1Prize: 4,
		age2Prize: 4,
		points: []int{-3, 0, 1, 3, 5, 7, 9, 10},
		resources: map[int]Resource {
			2: Wood,
			4: Wood,
			5: Skull,
		},
	}
}

func MakeTemples() *Temples {
	return &Temples{
		temples: []*Temple{Brown(), YellowT(), GreenT()},
	}
}

func (t *Temples) Step(p *Player, temple int, dir int) {
	t.temples[temple].playerLocations[p.color] += dir
	if t.temples[temple].playerLocations[p.color] < 0 {
		t.temples[temple].playerLocations[p.color] = 0
	}

	if t.temples[temple].playerLocations[p.color] >= t.temples[temple].steps {
		t.temples[temple].playerLocations[p.color] = t.temples[temple].steps - 1
		p.lightSide = true
	}
}

func (t *Temples) CanStep(p *Player, temple int, dir int) bool {
	if dir == -1 {
		return t.temples[temple].playerLocations[p.color] > 0
	} else if dir == 1 {
		return t.temples[temple].playerLocations[p.color] < t.temples[temple].steps - 1
	} else {
		return false 
	}
}

func (t *Temples) GainTempleStep(p *Player, o Option, dir int) []Option {
	options := make([]Option, 0)

	for i := 0; i < 3; i++ {
		if t.CanStep(p, i, dir) {
			options = append(options, Option{
				Execute: func() {
					t.Step(p, i, dir)
					o.Execute()
				},
				description: fmt.Sprintf("%s, %s temple %d", o.description, p.color.String(), dir),
			})
		}
	}

	return options
}

func (t *Temples) GainResources(p *Player) {
	for i := 0; i < 3; i++ {
		step := t.temples[i].playerLocations[p.color]
		for k, v := range t.temples[i].resources {
			if step >= k {
				p.resources[v] += 1
				fmt.Fprintf(os.Stdout, "giving %s to %s from temple %d\n", string(ResourceDebug[v]), p.color.String(), i)
			}
		}
	}
}

func (t *Temples) GainPoints(p *Player, age int) {
	for i := 0; i < 3; i++ {
		j := 0
		j += t.temples[i].points[t.temples[i].playerLocations[p.color]]

		isHighest := t.temples[i].IsHighest(p)
		if isHighest == 0 {
			if age == 1 {
				j += t.temples[i].age1Prize / 2
			} else {
				j += t.temples[i].age2Prize / 2
			}
		} else if isHighest == 1 {
			if age == 1 {
				j += t.temples[i].age1Prize
			} else {
				j += t.temples[i].age2Prize
			}
		}
		fmt.Fprintf(os.Stdout, "giving %d points to %s from temple %d\n", j, p.color.String(), i)
		p.points += j
	}
}

func (t *Temple) IsHighest(p *Player) int {
	step := t.playerLocations[p.color]

	highest := 1

	for i := 0; i < 4; i++ {
		if t.playerLocations[Color(i)] > step {
			return -1
		} else if Color(i) != p.color && t.playerLocations[Color(i)] == step {
			highest = 0
		}
	}

	return highest
}

func (t *Temples) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "----Temples------------\n")

	mostSteps := 0
	for _, temple := range t.temples {
		if temple.steps > mostSteps {
			mostSteps = temple.steps
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

	for i, temple := range t.temples {
		for c, step := range temple.playerLocations {
			fmt.Fprintf(&steps[step][i], "%s", c.String())
		}
	}

	for i := mostSteps; i >= 0; i-- {
		for j := 0; j < 3; j++ {
			if i == t.temples[j].steps {
				fmt.Fprintf(&br, " %s---------- ", string(TempleDebug[j]))
			} else if i > t.temples[j].steps {
				fmt.Fprintf(&br, "             ")
			} else {
				res, ok := t.temples[j].resources[i]
				if ok {
					fmt.Fprintf(&br, " |%s|", string(ResourceDebug[int(res)]))
				} else {
					fmt.Fprintf(&br, " | |")
				}

				fmt.Fprintf(&br, "%-4s|", steps[i][j].String())

				pts := t.temples[j].points[i]
				fmt.Fprintf(&br, "%2d| ", pts)
			}
		}
		fmt.Fprintf(&br, "\n")
	}

	steps = nil

	return br.String()
}