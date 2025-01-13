package main

import (
	"tzgolkin/model"
)

func main() {
	game := new(model.Game)
	game.Init()

	for !game.Over() {
		game.Round()
	}
}