package model

import (
    "math/rand"
)

func Building1() Building {
    return Building {
        id: 1,
        cost: [4]int{1, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func() {
                    p.freeWorkers += 1
                },
                description: "1 free worker",
            }}
        },
    }
}

func Building2() Building {
    return Building {
        id: 2,
        cost: [4]int{1, 2, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func() {
                    g.temples.Step(p, 0, 1)
                    g.temples.Step(p, 2, 1)
                },
                description: "1 BT, 1 GT",
            }}
        },
    }
}

func Building3() Building {
    b := Building1()
    return Building {
        id: 3,
        cost: b.cost,
        GetEffects: b.GetEffects,
    }
}

func Building4() Building {
    return Building {
        id: 4,
        cost: [4]int{2, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func() {
                    g.research.FreeResearch(p.color, Agriculture)
                },
                description: "free agr",
            }}
        },
    }
}

func Building5() Building {
    return Building {
        id: 5,
        cost: [4]int{4, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func() {
                    p.workerDeduction += 1
                }, 
                description: "1 worker deduction",
            }}
        },
    }
}

func Building6() Building {
    return Building {
        id: 6,
        cost: [4]int{0, 0, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func() {
                   g.research.FreeResearch(p.color, Construction)
                },
                description: "free const",
            }}
        },
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
    }
}

func Building8() Building {
    return Building {
        id: 8,
        cost: [4]int{2, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func() {
                    g.temples.Step(p, 0, 1)
                    g.temples.Step(p, 1, 1)
                },
                description: "1 BT 1 YT",
            }}
        },
    }
}

func Building9() Building {
    return Building {
        id: 9,
        cost: [4]int{1, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func() {
                    g.research.FreeResearch(p.color, Resources)
                    p.corn += 1
                },
                description: "free res, 1 corn",
            }}
        },
    }
}

func Building10() Building {
    return Building {
        id: 10,
        cost: [4]int{0, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func() {
                    g.research.FreeResearch(p.color, Theology)
                    g.temples.Step(p, 2, 1)
                },
                description: "free theo, 1 GT",
            }}
        },
    }
}

func Building11() Building {
    b := Building5()
    return Building {
        id: 11,
        cost: b.cost,
        GetEffects: b.GetEffects,
    }
}

func Building12() Building {
    return Building {
        id: 12,
        cost: [4]int{2, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func() {
                    g.research.FreeResearch(p.color, Resources)
                    p.resources[Gold] += 1
                },
                description: "free res, 1 G",
            }}
        },
    }
}

func Building13() Building {
    return Building {
        id: 13,
        cost: [4]int{3, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func() {
                    g.research.FreeResearch(p.color, Agriculture)
                    p.resources[Stone] += 1
                },
                description: "free agr, 1 S",
            }}
        },
    }
}

func Building14() Building {
    b := Building1()
    return Building {
        id: 14,
        cost: b.cost,
        GetEffects: b.GetEffects,
    }
}

func MakeAge1Buildings() []Building {
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
		j := rand.Intn(i + 1)
		buildings[i], buildings[j] = buildings[j], buildings[i]
	}

    return buildings
}