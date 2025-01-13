package main

import (
	. "tzgolkin/controller"
	. "tzgolkin/disp"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(1))

	ctrl := MakeController(r)
	disp := MakeDisplay(ctrl)

	disp.RunGame()
}