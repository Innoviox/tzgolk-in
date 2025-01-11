package model

type Science int

// todo get actual names
const (
	Agriculture Science = iota
	Resources
	Construction
	Theology
)


type Research struct {
	levels map[Color]map[Science]int
}

func MakeLevels() map[Science]int {
	return map[Science]int{
		Agriculture: 0,
		Resources: 0,
		Construction: 0,
		Theology: 0,
	}
}

func MakeResearch() *Research {
	return &Research{
		levels: map[Color]map[Science]int{
			Red: MakeLevels(),
			Green: MakeLevels(),
			Blue: MakeLevels(),
			Yellow: MakeLevels(),
		},
	}
}

func (r *Research) HasLevel(c Color, s Science, level int) bool {
	return r.levels[c][s] >= level
}