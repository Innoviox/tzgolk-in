package main

import (
	"fmt"
	"os"
	"tzgolkin/model"
)

func main() {
	game := new(model.Game)
	game.Init()

	p := game.GetPlayer(0)
	for i, move := range game.GenerateMoves(p) {
		fmt.Fprintf(os.Stdout, "%d %s\n", i, move.String())
	}
}