package main

import (
	"tzgolkin/model"
)

func main() {
	game := new(model.Game)
	game.Init()

	game.Round()
}