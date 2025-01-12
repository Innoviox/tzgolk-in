package model

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

type Option func()
type Options func(*Game, *Player) []Option


type Resource int

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
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	case Yellow:
		return "Yellow"
	}
	return "Unknown"
}

func MakeEmptyRetrievalMove() Move {
	return Move {
		placing: false,
		workers: make([]int, 0),
		positions: make([]*Position, 0),
		corn: 0,
	}
}

func MakeEmptyPlacementMove() Move {
	return Move {
		placing: true,
		workers: make([]int, 0),
		positions: make([]*Position, 0),
		corn: 0,
	}
}

func flatten(options []Options) Options {
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
			for _, r := range PayBlocks(newResources, nBlocks - 1) {
				result = append(result, r)
			}
		}
	}

	return result
}