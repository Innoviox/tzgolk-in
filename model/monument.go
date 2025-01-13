package model

type Monument struct {
    Id int
    Cost [4]int
    GetPoints func (g *Game, p *Player) int
    Color Color
}

func MakeMonument(id int, cost [4]int, getPoints func (g *Game, p *Player) int, color Color) *Monument {
    return &Monument {
        id: id,
        cost: cost,
        GetPoints: getPoints,
        color: color,
    }
}

func (m *Monument) CanBuild(player *Player) bool {
    for i := 0; i < 4; i++ {
        if player.resources[i] < m.cost[i] {
            return false
        }
    }

    return true
}