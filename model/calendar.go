package model

import (
	"fmt"
	"os"
	"strings"
	"github.com/innoviox/tzgolkin/model/types"
)

func (c *Calendar) Init() {
	// c.wheels = make([]*Wheel, 0)
	c.rotation = 0

	c.wheels = []*Wheel {
		MakePalenque(),
		MakeYaxchilan(),
		MakeTikal(),
		MakeUxmal(),
		MakeChichen(),
	}

	c.firstPlayer = -1
	c.clone = false

	// for i := 0; i < 5; i++ {
	// 	c.AddWheel(&Wheel {
	// 		id: i,
	// 		size: 5,
	// 		occupied: make([]int, 0),
	// 	})
	// }
}

func (c *Calendar) Clone() *Calendar {
	new_wheels := make([]*Wheel, 0) // todo map method?
	for _, wheel := range c.wheels {
		new_wheels = append(new_wheels, wheel.Clone())
	}

	return &Calendar {
		wheels: new_wheels,
		rotation: c.rotation,
		firstPlayer: c.firstPlayer,
		clone: true,
	}
}

func (c *Calendar) Execute(move Move, game *Game) {
	if !c.clone { fmt.Fprintf(os.Stdout, "Executing move %s\n", move.String()) }
	if (move.placing) {
		if !c.clone { fmt.Fprintf(os.Stdout, "Placing workers %s\n", move.String()) }
		for i := 0; i < len(move.workers); i++ {
			p := move.positions[i]
			worker := game.GetWorker(move.workers[i])

			if p.firstPlayer {
				if !c.clone { fmt.Fprintf(os.Stdout, "First playering!\n") }
				c.firstPlayer = move.workers[i]
				worker.available = false
			} else {
				if !c.clone { fmt.Fprintf(os.Stdout, "Placing worker %d on %s position %d\n", move.workers[i], c.wheels[p.wheel_id].name, p.corn) }
				c.wheels[p.wheel_id].AddWorker(p.corn, move.workers[i])

				worker.available = false
				worker.wheel_id = p.wheel_id
				worker.position = p.corn
			}
		}
	} else {
		for i := 0; i < len(move.workers); i++ {
			// steps:
			// - call the position's function on the game & player id
			// - return the worker to the player
			w := game.GetWorker(move.workers[i])
			p := move.positions[i]

			player := game.GetPlayerByColor(w.color)

			if !c.clone { fmt.Fprintf(os.Stdout, "Retrieving worker %d from %s position %d, executing %s\n", 
						w.id, c.wheels[p.wheel_id].name, p.corn, p.Execute.description) }
			p.Execute.Execute(game, player)
			w.ReturnFrom(c.wheels[p.wheel_id])
		}
	}
}

func (c *Calendar) LegalPositions() []*SpecificPosition {
	positions := make([]*SpecificPosition, 0)

	for _, wheel := range c.wheels {
		i := wheel.LowestUnoccupied()
		
		positions = append(positions, &SpecificPosition{
			wheel_id: wheel.id,
			corn: i,
		})
	}

	if c.firstPlayer == -1 {
		positions = append(positions, &SpecificPosition{
			wheel_id: -2,
			corn: 0,
			firstPlayer: true,
		})
	}

	return positions
}

// func (c *Calendar) Rotate(g *Game) {
// 	c.rotation = rotation
// 	for i := 0; i < len(c.wheels); i++ {
// 		c.wheels[i].SetRotation(rotation)
// 	}
// }

// todo rework when days are implemented?

func (c *Calendar) Rotate(g *Game) {
	// c.SetRotation(c.rotation + 1);
	for i := 0; i < len(c.wheels); i++ {
		c.wheels[i].Rotate(g)
	}
}

func (c *Calendar) String(workers []*Worker) string {
	var br strings.Builder

	fmt.Fprintf(&br, "----Calendar------------\n")
	for _, wheel := range c.wheels {
		fmt.Fprintf(&br, "| %s: ", wheel.name)

		out := make([]string, wheel.size)

		for k, v := range wheel.occupied {
			out[k] = workers[v].color.String()
		}

		for _, o := range out {
			if len(o) > 0 {
				fmt.Fprintf(&br, "%s", o)
			} else {
				fmt.Fprintf(&br, "_")
			}
		}
		fmt.Fprintf(&br, "(%v)\n", wheel.occupied)
	}
	fmt.Fprintf(&br, "------------------------\n")

	return br.String()
}