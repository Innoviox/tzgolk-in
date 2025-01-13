package model

import (
	"fmt"
	// "os"
)

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove[T any](slice []T, s int) []T {
	// fmt.Fprintf(os.Stderr, "removing %d from %v\n", s, slice)
    return append(slice[:s], slice[s+1:]...)
}

type Option struct {
	Execute func(*Game, *Player)
	Description string
	BuildingNum int
}

type Options func(*Game, *Player) []Option


type Resource int

const ResourceDebug = "WSGC"
const TempleDebug = "BYG"

const (
	Wood Resource = iota
	Stone
	Gold
	Skull
)

type Color int

const (
	Red Color = iota
	Green
	Blue
	Yellow
)

func (c Color) String() string {
	switch c {
	case Red:
		return "R"
	case Green:
		return "G"
	case Blue:
		return "B"
	case Yellow:
		return "Y"
	}
	return "Unknown"
}

func MakeEmptyRetrievalMove() Move {
	return Move {
		placing: false,
		workers: make([]int, 0),
		positions: make([]*SpecificPosition, 0),
		corn: 0,
	}
}

func MakeEmptyPlacementMove() Move {
	return Move {
		placing: true,
		workers: make([]int, 0),
		positions: make([]*SpecificPosition, 0),
		corn: 0,
	}
}

func flatten(options []Options) Options {
	// todo add "mirror" to description?
	return func (g *Game, p *Player) []Option {
		result := make([]Option, 0)
		for _, o := range options {
			result = append(result, o(g, p)...)
		}
		return result
	}
}

func PayBlocks(resources [4]int, nBlocks int) [][4]int {
	if nBlocks == 0 {
		return [][4]int{resources}
	}

	result := make([][4]int, 0)
	for i := 0; i < 3; i++ {
		if resources[i] > 0 {
			newResources := [4]int{}
			copy(newResources[:], resources[:])
			newResources[i] -= 1
			
			result = append(result, PayBlocks(newResources, nBlocks - 1)...)
		}
	}

	return result
}

func except(arr []int, n int) []int {
	new := make([]int, 0)
	for _, v := range arr {
		if v != n {
			new = append(new, v)
		}
	}
	return new
}

func CostString(cost [4]int) string {
	result := ""
	for i := 0; i < 4; i++ {
		if cost[i] > 0 {
			result += fmt.Sprintf("%d%s ", cost[i], string(ResourceDebug[i]))
		}
	}
	return result
}

func TotalCorn(p *Player) int {
	corn := p.corn
	corn += 2 * p.resources[Wood]
	corn += 3 * p.resources[Stone]
	corn += 4 * p.resources[Gold]
	
	return corn
}

type Tile struct {
	n int
	Execute func(*Game, *Player) // todo color type
}
