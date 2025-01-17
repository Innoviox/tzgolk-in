package buildings

import (
    . "tzgolkin/engine"
)

func Monument1() Monument {
    return Monument {
        Id: 1,
        Cost: [4]int{0, 3, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            highestLocation := 2
            for i := 0; i < 3; i++ {
                step := g.Temples.Temples[i].PlayerLocations[p.Color]
                if step > highestLocation {
                    highestLocation = step
                }
            }
            return highestLocation * 3
        },
        Color: Blue,
    }
}

func Monument2() Monument {
    return Monument {
        Id: 2,
        Cost: [4]int{2, 3, 1, 0},
        GetPoints: func (g *Game, p *Player) int {
            points := 0
            for k, v := range p.Buildings {
                if v == 1 && g.Buildings[k].Color == Green {
                    points += 4
                }
            }

            for k, v := range p.Monuments {
                if v == 1 && g.Monuments[k].Color == Green {
                    points += 4
                }
            }

            return points 
        },
        Color: Green,
    }
}

func Monument3() Monument {
    return Monument {
        Id: 3,
        Cost: [4]int{0, 4, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            points := 0
            for i := 0; i < 3; i++ {
                step := g.Temples.Temples[i].PlayerLocations[p.Color]
                points += g.Temples.Temples[i].Points[step]
            }
            return points
        },
        Color: Blue,
    }
}

func Monument4() Monument {
    return Monument {
        Id: 4,
        Cost: [4]int{3, 2, 1, 0},
        GetPoints: func (g *Game, p *Player) int {
            points := 0
            for k, v := range p.Buildings {
                if v == 1 && g.Buildings[k].Color == Red {
                    points += 4
                }
            }

            for k, v := range p.Monuments {
                if v == 1 && g.Monuments[k].Color == Red {
                    points += 4
                }
            }

            return points 
        },
        Color: Red,
    }
}

func Monument5() Monument {
    return Monument {
        Id: 5,
        Cost: [4]int{0, 2, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            points := 0
            for k, v := range p.Buildings {
                if v == 1 && g.Buildings[k].Color == Blue {
                    points += 4
                }
            }

            for k, v := range p.Monuments {
                if v == 1 && g.Monuments[k].Color == Blue {
                    points += 4
                }
            }

            return points 
        },
        Color: Blue,
    }
}

func Monument6() Monument {
    return Monument {
        Id: 6,
        Cost: [4]int{3, 0, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            workers := 0
            for _, w := range g.Workers {
                if w.Color == p.Color && w.Unlocked {
                    workers += 1
                }
            }

            return []int{0, 0, 0, 6, 12, 18}[workers]
        },
        Color: Green,
    }
}

func Monument7() Monument {
    return Monument {
        Id: 7,
        Cost: [4]int{1, 1, 4, 0},
        GetPoints: func (g *Game, p *Player) int {
            return 4 * p.CornTiles
        },
        Color: Red,
    }
}

func Monument8() Monument {
    return Monument {
        Id: 8,
        Cost: [4]int{1, 0, 4, 0},
        GetPoints: func (g *Game, p *Player) int {
            return 4 * p.WoodTiles
        },
        Color: Red,
    }
}

func Monument9() Monument {
    return Monument {
        Id: 9,
        Cost: [4]int{1, 3, 2, 0},
        GetPoints: func (g *Game, p *Player) int {
            return 2 * len(p.Buildings) + 2 * len(p.Monuments)
        },
        Color: Red,
    }
}

func Monument10() Monument {
    return Monument {
        Id: 10,
        Cost: [4]int{2, 2, 2, 0},
        GetPoints: func (g *Game, p *Player) int {
            n := 0
            for _, p := range g.Players {
                n += len(p.Monuments)
            }
            return []int{0, 6, 5, 4}[len(g.Players)] * n
        },
        Color: Red,
    }
}

func Monument11() Monument {
    return Monument {
        Id: 11,
        Cost: [4]int{1, 1, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            n := 0
            for _, l := range g.Research.Levels[p.Color] {
                if l == 3 {
                    n += 1
                }
            }

            return []int{0, 9, 20, 33, 33}[n]
        },
        Color: Red,
    }
}

func Monument12() Monument {
    return Monument {
        Id: 12,
        Cost: [4]int{2, 1, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            n := 0
            for _, l := range g.Research.Levels[p.Color] {
                n += l
            }

            return n * 3
        },
        Color: Red,
    }
}

func Monument13() Monument {
    return Monument {
        Id: 13,
        Cost: [4]int{0, 0, 4, 1},
        GetPoints: func (g *Game, p *Player) int {
            n := 0
            
            for _, p := range g.Calendar.Wheels[4].Positions {
                if p.CData.Full {
                    n += 1
                }
            }

            return n * 3
        },
        Color: Red,
    }
}

func MakeMonuments() map[int]Monument {
    monuments := make([]Monument, 0)

    monuments = append(monuments, Monument1())
    monuments = append(monuments, Monument2())
    monuments = append(monuments, Monument3())
    monuments = append(monuments, Monument4())
    monuments = append(monuments, Monument5())
    monuments = append(monuments, Monument6())
    monuments = append(monuments, Monument7())
    monuments = append(monuments, Monument8())
    monuments = append(monuments, Monument9())
    monuments = append(monuments, Monument10())
    monuments = append(monuments, Monument11())
    monuments = append(monuments, Monument12())
    monuments = append(monuments, Monument13())

    m := map[int]Monument{}
    for _, monument := range monuments {
        m[monument.Id] = monument
    }
    return m
}