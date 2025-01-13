package model

import (
	"fmt"
	"strings"
)

type Science int
type Levels map[Science]int

// todo get actual names
const (
	Agriculture Science = iota
	Resources
	Construction
	Theology
)

const ResearchDebug = "ARCT"


type Research struct {
	levels map[Color]Levels
}

func (r *Research) Clone() *Research {
	var newLevels = make(map[Color]Levels)

	for k, v := range r.levels {
		newSci := make(Levels)
		for k2, v2 := range v {
			newSci[k2] = v2
		}
		newLevels[k] = newSci
	}

	return &Research {
		levels: newLevels,
	}
}

func MakeLevels() Levels {
	return Levels{
		Agriculture: 0,
		Resources: 0,
		Construction: 0,
		Theology: 0,
	}
}

func MakeResearch() *Research {
	return &Research{
		levels: map[Color]Levels{
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

func (r *Research) Foresight(c Color) bool {
	return r.HasLevel(c, Theology, 1)
}

func (r *Research) CornBonus(player Color, positionColor Color) int {
	if positionColor == Blue {
		if r.HasLevel(player, Agriculture, 2) {
			return 1
		} else {
			return 0
		}
	} else if positionColor == Green {
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
	if r.HasLevel(c, Resources, int(res) + 1) {
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
		p.points += 2
	}
}

func (r *Research) BuiltString(p *Player) string {
	var br strings.Builder

	if r.HasLevel(p.color, Construction, 1) {
		fmt.Fprintf(&br, "1 corn ")
	}

	if r.HasLevel(p.color, Construction, 2) {
		fmt.Fprintf(&br, "2 points")
	}

	return br.String()
}

func (r *Research) Builder(c Color) bool {
	return r.HasLevel(c, Construction, 3)
}

func (r *Research) GetOptions(g *Game, p *Player, n int, free bool) []Option{ 
	resources := [4]int{}
	for i := 0; i < 4; i++ {
		resources[i] = p.resources[i]
	}

	return r.GetOptionsHelper(g, p, resources, r.levels[p.color], n, free)
}

func GenerateResearchDescription(r [4]int, nr [4]int, l Levels, nl Levels) string {
	return fmt.Sprintf("pay %s, gain %s", GeneratePaymentDescription(r, nr), GenerateLevelsDescription(l, nl))
}

func GeneratePaymentDescription(r [4]int, nr [4]int) string {
	payments := make([]string, 0)
	for i := 0; i < 4; i++ {
		if nr[i] < r[i] {
			payments = append(payments, fmt.Sprintf("%d %s", r[i] - nr[i], string(ResourceDebug[i])))
		}
	}
	return fmt.Sprintf("%v", payments)
}

func GenerateLevelsDescription(l Levels, nl Levels) string {
	descriptions := make([]string, 0)
	for s := 0; s < 4; s++ {
		if nl[Science(s)] > l[Science(s)] {
			descriptions = append(descriptions, fmt.Sprintf("%s %d", string(ResearchDebug[s]), nl[Science(s)] - l[Science(s)]))
		}
	}
	return fmt.Sprintf("%v", descriptions)
}

func (r *Research) GetOptionsHelper(g *Game, p *Player, resources [4]int, levels Levels, n int, free bool) []Option {
	options := make([]Option, 0)
	for s := 0; s < 4; s++ {
		level := levels[Science(s)]
		possResources := [][4]int{resources}
		if level < 3 {
			if !free {
				possResources = PayBlocks(resources, level + 1)
			}
			for _, newResources := range possResources {
				newLevels := make(Levels)
				for k, v := range levels {
					newLevels[k] = v
				}
				newLevels[Science(s)] += 1

				if n == 1 {
					options = append(options, Option{
						Execute: func(g *Game, p *Player) {
							p.resources = newResources
							g.research.levels[p.color] = newLevels
						},
						description: GenerateResearchDescription(resources, newResources, levels, newLevels),
					})
				} else {
					options = append(options, r.GetOptionsHelper(g, p, newResources, newLevels, n - 1, free)...)
				}
			}
		} else {
			advancedOptions := make([]Option, 0)
			if !free {
				possResources = PayBlocks(resources, 1)
			}
			for _, newResources := range possResources {
				switch Science(s) {
				case Agriculture:
					advancedOptions = append(advancedOptions, g.temples.GainTempleStep(p, Option{
						Execute: func(g *Game, p *Player) {
							p.resources = newResources
						},
						description: fmt.Sprintf("[agr tier 4] pay %s", GeneratePaymentDescription(resources, newResources)),
					}, 1)...)
				case Resources:
					for i := 0; i < 3; i++ {
						for j := 0; j < 3; j++ {
							advancedOptions = append(advancedOptions, Option{
								Execute: func(g *Game, p *Player) {
									p.resources = newResources
									p.resources[i] += 1
									p.resources[j] += 1
								},
								description: fmt.Sprintf("[res tier 4] pay %s, 1 %s, 1 %s", GeneratePaymentDescription(resources, newResources), string(ResourceDebug[i]), string(ResourceDebug[j])),
							})
						}
					}
				case Construction:
					advancedOptions = append(advancedOptions, Option{
						Execute: func(g *Game, p *Player) {
							p.resources = newResources
							p.points += 3
						},
						description: fmt.Sprintf("[cons tier 4] pay %s, 3 points", GeneratePaymentDescription(resources, newResources)),
					})
				case Theology:
					advancedOptions = append(advancedOptions, Option{
						Execute: func(g *Game, p *Player) {
							p.resources = newResources
							p.resources[Skull] += 1
						},
						description: fmt.Sprintf("[theo tier 4] pay %s, 1 skull", GeneratePaymentDescription(resources, newResources)),
					})
				}
			}

			if n == 1 {
				options = append(options, advancedOptions...)
			} else {
				for _, o1 := range advancedOptions {
					for _, o2 := range r.GetOptionsHelper(g, p, resources, levels, n - 1, free) {
						options = append(options, Option{
							Execute: func(g *Game, p *Player) {
								o1.Execute(g, p)
								o2.Execute(g, p)
							},
							description: fmt.Sprintf("%s; %s", o1.description, o2.description),
						})
					}
				}
			}
		}
	}

	return options
}

func (r *Research) FreeResearch(c Color, s Science) {
	r.levels[c][s] += 1
}

func (r *Research) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "----Research------------\n")

	var matrix [][]strings.Builder

	for j := 0; j < 4; j++ {
		row := make([]strings.Builder, 0)
		for k := 0; k < 4; k++ {
			row = append(row, strings.Builder{})
		}
		matrix = append(matrix, row)
	}

	for c, l := range r.levels {
		for s, n := range l {
			fmt.Fprintf(&matrix[int(s)][n], "%s", c.String())
		}
	}
	
	for i := 0; i < 4; i++ {
		fmt.Fprintf(&br, "|%s |", string(ResearchDebug[i]))

		for j := 0; j < 4; j++ {
			fmt.Fprintf(&br, "%-4s|", matrix[i][j].String())
		}

		fmt.Fprintf(&br, "\n")
	}

	fmt.Fprintf(&br, "------------------------\n")

	matrix = nil

	return br.String()
}