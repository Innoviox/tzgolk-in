package main

import (
	. "tzgolkin/controller"
	. "tzgolkin/disp"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(78))

	ctrl := MakeController(r)
	disp := MakeDisplay(ctrl)

	disp.Run()
}