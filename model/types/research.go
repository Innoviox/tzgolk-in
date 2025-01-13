package types

type Research struct {
	levels map[Color]Levels
}

type Levels map[Science]int
