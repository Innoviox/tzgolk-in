package main

import (
	"tzgolkin/model"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(1))

	game := new(model.Game)
	game.Init(r)

	for !game.Over() {
		game.Round()
	}
}