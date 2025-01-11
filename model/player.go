package model

type Player struct {
	resources [4]int
	corn int 
	color Color

	points int
	cornTiles int
	woodTiles int
}