package buildings

import (
    "math/rand"
    . "tzgolkin/model"
)

func Building1() Building {
    return Building {
        Id: 1,
        Cost: [4]int{1, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    p.FreeWorkers += 1
                },
                Description: "1 free worker",
            }}
        },
        Color: Yellow,
    }
}

func Building2() Building {
    return Building {
        Id: 2,
        Cost: [4]int{1, 2, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Temples.Step(p, 0, 1)
                    g.Temples.Step(p, 2, 1)
                },
                Description: "1 BT, 1 GT",
            }}
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
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Research.FreeResearch(p.Color, Agriculture)
                },
                Description: "free agr",
            }}
        },
        Color: Green,
    }
}

func Building5() Building {
    return Building {
        Id: 5,
        Cost: [4]int{4, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    p.WorkerDeduction += 1
                }, 
                Description: "1 worker deduction",
            }}
        },
        Color: Yellow,
    }
}

func Building6() Building {
    return Building {
        Id: 6,
        Cost: [4]int{0, 0, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                   g.Research.FreeResearch(p.Color, Construction)
                },
                Description: "free const",
            }}
        },
        Color: Blue,
    }
}

func Building7() Building {
    return Building {
        Id: 7,
        Cost: [4]int{1, 0, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            
            for _, o := range g.GetBuildingOptions(p, 7, true) {
                options = append(options, g.Temples.GainTempleStep(p, o, 1)...)
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
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Temples.Step(p, 0, 1)
                    g.Temples.Step(p, 1, 1)
                },
                Description: "1 BT 1 YT",
            }}
        },
        Color: Red,
    }
}

func Building9() Building {
    return Building {
        Id: 9,
        Cost: [4]int{1, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Research.FreeResearch(p.Color, Resources)
                    p.Corn += 1
                },
                Description: "free res, 1 Corn",
            }}
        },
        Color: Green,
    }
}

func Building10() Building {
    return Building {
        Id: 10,
        Cost: [4]int{0, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Research.FreeResearch(p.Color, Theology)
                    g.Temples.Step(p, 2, 1)
                },
                Description: "free theo, 1 GT",
            }}
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
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Research.FreeResearch(p.Color, Resources)
                    p.Resources[Gold] += 1
                },
                Description: "free res, 1 G",
            }}
        },
        Color: Green,
    }
}

func Building13() Building {
    return Building {
        Id: 13,
        Cost: [4]int{3, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Research.FreeResearch(p.Color, Agriculture)
                    p.Resources[Stone] += 1
                },
                Description: "free agr, 1 S",
            }}
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

func MakeAge1Buildings(r *rand.Rand) []Building {
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

    // shuffle buildings
	for i := range buildings {
		j := r.Intn(i + 1)
		buildings[i], buildings[j] = buildings[j], buildings[i]
	}

    return buildings
}