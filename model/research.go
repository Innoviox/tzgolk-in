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

func (r *Research) GetOptions(g *Game, p *Player, n int) []Option{ 
	resources := [4]int{}
	for i := 0; i < 4; i++ {
		resources[i] = p.resources[i]
	}

	return r.GetOptionsHelper(g, p, resources, r.levels[p.color], n)
}

func (r *Research) GetOptionsHelper(g *Game, p *Player, resources [4]int, levels map[Science]int, n int) []Option {
	options := make([]Option, 0)
	for s := 0; s < 4; s++ {
		level := levels[Science(s)]
		if level < 3 {
			for _, newResources := range PayBlocks(resources, level) {
				newLevels := make(map[Science]int)
				for k, v := range levels {
					newLevels[k] = v
				}
				newLevels[Science(s)] += 1

				if n == 1 {
					options = append(options, func() {
						p.resources = newResources
						r.levels[p.color] = newLevels
					})
				} else {
					options = append(options, r.GetOptionsHelper(g, p, newResources, newLevels, n - 1)...)
				}
			}
		} else {
			advancedOptions := make([]Option, 0)
			for _, newResources := range PayBlocks(resources, 1) {
				switch Science(s) {
				case Agriculture:
					advancedOptions = append(advancedOptions, g.temples.GainTempleStep(p.color, func() {
						p.resources = newResources
					}, 1)...)
				case Resources:
					for i := 0; i < 3; i++ {
						for j := 0; j < 3; j++ {
							advancedOptions = append(advancedOptions, func() {
								p.resources = newResources
								p.resources[i] += 1
								p.resources[j] += 1
							})
						}
					}
				case Construction:
					advancedOptions = append(advancedOptions, func() {
						p.resources = newResources
						p.points += 3
					})
				case Theology:
					advancedOptions = append(advancedOptions, func() {
						p.resources = newResources
						p.resources[Skull] += 1
					})
				}
			}

			if n == 1 {
				options = append(options, advancedOptions...)
			} else {
				for _, o1 := range advancedOptions {
					for _, o2 := range r.GetOptionsHelper(g, p, resources, levels, n - 1) {
						options = append(options, func() {
							o1()
							o2()
						})
					}
				}
			}
		}
	}

	return options
}