package types

type Monument struct {
    id int
    cost [4]int
    GetPoints func (g *Game, p *Player) int
    color Color
}