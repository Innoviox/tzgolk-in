package types

type Wheel struct {
	// this is all todo
	id int
	size int

	// occupied []int
	// workers []int
	// map from position => worker id
	occupied map[int]int


	positions []*Position
	name string
}