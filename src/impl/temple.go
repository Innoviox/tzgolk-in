package impl 

import (
    . "tzgolkin/engine"
)

// todo real temple names
func Brown() *Temple {
	return &Temple {
		Steps: 7,
		PlayerLocations: map[Color]int{
			Red: 1,
			Green: 1,
			Blue: 1,
			Yellow: 1,
		},
		Age1Prize: 6,
		Age2Prize: 2,
		Points: []int{-1, 0, 2, 4, 6, 7, 8},
		Resources: map[int]Resource {
			2: Stone,
			4: Stone,
		},
	}
}

func YellowT() *Temple {
	return &Temple {
		Steps: 9,
		PlayerLocations: map[Color]int{
			Red: 1,
			Green: 1,
			Blue: 1,
			Yellow: 1,
		},
		Age1Prize: 2,
		Age2Prize: 6,
		Points: []int{-2, 0, 1, 2, 4, 6, 9, 12, 13},
		Resources: map[int]Resource {
			3: Gold,
			5: Gold,
		},
	}
}

func GreenT() *Temple {
	return &Temple {
		Steps: 8,
		PlayerLocations: map[Color]int {
			Red: 1,
			Green: 1,
			Blue: 1,
			Yellow: 1,
		},
		Age1Prize: 4,
		Age2Prize: 4,
		Points: []int{-3, 0, 1, 3, 5, 7, 9, 10},
		Resources: map[int]Resource {
			2: Wood,
			4: Wood,
			5: Skull,
		},
	}
}