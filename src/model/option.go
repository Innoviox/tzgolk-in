package model

import (
    "fmt"
)

type Option struct {
	Execute func(*Game, *Player)
	Description string
	BuildingNum int
}

type Options func(*Game, *Player) []Option

func Flatten(options []Options) Options {
	// todo add "mirror" to Description?
	return func (g *Game, p *Player) []Option {
		result := make([]Option, 0)
		for _, o := range options {
			result = append(result, o(g, p)...)
		}
		return result
	}
}

func (g *Game) GetBuildingOptions(p *Player, exclude int, useResearch bool) []Option {
	options := make([]Option, 0)

	for _, b := range g.CurrentBuildings {
		if b.Id == exclude {
			continue
		}

		costs := b.GetCosts(g, p)
		for _, cost := range costs {
			for _, effect := range b.GetEffects(g, p) {
				options = append(options, Option{
					Execute: func(g *Game, p *Player) {
						for i := 0; i < 4; i++ {
							p.Resources[i] -= cost[i]
						}

						effect.Execute(g, p)

						if useResearch {
							g.Research.Built(p)
						}

						p.Buildings = append(p.Buildings, b)

						// g.RemoveBuilding(b)
					},
					Description: fmt.Sprintf("[build %d] pay %s, %s +%s", b.Id, CostString(cost), effect.Description, g.Research.BuiltString(p)),
					BuildingNum: b.Id,
				})
			}
		}
	}

	return options
}

func (g *Game) GetMonumentOptions(p *Player) []Option {
	options := make([]Option, 0)

	for _, m := range g.CurrentMonuments {
		if p.CanPay(m.Cost) {
			options = append(options, Option{
				Execute: func(g *Game, p *Player) {
					for i := 0; i < 4; i++ {
						p.Resources[i] -= m.Cost[i]
					}

					p.Monuments = append(p.Monuments, m)

					// g.RemoveMonument(m)
				},
				Description: fmt.Sprintf("[build %d] pay %s, get monument %d", m.Id, CostString(m.Cost), m.Id),
			})
		}
	}

	return options
}

func (t *Temples) GainTempleStep(p *Player, o Option, dir int) []Option {
	options := make([]Option, 0)

	for i := 0; i < 3; i++ {
		if t.CanStep(p, i, dir) {
			options = append(options, Option{
				Execute: func(g *Game, p *Player) {
					g.Temples.Step(p, i, dir)
					o.Execute(g, p)
				},
				Description: fmt.Sprintf("%s, %s temple %d", o.Description, p.Color.String(), dir),
			})
		}
	}

	return options
}

func (r *Research) GetOptions(g *Game, p *Player, n int, free bool) []Option{ 
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

func (r *Research) GetOptionsHelper(g *Game, p *Player, resources [4]int, levels Levels, n int, free bool) []Option {
	options := make([]Option, 0)
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

				opt := Option{
					Execute: func(g *Game, p *Player) {
						p.Resources = newResources
						g.Research.Levels[p.Color] = newLevels
					},
					Description: GenerateResearchDescription(resources, newResources, levels, newLevels),
				}

				if n == 1 {
					options = append(options, opt)
				} else {
					for _, o := range r.GetOptionsHelper(g, p, newResources, newLevels, n - 1, free) {
						options = append(options, Option{
							Execute: func(g *Game, p *Player) {
								opt.Execute(g, p)
								o.Execute(g, p)
							},
							Description: fmt.Sprintf("%s; %s", opt.Description, o.Description),
						})
					}
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
					advancedOptions = append(advancedOptions, g.Temples.GainTempleStep(p, Option{
						Execute: func(g *Game, p *Player) {
							p.Resources = newResources
						},
						Description: fmt.Sprintf("[agr tier 4] pay %s", GeneratePaymentDescription(resources, newResources)),
					}, 1)...)
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
							Description: fmt.Sprintf("%s; %s", o1.Description, o2.Description),
						})
					}
				}
			}
		}
	}

	return options
}