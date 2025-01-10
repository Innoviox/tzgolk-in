package model

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

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