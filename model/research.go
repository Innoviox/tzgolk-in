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

func (r *Research) Devout(c Color) bool {
	return r.HasLevel(c, Theology, 2)
}

func (r *Research) CornBonus(player Color, tile Color) int {
	if tile == Blue {
		if r.HasLevel(player, Agriculture, 2) {
			return 1
		} else {
			return 0
		}
	} else if tile == Green {
		if r.HasLevel(player, Agriculture, 1) {
			return 1
		} else if r.HasLevel(player, Agriculture, 3) {
			return 3
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func (r *Research) Irrigation(c Color) bool {
	return r.HasLevel(c, Agriculture, 2)
}

func (r *Research) ResourceBonus(c Color, res Resource) int {
	if r.HasLevel(c, Resources, int(res)) {
		return 1
	} else {
		return 0
	}
}

func (r *Research) Built(p *Player) {
	if r.HasLevel(p.color, Construction, 1) {
		p.corn += 1
	}

	if r.HasLevel(p.color, Construction, 2) {
		p.points += 1
	}
}

func (r *Research) Builder(c Color) bool {
	return r.HasLevel(c, Construction, 3)
}

func (r *Research) GetOptions(p *Player, n int) []Option{ 

}

func (r *Research) GetOptionsHelper(resources [4]int, levels map[Science]int) {
	// options := make([]Option, 0)
	for s := 0; s < 4; s++ {
		level := levels[Science(s)]
		if level < 3 {
			for _, newResources := range PayBlocks(resources, level) {
				// yield newResources, newLevels
			}
		} else {
			for _, newResources := range PayBlocks(resources, 1) {
				switch Science(s) {
				case Agriculture:
					for i := 0; i < 3; i++ {
						
					}
				}
			}
		}
	}
}