package engine

import (
	"fmt"
	"math/rand"
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

func InvCost(cost [4]int) [4]int {
	result := [4]int{}
	for i := 0; i < 4; i++ {
		result[i] = -cost[i]
	}
	return result
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	new := make(map[K]V)
	for k, v := range m {
		new[k] = v
	}
	return new
}

func RandZeros(m map[int]int, lo int, hi int, n int) []int {
	result := make([]int, 0)

	for i := lo; i < hi; i++ {
		if m[i] == 0 {
			result = append(result, i)
		}
	}

	if len(result) <= n {
		return result
	}

	rand.Shuffle(len(result), func(i, j int) { 
		result[i], result[j] = result[j], result[i] 
	})
	return result[:n]
}

func CountValues(m map[int]int, v int) int {
	count := 0
	for _, val := range m {
		if val == v {
			count += 1
		}
	}
	return count
}