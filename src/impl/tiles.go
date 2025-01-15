package impl

import (
	"math/rand"
	. "tzgolkin/engine"
)

func Tile1(g *Game, p *Player) {
	p.Corn += 6
	p.Resources[Stone] += 2
}

func Tile2(g *Game, p *Player) {
	p.Corn += 3
	p.Resources[Wood] += 1
	p.FreeWorkers += 1
}

func Tile3(g *Game, p *Player) {
	p.Resources[Stone] += 1
	p.Resources[Gold] += 1
	g.Research.FreeResearch(g, p, Agriculture)[0].Execute(g, p)
}

func Tile4(g *Game, p *Player) {
	p.Corn += 3
	p.Resources[Gold] += 1
	g.Research.FreeResearch(g, p, Construction)[0].Execute(g, p)
}

func Tile5(g *Game, p *Player) {
	p.Corn += 5
	p.Resources[Stone] += 1
	g.Research.FreeResearch(g, p, Theology)[0].Execute(g, p)
}

func Tile6(g *Game, p *Player) {
	p.Resources[Wood] += 1
	g.Temples.Step(p, 2, 1)
	g.Research.FreeResearch(g, p, Resources)[0].Execute(g, p)
}

func Tile7(g *Game, p *Player) {
	p.Corn += 2
	p.Resources[Wood] += 2
	g.Temples.Step(p, 2, 1)
}

func Tile8(g *Game, p *Player) {
	p.Corn += 5
	p.Resources[Stone] += 1
	g.Temples.Step(p, 0, 1)
}

func Tile9(g *Game, p *Player) {
	p.Corn += 7
	p.Resources[Wood] += 2
}

func Tile10(g *Game, p *Player) {
	p.Corn += 4
	p.Resources[Wood] += 1
	g.Research.FreeResearch(g, p, Resources)[0].Execute(g, p)
}

func Tile11(g *Game, p *Player) {
	p.Corn += 2
	g.Temples.Step(p, 0, 1)
	g.Research.FreeResearch(g, p, Construction)[0].Execute(g, p)
}

func Tile12(g *Game, p *Player) {
	p.Corn += 4
	p.Resources[Wood] += 1
	p.Resources[Skull] += 1
}

func Tile13(g *Game, p *Player) {
	p.Corn += 9
	p.Resources[Stone] += 1
}

func Tile14(g *Game, p *Player) {
	g.UnlockWorker(p.Color)
}

func Tile15(g *Game, p *Player) {
	p.Corn += 2
	p.Resources[Wood] += 2
	g.Research.FreeResearch(g, p, Theology)[0].Execute(g, p)
}

func Tile16(g *Game, p *Player) {
	p.Corn += 4
	p.Resources[Wood] += 3
}

func Tile17(g *Game, p *Player) {
	p.Corn += 5
	p.Resources[Gold] += 1
	g.Temples.Step(p, 1, 1)
}

func Tile18(g *Game, p *Player) {
	p.Corn += 3
	g.Temples.Step(p, 1, 1)
	g.Research.FreeResearch(g, p, Agriculture)[0].Execute(g, p)
}

func Tile19(g *Game, p *Player) {
	p.Corn += 3
	p.Resources[Wood] += 2
	p.Resources[Stone] += 1
}

func Tile20(g *Game, p *Player) {
	p.Corn += 8
	p.Resources[Gold] += 1
}

func Tile21(g *Game, p *Player) {
	p.Corn += 6
	p.Resources[Wood] += 1
	p.Resources[Stone] += 1
}

func MakeWealthTiles(r *rand.Rand) []Tile {
	tiles := make([]Tile, 0)

	tiles = append(tiles, Tile{ 1, Tile1 })
	tiles = append(tiles, Tile{ 2, Tile2 })
	tiles = append(tiles, Tile{ 3, Tile3 })
	tiles = append(tiles, Tile{ 4, Tile4 })
	tiles = append(tiles, Tile{ 5, Tile5 })
	tiles = append(tiles, Tile{ 6, Tile6 })
	tiles = append(tiles, Tile{ 7, Tile7 })
	tiles = append(tiles, Tile{ 8, Tile8 })
	tiles = append(tiles, Tile{ 9, Tile9 })
	tiles = append(tiles, Tile{ 10, Tile10 })
	tiles = append(tiles, Tile{ 11, Tile11 })
	tiles = append(tiles, Tile{ 12, Tile12 })
	tiles = append(tiles, Tile{ 13, Tile13 })
	tiles = append(tiles, Tile{ 14, Tile14 })
	tiles = append(tiles, Tile{ 15, Tile15 })
	tiles = append(tiles, Tile{ 16, Tile16 })
	tiles = append(tiles, Tile{ 17, Tile17 })
	tiles = append(tiles, Tile{ 18, Tile18 })
	tiles = append(tiles, Tile{ 19, Tile19 })
	tiles = append(tiles, Tile{ 20, Tile20 })
	tiles = append(tiles, Tile{ 21, Tile21 })

	// shuffle tiles
	for i := range tiles {
		j := r.Intn(i + 1)
		tiles[i], tiles[j] = tiles[j], tiles[i]
	}

	return tiles
}