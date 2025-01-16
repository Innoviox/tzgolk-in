package buildings

import (
    // "fmt"
    // "math/rand"
    . "tzgolkin/engine"
)

func Building1() Building {
    return Building {
        Id: 1,
        Cost: [4]int{1, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            d := PlayerDeltaWrapper(p.Color, PlayerDelta{
                FreeWorkers: 1,
            })
            d.Description = "1 free worker"
            return []*Delta{d}
        },
        Color: Yellow,
    }
}

func Building2() Building {
    return Building {
        Id: 2,
        Cost: [4]int{1, 2, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            return []*Delta{Combine(
                g.Temples.Step(p, 0, 1),
                g.Temples.Step(p, 2, 1),
            )}
        },
        Color: Red,
    }
}

func Building3() Building {
    b := Building1()
    return Building {
        Id: 3,
        Cost: b.Cost,
        GetEffects: b.GetEffects,
        Color: b.Color,
    }
}

func Building4() Building {
    return Building {
        Id: 4,
        Cost: [4]int{2, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            return g.Research.FreeResearch(g, p, Agriculture)
        },
        Color: Green,
    }
}

func Building5() Building {
    return Building {
        Id: 5,
        Cost: [4]int{4, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            d := PlayerDeltaWrapper(p.Color, PlayerDelta{
                WorkerDeduction: 1,
            })
            d.Description = "1 worker deduction"
            return []*Delta{d}
        },
        Color: Yellow,
    }
}

func Building6() Building {
    return Building {
        Id: 6,
        Cost: [4]int{0, 0, 1, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            return g.Research.FreeResearch(g, p, Construction)
        },
        Color: Blue,
    }
}

func Building7() Building {
    return Building {
        Id: 7,
        Cost: [4]int{1, 0, 1, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)

            new_player := p.Clone()
            new_player.Resources[Wood] -= 1
            new_player.Resources[Gold] -= 1
            
            for _, o := range g.GetBuildingOptions(new_player, 7, true) {
                options = append(options, g.Temples.GainTempleStep(new_player, o, 1)...)
            }

            return options
        },
        Color: Red,
    }
}

func Building8() Building {
    return Building {
        Id: 8,
        Cost: [4]int{2, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            return []*Delta{Combine(
                g.Temples.Step(p, 0, 1),
                g.Temples.Step(p, 1, 1),
            )}
        },
        Color: Red,
    }
}

func Building9() Building {
    return Building {
        Id: 9,
        Cost: [4]int{1, 1, 0, 0},   
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
            for _, o := range g.Research.FreeResearch(g, p, Resources) {
                options = append(options, Combine(o, PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Corn: 1,
                })))
            }
            return options
        },
        Color: Green,
    }
}

func Building10() Building {
    return Building {
        Id: 10,
        Cost: [4]int{0, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
            for _, o := range g.Research.FreeResearch(g, p, Theology) {
                options = append(options, Combine(o, g.Temples.Step(p, 2, 1)))
            }
            return options
        },
        Color: Blue,
    }
}

func Building11() Building {
    b := Building5()
    return Building {
        Id: 11,
        Cost: b.Cost,
        GetEffects: b.GetEffects,
        Color: b.Color,
    }
}

func Building12() Building {
    return Building {
        Id: 12,
        Cost: [4]int{2, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
            for _, o := range g.Research.FreeResearch(g, p, Resources) {
                options = append(options, Combine(o, PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Resources: [4]int{0, 0, 1, 0},
                })))
            }
            return options
        },
        Color: Green,
    }
}

func Building13() Building {
    return Building {
        Id: 13,
        Cost: [4]int{3, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
            for _, o := range g.Research.FreeResearch(g, p, Agriculture) {
                options = append(options, Combine(o, PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Resources: [4]int{0, 1, 0, 0},
                })))
            }
            return options
        },
        Color: Green,
    }
}

func Building14() Building {
    b := Building1()
    return Building {
        Id: 14,
        Cost: b.Cost,
        GetEffects: b.GetEffects,
        Color: Yellow,
    }
}

func MakeAge1Buildings() map[int]Building {
    buildings := make([]Building, 0)

    buildings = append(buildings, Building1())
    buildings = append(buildings, Building2())
    buildings = append(buildings, Building3())
    buildings = append(buildings, Building4())
    buildings = append(buildings, Building5())
    buildings = append(buildings, Building6())
    buildings = append(buildings, Building7())
    buildings = append(buildings, Building8())
    buildings = append(buildings, Building9())
    buildings = append(buildings, Building10())
    buildings = append(buildings, Building11())
    buildings = append(buildings, Building12())
    buildings = append(buildings, Building13())
    buildings = append(buildings, Building14())

    b := map[int]Building{}
    for _, building := range buildings {
        b[building.Id] = building
    }
    return b
}