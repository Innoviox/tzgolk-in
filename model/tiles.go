package model

import (
	"math/rand"
)

type Tile struct {
	Execute func(*Game, *Player) // todo color type
}

func Tile1(g *Game, p *Player) {
	p.corn += 6
	p.resources[Stone] += 2
}

func Tile2(g *Game, p *Player) {
	p.corn += 3
	p.resources[Wood] += 1
	p.freeWorkers += 1
}

func Tile3(g *Game, p *Player) {
	p.resources[Stone] += 1
	p.resources[Gold] += 1
	g.research.FreeResearch(p.color, Agriculture)
}

func Tile4(g *Game, p *Player) {
	p.corn += 3
	p.resources[Gold] += 1
	g.research.FreeResearch(p.color, Construction)
}

func Tile5(g *Game, p *Player) {
	p.corn += 5
	p.resources[Stone] += 1
	g.research.FreeResearch(p.color, Theology)
}

func Tile6(g *Game, p *Player) {
	p.resources[Wood] += 1
	g.temples.Step(p.color, 2, 1)
	g.research.FreeResearch(p.color, Resources)
}

func Tile7(g *Game, p *Player) {
	p.corn += 2
	p.resources[Wood] += 2
	g.temples.Step(p.color, 2, 1)
}

func Tile8(g *Game, p *Player) {
	p.corn += 5
	p.resources[Stone] += 1
	g.temples.Step(p.color, 0, 1)
}

func Tile9(g *Game, p *Player) {
	p.corn += 7
	p.resources[Wood] += 2
}

func Tile10(g *Game, p *Player) {
	p.corn += 4
	p.resources[Wood] += 1
	g.research.FreeResearch(p.color, Resources)
}

func Tile11(g *Game, p *Player) {
	p.corn += 2
	g.temples.Step(p.color, 0, 1)
	g.research.FreeResearch(p.color, Construction)
}

func Tile12(g *Game, p *Player) {
	p.corn += 4
	p.resources[Wood] += 1
	p.resources[Skull] += 1
}

func Tile13(g *Game, p *Player) {
	p.corn += 9
	p.resources[Stone] += 1
}

func Tile14(g *Game, p *Player) {
	g.UnlockWorker(p.color)
}

func Tile15(g *Game, p *Player) {
	p.corn += 2
	p.resources[Wood] += 2
	g.research.FreeResearch(p.color, Theology)
}

func Tile16(g *Game, p *Player) {
	p.corn += 4
	p.resources[Wood] += 3
}

func Tile17(g *Game, p *Player) {
	p.corn += 5
	p.resources[Gold] += 1
	g.temples.Step(p.color, 1, 1)
}

func Tile18(g *Game, p *Player) {
	p.corn += 3
	g.temples.Step(p.color, 1, 1)
	g.research.FreeResearch(p.color, Agriculture)
}

func Tile19(g *Game, p *Player) {
	p.corn += 3
	p.resources[Wood] += 2
	p.resources[Stone] += 1
}

func Tile20(g *Game, p *Player) {
	p.corn += 8
	p.resources[Gold] += 1
}

func Tile21(g *Game, p *Player) {
	p.corn += 6
	p.resources[Wood] += 1
	p.resources[Stone] += 1
}

func MakeWealthTiles() []Tile {
	tiles := make([]Tile, 0)

	tiles = append(tiles, Tile{ Tile1 })
	tiles = append(tiles, Tile{ Tile2 })
	tiles = append(tiles, Tile{ Tile3 })
	tiles = append(tiles, Tile{ Tile4 })
	tiles = append(tiles, Tile{ Tile5 })
	tiles = append(tiles, Tile{ Tile6 })
	tiles = append(tiles, Tile{ Tile7 })
	tiles = append(tiles, Tile{ Tile8 })
	tiles = append(tiles, Tile{ Tile9 })
	tiles = append(tiles, Tile{ Tile10 })
	tiles = append(tiles, Tile{ Tile11 })
	tiles = append(tiles, Tile{ Tile12 })
	tiles = append(tiles, Tile{ Tile13 })
	tiles = append(tiles, Tile{ Tile14 })
	tiles = append(tiles, Tile{ Tile15 })
	tiles = append(tiles, Tile{ Tile16 })
	tiles = append(tiles, Tile{ Tile17 })
	tiles = append(tiles, Tile{ Tile18 })
	tiles = append(tiles, Tile{ Tile19 })
	tiles = append(tiles, Tile{ Tile20 })
	tiles = append(tiles, Tile{ Tile21 })

	// shuffle tiles
	for i := range tiles {
		j := rand.Intn(i + 1)
		tiles[i], tiles[j] = tiles[j], tiles[i]
	}

	return tiles
}