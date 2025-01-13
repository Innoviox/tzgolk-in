package buildings

import (
    "math/rand"
    . "tzgolkin/model"
)

func Building1() Building {
    return MakeBuilding(
        1,
        [4]int{1, 0, 0, 0},
        func (g *Game, p *Player) []Option {
            return []Option {MakeOption(
                func(g *Game, p *Player) {
                    p.freeWorkers += 1
                },
                "1 free worker",
            )}
        },
        Yellow,
    )
}

func Building2() Building {
    return Building {
        id: 2,
        cost: [4]int{1, 2, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.temples.Step(p, 0, 1)
                    g.temples.Step(p, 2, 1)
                },
                description: "1 BT, 1 GT",
            }}
        },
        color: Red,
    }
}

func Building3() Building {
    b := Building1()
    return Building {
        id: 3,
        cost: b.cost,
        GetEffects: b.GetEffects,
        color: b.color,
    }
}

func Building4() Building {
    return Building {
        id: 4,
        cost: [4]int{2, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.research.FreeResearch(p.color, Agriculture)
                },
                description: "free agr",
            }}
        },
        color: Green,
    }
}

func Building5() Building {
    return Building {
        id: 5,
        cost: [4]int{4, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    p.workerDeduction += 1
                }, 
                description: "1 worker deduction",
            }}
        },
        color: Yellow,
    }
}

func Building6() Building {
    return Building {
        id: 6,
        cost: [4]int{0, 0, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                   g.research.FreeResearch(p.color, Construction)
                },
                description: "free const",
            }}
        },
        color: Blue,
    }
}

func Building7() Building {
    return Building {
        id: 7,
        cost: [4]int{1, 0, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            
            for _, o := range g.GetBuildingOptions(p, 7, true) {
                options = append(options, g.temples.GainTempleStep(p, o, 1)...)
            }

            return options
        },
        color: Red,
    }
}

func Building8() Building {
    return Building {
        id: 8,
        cost: [4]int{2, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.temples.Step(p, 0, 1)
                    g.temples.Step(p, 1, 1)
                },
                description: "1 BT 1 YT",
            }}
        },
        color: Red,
    }
}

func Building9() Building {
    return Building {
        id: 9,
        cost: [4]int{1, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.research.FreeResearch(p.color, Resources)
                    p.corn += 1
                },
                description: "free res, 1 corn",
            }}
        },
        color: Green,
    }
}

func Building10() Building {
    return Building {
        id: 10,
        cost: [4]int{0, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.research.FreeResearch(p.color, Theology)
                    g.temples.Step(p, 2, 1)
                },
                description: "free theo, 1 GT",
            }}
        },
        color: Blue,
    }
}

func Building11() Building {
    b := Building5()
    return Building {
        id: 11,
        cost: b.cost,
        GetEffects: b.GetEffects,
        color: b.color,
    }
}

func Building12() Building {
    return Building {
        id: 12,
        cost: [4]int{2, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.research.FreeResearch(p.color, Resources)
                    p.resources[Gold] += 1
                },
                description: "free res, 1 G",
            }}
        },
        color: Green,
    }
}

func Building13() Building {
    return Building {
        id: 13,
        cost: [4]int{3, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.research.FreeResearch(p.color, Agriculture)
                    p.resources[Stone] += 1
                },
                description: "free agr, 1 S",
            }}
        },
        color: Green,
    }
}

func Building14() Building {
    b := Building1()
    return Building {
        id: 14,
        cost: b.cost,
        GetEffects: b.GetEffects,
        color: Yellow,
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