package model

import (
	"fmt"
)

type Temple struct {
	steps int
	playerLocations map[Color]int
	age1Prize int
	age2Prize int
	points []int
	resources map[int]Resource
}

type Temples struct {
	temples []*Temple
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

func (t *Temples) Step(c Color, temple int, dir int) {
	t.temples[temple].playerLocations[c] += dir
}

func (t *Temples) CanStep(c Color, temple int, dir int) bool {
	if dir == -1 {
		return t.temples[temple].playerLocations[c] > 0
	} else if dir == 1 {
		return t.temples[temple].playerLocations[c] < t.temples[temple].steps - 1
	} else {
		return false 
	}
}

func (t *Temples) GainTempleStep(c Color, o Option, dir int) []Option {
	options := make([]Option, 0)

	for i := 0; i < 3; i++ {
		if t.CanStep(c, i, dir) {
			options = append(options, Option{
				Execute: func() {
					t.Step(c, i, dir)
					o.Execute()
				},
				description: fmt.Sprintf("%s, %s temple %d", o.description, c.String(), dir),
			})
		}
	}

	return options
}