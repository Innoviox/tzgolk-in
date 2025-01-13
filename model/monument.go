package model

import (
    "math/rand"
)

type Monument struct {
    id int
    cost [4]int
    GetPoints func (g *Game, p *Player) int
    color Color
}

func Monument1() Monument {
    return Monument {
        id: 1,
        cost: [4]int{0, 3, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            highestLocation := 2
            for i := 0; i < 3; i++ {
                step := g.temples.temples[i].playerLocations[p.color]
                if step > highestLocation {
                    highestLocation = step
                }
            }
            return highestLocation * 3
        },
        color: Blue,
    }
}

func Monument2() Monument {
    return Monument {
        id: 2,
        cost: [4]int{2, 3, 1, 0},
        GetPoints: func (g *Game, p *Player) int {
            points := 0
            for _, b := range p.buildings {
                if b.color == Green {
                    points += 4
                }
            }

            for _, b := range p.monuments {
                if b.color == Green {
                    points += 4
                }
            }

            return points 
        },
        color: Green,
    }
}

func Monument3() Monument {
    return Monument {
        id: 3,
        cost: [4]int{0, 4, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            points := 0
            for i := 0; i < 3; i++ {
                step := g.temples.temples[i].playerLocations[p.color]
                points += g.temples.temples[i].points[step]
            }
            return points
        },
        color: Blue,
    }
}

func Monument4() Monument {
    return Monument {
        id: 4,
        cost: [4]int{3, 2, 1, 0},
        GetPoints: func (g *Game, p *Player) int {
            points := 4
            for _, b := range p.buildings {
                if b.color == Red {
                    points += 4
                }
            }

            return points 
        },
        color: Red,
    }
}

func Monument5() Monument {
    return Monument {
        id: 5,
        cost: [4]int{0, 2, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            points := 0
            for _, b := range p.buildings {
                if b.color == Blue {
                    points += 4
                }
            }

            for _, b := range p.monuments {
                if b.color == Blue {
                    points += 4
                }
            }

            return points 
        },
        color: Blue,
    }
}

func Monument6() Monument {
    return Monument {
        id: 6,
        cost: [4]int{3, 0, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            workers := 0
            for _, w := range g.workers {
                if w.color == p.color && w.available {
                    workers += 1
                }
            }

            return []int{0, 0, 0, 6, 12, 18}[workers]
        },
        color: Green,
    }
}

func Monument7() Monument {
    return Monument {
        id: 7,
        cost: [4]int{1, 1, 4, 0},
        GetPoints: func (g *Game, p *Player) int {
            return 4 * p.cornTiles
        },
        color: Red,
    }
}

func Monument8() Monument {
    return Monument {
        id: 8,
        cost: [4]int{1, 0, 4, 0},
        GetPoints: func (g *Game, p *Player) int {
            return 4 * p.woodTiles
        },
        color: Red,
    }
}

func Monument9() Monument {
    return Monument {
        id: 9,
        cost: [4]int{1, 3, 2, 0},
        GetPoints: func (g *Game, p *Player) int {
            return 2 * len(p.buildings) + 2 * len(p.monuments)
        },
        color: Red,
    }
}

func Monument10() Monument {
    return Monument {
        id: 10,
        cost: [4]int{2, 2, 2, 0},
        GetPoints: func (g *Game, p *Player) int {
            n := 0
            for _, p := range g.players {
                n += len(p.monuments)
            }
            return []int{0, 6, 5, 4}[len(g.players)] * n
        },
        color: Red,
    }
}

func Monument11() Monument {
    return Monument {
        id: 11,
        cost: [4]int{1, 1, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            n := 0
            for _, l := range g.research.levels[p.color] {
                if l == 3 {
                    n += 1
                }
            }

            return []int{0, 9, 20, 33, 33}[n]
        },
        color: Red,
    }
}

func Monument12() Monument {
    return Monument {
        id: 12,
        cost: [4]int{2, 1, 3, 0},
        GetPoints: func (g *Game, p *Player) int {
            n := 0
            for _, l := range g.research.levels[p.color] {
                n += l
            }

            return n * 3
        },
        color: Red,
    }
}

func Monument13() Monument {
    return Monument {
        id: 13,
        cost: [4]int{0, 0, 4, 1},
        GetPoints: func (g *Game, p *Player) int {
            n := 0
            
            for _, p := range g.calendar.wheels[4].positions {
                if p.cData.full {
                    n += 1
                }
            }

            return n * 3
        },
        color: Red,
    }
}

func MakeMonuments() []Monument {
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

    for i := range monuments {
		j := rand.Intn(i + 1)
		monuments[i], monuments[j] = monuments[j], monuments[i]
	}

    return monuments
}

func (m *Monument) CanBuild(player *Player) bool {
    for i := 0; i < 4; i++ {
        if player.resources[i] < m.cost[i] {
            return false
        }
    }

    return true
}