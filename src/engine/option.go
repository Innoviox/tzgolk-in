package engine

import (
    "fmt"
)

// type Option struct {
// 	Execute func(*Game, *Player)
// 	Description string
// 	BuildingNum int
// }

type Options func(*Game, *Player) []*Delta

func Flatten(options []Options) Options {
	// todo add "mirror" to Description?
	return func (g *Game, p *Player) []*Delta {
		result := make([]*Delta, 0)
		for _, o := range options {
			result = append(result, o(g, p)...)
		}
		return result
	}
}

func Skip() []*Delta {
	return []*Delta{&Delta{Description: "skip"}}
}

func SkipWrapper(o []*Delta) []*Delta {
	if len(o) == 0 {
		return Skip()
	} else {
		return o
	}
}

func (g *Game) GetBuildingOptions(p *Player, exclude int, useResearch bool) []*Delta {
	options := make([]*Delta, 0)

	for _, b := range g.CurrentBuildings {
		if b.Id == exclude {
			continue
		}

		costs := b.GetCosts(g, p)
		for _, cost := range costs {
			for _, effect := range b.GetEffects(g, p) {
				d := &Delta{
					PlayerDeltas: map[Color]PlayerDelta{
						p.Color: PlayerDelta{
							Resources: InvCost(b.Cost),
							// todo building delta
						},
					},
					Description: fmt.Sprintf("[build %d] pay %s ", b.Id, CostString(cost)),
					BuildingNum: b.Id,
				}
				d.Add(effect) // effect.Descriptio
				if useResearch {
					d.Add(g.Research.Built(p)) // g.Research.BuiltString(p)
				}
				options = append(options, d)
			}
		}
	}

	return options
}

func (g *Game) GetMonumentOptions(p *Player) []Delta {
	options := make([]Delta, 0)

	for _, m := range g.CurrentMonuments {
		if p.CanPay(m.Cost) {
			options = append(options, Delta{
				PlayerDeltas: map[Color]PlayerDelta{
					p.Color: PlayerDelta{
						Resources: InvCost(m.Cost),
						// todo monument delta
					},
				},
				Description: fmt.Sprintf("[build %d] pay %s, get monument %d", m.Id, CostString(m.Cost), m.Id),
			})
		}
	}

	return options
}

func (t *Temples) GainTempleStep(p *Player, o *Delta, dir int) []*Delta {
	options := make([]*Delta, 0)

	for i := 0; i < 3; i++ {
		if t.CanStep(p, i, dir) {
			d := t.Step(p, i, dir)
			d.Add(o)
			options = append(options, d)
		}
	}

	return options
}

func (r *Research) GetOptions(g *Game, p *Player, n int, free bool) []*Delta{ 
	resources := [4]int{}
	for i := 0; i < 4; i++ {
		resources[i] = p.Resources[i]
	}

	return r.GetOptionsHelper(g, p, resources, r.Levels[p.Color], n, free)
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
	Descriptions := make([]string, 0)
	for s := 0; s < 4; s++ {
		if nl[Science(s)] > l[Science(s)] {
			Descriptions = append(Descriptions, fmt.Sprintf("%s %d", string(ResearchDebug[s]), nl[Science(s)] - l[Science(s)]))
		}
	}
	return fmt.Sprintf("%v", Descriptions)
}

func (r *Research) GetOptionsHelper(g *Game, p *Player, resources [4]int, levels Levels, n int, free bool) []*Delta {
	options := make([]*Delta, 0)
	for s := 0; s < 4; s++ {
		level := levels[Science(s)]
		possResources := [][4]int{resources} // the resources you end up with after paying
											 // (it's current resources since it's free)
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

				d := ResourcesDelta(p.Color, p.Resources, newResources)
				d.Add(&Delta{ResearchDelta: ResearchDelta{Levels: map[Color]Levels{p.Color: map[Science]int{
					Science(s): 1,
				}}}})
				d.Description = GenerateResearchDescription(resources, newResources, levels, newLevels)

				if n == 1 {
					options = append(options, d)
				} else {
					for _, o := range r.GetOptionsHelper(g, p, newResources, newLevels, n - 1, free) {
						options = append(options, Combine(d, o))
					}
				}
			}
		} else {
			advancedOptions := r.GetAdvancedOptions(g, p, resources, free, Science(s))

			if n == 1 {
				options = append(options, advancedOptions...)
			} else {
				for _, o1 := range advancedOptions {
					for _, o2 := range r.GetOptionsHelper(g, p, resources, levels, n - 1, free) {
						options = append(options, Combine(o1, o2))
					}
				}
			}
		}
	}

	return options
}

func (r *Research) GetAdvancedOptions(g *Game, p *Player, resources [4]int, free bool, s Science) []*Delta {
	advancedOptions := make([]*Delta, 0)
	possResources := [][4]int{resources}
	if !free {
		possResources = PayBlocks(resources, 1)
	}
	for _, newResources := range possResources {
		switch Science(s) {
		case Agriculture:
			d := ResourcesDelta(p.Color, p.Resources, newResources)
			d.Description = fmt.Sprintf("[agr tier 4] pay %s", GeneratePaymentDescription(resources, newResources))
			advancedOptions = append(advancedOptions, g.Temples.GainTempleStep(p, d, 1)...)
		case Resources:
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					advancedOptions = append(advancedOptions, Option{
						Execute: func(g *Game, p *Player) {
							p.Resources = newResources
							p.Resources[i] += 1
							p.Resources[j] += 1
						},
						Description: fmt.Sprintf("[res tier 4] pay %s, 1 %s, 1 %s", GeneratePaymentDescription(resources, newResources), string(ResourceDebug[i]), string(ResourceDebug[j])),
					})
				}
			}
		case Construction:
			advancedOptions = append(advancedOptions, Option{
				Execute: func(g *Game, p *Player) {
					p.Resources = newResources
					p.Points += 3
				},
				Description: fmt.Sprintf("[cons tier 4] pay %s, 3 points", GeneratePaymentDescription(resources, newResources)),
			})
		case Theology:
			advancedOptions = append(advancedOptions, Option{
				Execute: func(g *Game, p *Player) {
					p.Resources = newResources
					p.Resources[Skull] += 1
				},
				Description: fmt.Sprintf("[theo tier 4] pay %s, 1 skull", GeneratePaymentDescription(resources, newResources)),
			})
		}
	}
	return advancedOptions
}


func (r *Research) FreeResearch(g *Game, p *Player, s Science) []*Delta {
	if g.Research.HasLevel(p.Color, s, 3) {
		return r.GetAdvancedOptions(g, p, p.Resources, true, s)
	} else {
		return []*Delta{&Delta{
			ResearchDelta: ResearchDelta{
				Levels: map[Color]Levels{
					p.Color: Levels{
						s: 1,
					},
				},
			},
			Description: fmt.Sprintf("free %s", string(ResearchDebug[s])),
		}}
	}
}