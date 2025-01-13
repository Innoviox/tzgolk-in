package buildings

import (
    "fmt"
    "math/rand"
    . "tzgolkin/model"
    . "tzgolkin/impl/wheels"
)

func Age2Building1() Building {
    return Building {
        Id: 1,
        Cost: [4]int{0, 0, 2, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Temples.Step(p, 2, 2)
                    p.Points += 3
                },
                Description: "2 GT, 3 points",
            }}
        },
        Color: Blue,
    }
}

func Age2Building2() Building {
    return Building {
        Id: 2,
        Cost: [4]int{0, 2, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Temples.Step(p, 0, 2)
                    p.Points += 2
                },
                Description: "2 BT, 2 points",
            }}
        },
        Color: Blue,
    }
}

func Age2Building3() Building {
    return Building {
        Id: 3,
        Cost: [4]int{0, 0, 3, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Temples.Step(p, 1, 2)
                },
                Description: "2 YT, 4 points",
            }}
        },
        Color: Blue,
    }
}

func Age2Building4() Building {
    return Building {
        Id: 4,
        Cost: [4]int{0, 1, 2, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return g.Research.GetOptions(g, p, 2, true)
        },
        Color: Blue,
    }
}

func Age2Building5() Building {
    return Building {
        Id: 5,
        Cost: [4]int{0, 2, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Research.FreeResearch(p.Color, Theology)
                    g.Temples.Step(p, 0, 1)
                    g.Temples.Step(p, 2, 1)
                },
                Description: "free theo, 1 BT, 1 GT",
            }}
        },
        Color: Blue,
    }
}

func Age2Building6() Building {
    return Building {
        Id: 6,
        Cost: [4]int{3, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            for _, o := range g.Research.GetOptions(g, p, 1, true) {
                options = append(options, Option{
                    Execute: func(g *Game, p *Player) {
                        o.Execute(g, p)
                        p.Resources[Stone] += 1
                    },
                    Description: fmt.Sprintf("%s, 1 stone", o.Description),
                })
            }

            return options
        },
        Color: Green,
    }
}

func Age2Building7() Building {
    return Building {
        Id: 7,
        Cost: [4]int{1, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.UnlockWorker(p.Color)
                    p.Points += 6
                },
                Description: "unlock worker, 6 points",
            }}
        },
        Color: Red,

    }
}

func Age2Building8() Building {
    return Building {
        Id: 8,
        Cost: [4]int{0, 0, 3, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            for _, o := range Uxmal2(g, p) {
                options = append(options, Option{
                    Execute: func(g *Game, p *Player) {
                        o.Execute(g, p)
                        p.Points += 6
                    },
                    Description: fmt.Sprintf("%s, 6 points", o.Description),
                })
            }

            options = append(options, Option{
                Execute: func(g *Game, p *Player) {
                    p.Points += 6
                },
                Description: "6 points",
            })

            return options
        },
        Color: Red,
    }
}

func Age2Building9() Building {
    return Building {
        Id: 9,
        Cost: [4]int{1, 0, 2, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    p.Points += 8
                },
                Description: "8 points",
            }}
        },
        Color: Red,
    }
}

func Age2Building10() Building {
    return Building {
        Id: 10,
        Cost: [4]int{2, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    p.FreeWorkers += 3
                },
                Description: "3 free workers",
            }}
        },
        Color: Yellow,
    }
}

func Age2Building11() Building {
    return Building {
        Id: 11,
        Cost: [4]int{1, 2, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    for i := 0; i < 3; i++ {
                        g.Temples.Step(p, i, 1)
                    }
                    p.Points += 3
                },
                Description: "1 BT 1 GT 1 YT, 3 points",
            }}
        },
        Color: Red,
    }
}

func Age2Building12() Building {
    return Building {
        Id: 12,
        Cost: [4]int{2, 2, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            for _, o := range g.Research.GetOptions(g, p, 1, true) {
                options = append(options, Option{
                    Execute: func(g *Game, p *Player) {
                        o.Execute(g, p)
                        p.Resources[Skull] += 1
                    },
                    Description: fmt.Sprintf("%s, 1 skull", o.Description),
                })
            }
            return options
        },
        Color: Green,
    }
}

func Age2Building13() Building {
    return Building {
        Id: 13,
        Cost: [4]int{0, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            return []Option {Option{
                Execute: func(g *Game, p *Player) {
                    g.Research.FreeResearch(p.Color, Construction)
                    p.Points += 3
                },
                Description: "free const, 3 points",
            }}
        },
        Color: Blue,
    }
}

func Age2Building14() Building {
    b := Age2Building10()
    return Building {
        Id: 14,
        Cost: b.Cost,
        GetEffects: b.GetEffects,
        Color: Yellow,
    }
}

func Age2Building15() Building {
    b := Age2Building10()
    return Building {
        Id: 15,
        Cost: b.Cost,
        GetEffects: b.GetEffects,
        Color: Yellow,
    }
}

func Age2Building16() Building {
    return Building {
        Id: 16,
        Cost: [4]int{2, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            // for _, o := range Uxmal5(g, p) {
            //     options = append(options, Option{
            //         Execute: func(g *Game, p *Player) {
            //             o.Execute(g, p)
            //             p.Points += 2
            //         },
            //         Description: fmt.Sprintf("%s, 2 points", o.Description),
            //     })
            // }
            return options
        },
        Color: Red,
    }
}

func Age2Building17() Building {
    return Building {
        Id: 17,
        Cost: [4]int{0, 0, 1, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            for _, o := range g.Research.GetOptions(g, p, 1, true) {
                options = append(options, Option{
                    Execute: func(g *Game, p *Player) {
                        o.Execute(g, p)
                        p.Corn += 6
                    },
                    Description: fmt.Sprintf("%s, 6 Corn", o.Description),
                })
            }
            return options
        },
        Color: Green,
    }
}

func Age2Building18() Building {
    return Building {
        Id: 18,
        Cost: [4]int{2, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []Option {
            options := make([]Option, 0)
            for _, o := range g.Research.GetOptions(g, p, 1, true) {
                options = append(options, Option{
                    Execute: func(g *Game, p *Player) {
                        o.Execute(g, p)
                        p.Resources[Gold] += 1
                    },
                    Description: fmt.Sprintf("%s, 1 G", o.Description),
                })
            }
            return options
        },
        Color: Green,
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