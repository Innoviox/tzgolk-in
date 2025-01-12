package model

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
		resources: map[int][4]int {
			2: Stone,
			6: Stone,
		},
	}
}

func MakeTemples() *Temples {
	return &Temples{
		temples: []Temple{Brown()},
	}
}

func (t *Temples) Step(c Color, temple int, dir int) {
	t.temples[temple].playerLocations[c] += dir
}

// func (t *Temples) GetSteps(p *Player) Options