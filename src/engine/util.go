package engine

import (
	"fmt"
	// "os"
)

const TempleDebug = "BYG"

type Resource int
const ResourceDebug = "WSGC"
const (
	Wood Resource = iota
	Stone
	Gold
	Skull
)

// todo get actual names
const (
	Agriculture Science = iota
	Resources
	Construction
	Theology
)

const ResearchDebug = "ARCT"


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

type Tile struct {
	N int
	Execute func(*Game, *Player) // todo color type
}

// https://stackoverflow.Com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove[T any](slice []T, s int) []T {
	// fmt.Fprintf(os.Stderr, "removing %d from %v\n", s, slice)
    return append(slice[:s], slice[s+1:]...)
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

func CostString(cost [4]int) string {
	result := ""
	for i := 0; i < 4; i++ {
		if cost[i] > 0 {
			result += fmt.Sprintf("%d%s ", cost[i], string(ResourceDebug[i]))
		}
	}
	return result
}
