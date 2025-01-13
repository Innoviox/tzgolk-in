package buildings

import (
    "fmt"
    "math/rand"
    . "tzgolkin/model"
)

func Age2Building1() Building {
    return Building {
        id: 1,
        cost: [4]int{0, 0, 2, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.temples.Step(p, 2, 2)
                    p.points += 3
                },
                description: "2 GT, 3 points",
            }}
        },
        color: Blue,
    }
}

func Age2Building2() Building {
    return Building {
        id: 2,
        cost: [4]int{0, 2, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.temples.Step(p, 0, 2)
                    p.points += 2
                },
                description: "2 BT, 2 points",
            }}
        },
        color: Blue,
    }
}

func Age2Building3() Building {
    return Building {
        id: 3,
        cost: [4]int{0, 0, 3, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.temples.Step(p, 1, 2)
                },
                description: "2 YT, 4 points",
            }}
        },
        color: Blue,
    }
}

func Age2Building4() Building {
    return Building {
        id: 4,
        cost: [4]int{0, 1, 2, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return g.research.GetOptions(g, p, 2, true)
        },
        color: Blue,
    }
}

func Age2Building5() Building {
    return Building {
        id: 5,
        cost: [4]int{0, 2, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.research.FreeResearch(p.color, Theology)
                    g.temples.Step(p, 0, 1)
                    g.temples.Step(p, 2, 1)
                },
                description: "free theo, 1 BT, 1 GT",
            }}
        },
        color: Blue,
    }
}

func Age2Building6() Building {
    return Building {
        id: 6,
        cost: [4]int{3, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            for _, o := range g.research.GetOptions(g, p, 1, true) {
                options = append(options, Option{
                    Execute: func(g *Game, p *Player) {
                        o.Execute(g, p)
                        p.resources[Stone] += 1
                    },
                    description: fmt.Sprintf("%s, 1 stone", o.description),
                })
            }

            return options
        },
        color: Green,
    }
}

func Age2Building7() Building {
    return Building {
        id: 7,
        cost: [4]int{1, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.UnlockWorker(p.color)
                    p.points += 6
                },
                description: "unlock worker, 6 points",
            }}
        },
        color: Red,

    }
}

func Age2Building8() Building {
    return Building {
        id: 8,
        cost: [4]int{0, 0, 3, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            for _, o := range Uxmal2(g, p) {
                options = append(options, Option{
                    Execute: func(g *Game, p *Player) {
                        o.Execute(g, p)
                        p.points += 6
                    },
                    description: fmt.Sprintf("%s, 6 points", o.description),
                })
            }

            options = append(options, Option{
                Execute: func(g *Game, p *Player) {
                    p.points += 6
                },
                description: "6 points",
            })

            return options
        },
        color: Red,
    }
}

func Age2Building9() Building {
    return Building {
        id: 9,
        cost: [4]int{1, 0, 2, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    p.points += 8
                },
                description: "8 points",
            }}
        },
        color: Red,
    }
}

func Age2Building10() Building {
    return Building {
        id: 10,
        cost: [4]int{2, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    p.freeWorkers += 3
                },
                description: "3 free workers",
            }}
        },
        color: Yellow,
    }
}

func Age2Building11() Building {
    return Building {
        id: 11,
        cost: [4]int{1, 2, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    for i := 0; i < 3; i++ {
                        g.temples.Step(p, i, 1)
                    }
                    p.points += 3
                },
                description: "1 BT 1 GT 1 YT, 3 points",
            }}
        },
        color: Red,
    }
}

func Age2Building12() Building {
    return Building {
        id: 12,
        cost: [4]int{2, 2, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            for _, o := range g.research.GetOptions(g, p, 1, true) {
                options = append(options, Option{
                    Execute: func(g *Game, p *Player) {
                        o.Execute(g, p)
                        p.resources[Skull] += 1
                    },
                    description: fmt.Sprintf("%s, 1 skull", o.description),
                })
            }
            return options
        },
        color: Green,
    }
}

func Age2Building13() Building {
    return Building {
        id: 13,
        cost: [4]int{0, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.research.FreeResearch(p.color, Construction)
                    p.points += 3
                },
                description: "free const, 3 points",
            }}
        },
        color: Blue,
    }
}

func Age2Building14() Building {
    b := Age2Building10()
    return Building {
        id: 14,
        cost: b.cost,
        GetEffects: b.GetEffects,
        color: Yellow,
    }
}

func Age2Building15() Building {
    b := Age2Building10()
    return Building {
        id: 15,
        cost: b.cost,
        GetEffects: b.GetEffects,
        color: Yellow,
    }
}

func Age2Building16() Building {
    return Building {
        id: 16,
        cost: [4]int{2, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            // for _, o := range Uxmal5(g, p) {
            //     options = append(options, Option{
            //         Execute: func(g *Game, p *Player) {
            //             o.Execute(g, p)
            //             p.points += 2
            //         },
            //         description: fmt.Sprintf("%s, 2 points", o.description),
            //     })
            // }
            return options
        },
        color: Red,
    }
}

func Age2Building17() Building {
    return Building {
        id: 17,
        cost: [4]int{0, 0, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            for _, o := range g.research.GetOptions(g, p, 1, true) {
                options = append(options, Option{
                    Execute: func(g *Game, p *Player) {
                        o.Execute(g, p)
                        p.corn += 6
                    },
                    description: fmt.Sprintf("%s, 6 corn", o.description),
                })
            }
            return options
        },
        color: Green,
    }
}

func Age2Building18() Building {
    return Building {
        id: 18,
        cost: [4]int{2, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            for _, o := range g.research.GetOptions(g, p, 1, true) {
                options = append(options, Option{
                    Execute: func(g *Game, p *Player) {
                        o.Execute(g, p)
                        p.resources[Gold] += 1
                    },
                    description: fmt.Sprintf("%s, 1 G", o.description),
                })
            }
            return options
        },
        color: Green,
    }
}

func MakeAge2Buildings(r *rand.Rand) []Building {
    buildings := make([]Building, 0)

    buildings = append(buildings, Age2Building1())
    buildings = append(buildings, Age2Building2())
    buildings = append(buildings, Age2Building3())
    buildings = append(buildings, Age2Building4())
    buildings = append(buildings, Age2Building5())
    buildings = append(buildings, Age2Building6())
    buildings = append(buildings, Age2Building7())
    buildings = append(buildings, Age2Building8())
    buildings = append(buildings, Age2Building9())
    buildings = append(buildings, Age2Building10())
    buildings = append(buildings, Age2Building11())
    buildings = append(buildings, Age2Building12())
    buildings = append(buildings, Age2Building13())
    buildings = append(buildings, Age2Building14())
    buildings = append(buildings, Age2Building15())
    buildings = append(buildings, Age2Building16())
    buildings = append(buildings, Age2Building17())
    buildings = append(buildings, Age2Building18())

    // shuffle buildings
	for i := range buildings {
		j := r.Intn(i + 1)
		buildings[i], buildings[j] = buildings[j], buildings[i]
	}

    return buildings
}