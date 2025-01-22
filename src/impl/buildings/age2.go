package buildings

import (
    // "fmt"
    // "math/rand"
    . "tzgolkin/engine"
    . "tzgolkin/impl/wheels"
)

func Age2Building1() Building {
    return Building {
        Id: 15,
        Cost: [4]int{0, 0, 2, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            return []*Delta {Combine(
                g.Temples.Step(p, 2, 2),
                PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Points: 3,
                }),
            )}
        },
        Color: Blue,
    }
}

func Age2Building2() Building {
    return Building {
        Id: 16,
        Cost: [4]int{0, 2, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            return []*Delta {Combine(
                g.Temples.Step(p, 0, 2),
                PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Points: 2,
                }),
            )}
        },
        Color: Blue,
    }
}

func Age2Building3() Building {
    return Building {
        Id: 17,
        Cost: [4]int{0, 0, 3, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            return []*Delta {Combine(
                g.Temples.Step(p, 1, 2),
                PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Points: 4,
                }),
            )}
        },
        Color: Blue,
    }
}

func Age2Building4() Building {
    return Building {
        Id: 18,
        Cost: [4]int{0, 1, 2, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            return g.Research.GetOptions(g, p, 2, true)
        },
        Color: Blue,
    }
}

func Age2Building5() Building {
    return Building {
        Id: 19,
        Cost: [4]int{0, 2, 1, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
            for _, o := range g.Research.FreeResearch(g, p, Theology) {
                d := g.Temples.Step(p, 0, 1)
                d.Add(g.Temples.Step(p, 2, 1), true)
                d.Add(o, true)
                options = append(options, d)
            }
            return options
        },
        Color: Blue,
    }
}

func Age2Building6() Building {
    return Building {
        Id: 20,
        Cost: [4]int{3, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
            for _, o := range g.Research.GetOptions(g, p, 1, true) {
                d := PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Resources: [4]int{0, 1, 0, 0},
                })

                d.Add(o, true)
                options = append(options, d)
            }

            return options
        },
        Color: Green,
    }
}

func Age2Building7() Building {
    return Building {
        Id: 21,
        Cost: [4]int{1, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            return []*Delta {Combine(
                g.UnlockWorker(p.Color),
                PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Points: 6,
                }),
            )}
        },
        Color: Red,

    }
}

func Age2Building8() Building {
    return Building {
        Id: 22,
        Cost: [4]int{0, 0, 3, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
            for _, o := range Uxmal2(g, p) {
                options = append(options, Combine(o, PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Points: 6,
                })))
            }

            options = append(options, PlayerDeltaWrapper(p.Color, PlayerDelta{
                Points: 6,
            }))

            return options
        },
        Color: Red,
    }
}

func Age2Building9() Building {
    return Building {
        Id: 23,
        Cost: [4]int{1, 0, 2, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            return []*Delta{PlayerDeltaWrapper(p.Color, PlayerDelta{
                Points: 8,
            })}
        },
        Color: Red,
    }
}

func Age2Building10() Building {
    return Building {
        Id: 24,
        Cost: [4]int{2, 0, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            return []*Delta{PlayerDeltaWrapper(p.Color, PlayerDelta{
                FreeWorkers: 3,
            })}
        },
        Color: Yellow,
    }
}

func Age2Building11() Building {
    return Building {
        Id: 25,
        Cost: [4]int{1, 2, 1, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            d := PlayerDeltaWrapper(p.Color, PlayerDelta{
                Points: 3,
            })
            for i := 0; i < 3; i++ {
                d.Add(g.Temples.Step(p, i, 1), true)
            }
            return []*Delta{d}
        },
        Color: Red,
    }
}

func Age2Building12() Building {
    return Building {
        Id: 26,
        Cost: [4]int{2, 2, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
            for _, o := range g.Research.GetOptions(g, p, 1, true) {
                d := PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Resources: [4]int{0, 0, 0, 1},
                })
                d.Add(o, true)
                options = append(options, d)
            }
            return options
        },
        Color: Green,
    }
}

func Age2Building13() Building {
    return Building {
        Id: 27,
        Cost: [4]int{0, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
            for _, o := range g.Research.FreeResearch(g, p, Construction) {
                d := PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Points: 3,
                })
                d.Add(o, true)
                options = append(options, d)
            }
            return options
        },
        Color: Blue,
    }
}

func Age2Building14() Building {
    b := Age2Building10()
    return Building {
        Id: 28,
        Cost: b.Cost,
        GetEffects: b.GetEffects,
        Color: Yellow,
    }
}

func Age2Building15() Building {
    b := Age2Building10()
    return Building {
        Id: 29,
        Cost: b.Cost,
        GetEffects: b.GetEffects,
        Color: Yellow,
    }
}

func Age2Building16() Building {
    return Building {
        Id: 30,
        Cost: [4]int{2, 1, 1, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
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
        Id: 31,
        Cost: [4]int{3, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
            for _, o := range g.Research.GetOptions(g, p, 1, true) {
                d := PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Corn: 6,
                })
                d.Add(o, true)
                options = append(options, d)
            }
            return options
        },
        Color: Green,
    }
}

func Age2Building18() Building {
    return Building {
        Id: 32,
        Cost: [4]int{2, 1, 0, 0},
        GetEffects: func (g *Game, p *Player) []*Delta {
            options := make([]*Delta, 0)
            for _, o := range g.Research.GetOptions(g, p, 1, true) {
                d := PlayerDeltaWrapper(p.Color, PlayerDelta{
                    Resources: [4]int{0, 0, 1, 0},
                })
                d.Add(o, true)
                options = append(options, d)
            }
            return options
        },
        Color: Green,
    }
}

func MakeAge2Buildings() map[int]Building {
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

    b := map[int]Building{}
    for _, building := range buildings {
        b[building.Id] = building
    }
    return b
}