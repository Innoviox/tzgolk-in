package main

import (
	. "tzgolkin/controller"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(1))

	ctrl := MakeController(r)

	ctrl.RunGame()
}