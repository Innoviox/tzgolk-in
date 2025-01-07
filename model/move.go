package model 

type Move struct {
	placing bool
	workers []int

	positions []*Position

	corn int
}