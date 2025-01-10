package model

type Tile struct {
	Execute func(*Game, int) // todo color type
}

func Tile1(g *Game, playerId int) {
	player := g.GetPlayer(playerId)

	player.corn += 6
	player.resources[Stone] += 2
}

func MakeWealthTiles() []Tile {
	tiles := make([]Tile, 0)

	tiles = append(tiles, Tile{ Tile1 })

	return tiles
}