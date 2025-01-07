package main

import (
	"fmt"
	"os"
	"tzgolkin/model"
	
)

func main() {
	game := new(model.Game)
	game.Init()

	game.Round()
}