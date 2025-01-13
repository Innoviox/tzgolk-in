package model

type Monument struct {
    Id int
    Cost [4]int
    GetPoints func (g *Game, p *Player) int
    Color Color
}

// -- MARK -- Basic methods
func MakeMonument(id int, cost [4]int, getPoints func (g *Game, p *Player) int, color Color) *Monument {
    return &Monument {
        Id: id,
        Cost: cost,
        GetPoints: getPoints,
        Color: color,
    }
}
