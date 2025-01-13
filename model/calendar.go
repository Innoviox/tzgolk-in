package model

import (
	"fmt"
	"os"
	"strings"
)

type Calendar struct {
	Wheels []*Wheel
	Rotation int
	FirstPlayer int
	IsClone bool
	// tojm9 do food days & such
}

// func (c *Calendar) Init() {
func MakeCalendar(wheels []*Wheel) *Calendar {
	// c.Wheels = make([]*Wheel, 0)
	return &Calendar {
		Rotation: 0,
		Wheels: wheels,
		FirstPlayer: -1,
		IsClone: false,
	}
	// c.Rotation = 0

	// // c.Wheels = []*Wheel {
	// // 	MakePalenque(),
	// // 	MakeYaxchilan(),
	// // 	MakeTikal(),
	// // 	MakeUxmal(),
	// // 	MakeChichen(),
	// // }
	// c.Wheels = wheels

	// c.FirstPlayer = -1
	// c.Clone = false

	// for i := 0; i < 5; i++ {
	// 	c.AddWheel(&Wheel {
	// 		Id: i,
	// 		Size: 5,
	// 		Occupied: make([]int, 0),
	// 	})
	// }
}

func (c *Calendar) Clone() *Calendar {
	new_wheels := make([]*Wheel, 0) // todo map method?
	for _, wheel := range c.Wheels {
		new_wheels = append(new_wheels, wheel.Clone())
	}

	return &Calendar {
		Wheels: new_wheels,
		Rotation: c.Rotation,
		FirstPlayer: c.FirstPlayer,
		IsClone: true,
	}
}

func (c *Calendar) Execute(move Move, game *Game) {
	if !c.IsClone { fmt.Fprintf(os.Stdout, "Executing move %s\n", move.String()) }
	if (move.Placing) {
		if !c.IsClone { fmt.Fprintf(os.Stdout, "Placing workers %s\n", move.String()) }
		for i := 0; i < len(move.Workers); i++ {
			p := move.Positions[i]
			worker := game.GetWorker(move.Workers[i])

			if p.FirstPlayer {
				if !c.IsClone { fmt.Fprintf(os.Stdout, "First playering!\n") }
				c.FirstPlayer = move.Workers[i]
				worker.Available = false
			} else {
				if !c.IsClone { fmt.Fprintf(os.Stdout, "Placing worker %d on %s position %d\n", move.Workers[i], c.Wheels[p.Wheel_id].Name, p.Corn) }
				c.Wheels[p.Wheel_id].AddWorker(p.Corn, move.Workers[i])

				worker.Available = false
				worker.Wheel_id = p.Wheel_id
				worker.Position = p.Corn
			}
		}
	} else {
		for i := 0; i < len(move.Workers); i++ {
			// steps:
			// - call the position's function on the game & player id
			// - return the worker to the player
			w := game.GetWorker(move.Workers[i])
			p := move.Positions[i]

			player := game.GetPlayerByColor(w.Color)

			if !c.IsClone { fmt.Fprintf(os.Stdout, "Retrieving worker %d from %s position %d, executing %s\n", 
						w.Id, c.Wheels[p.Wheel_id].Name, p.Corn, p.Execute.Description) }
			p.Execute.Execute(game, player)
			w.ReturnFrom(c.Wheels[p.Wheel_id])
		}
	}
}

func (c *Calendar) LegalPositions() []*SpecificPosition {
	positions := make([]*SpecificPosition, 0)

	for _, wheel := range c.Wheels {
		i := wheel.LowestUnoccupied()
		
		positions = append(positions, &SpecificPosition{
			Wheel_id: wheel.Id,
			Corn: i,
		})
	}

	if c.FirstPlayer == -1 {
		positions = append(positions, &SpecificPosition{
			Wheel_id: -2,
			Corn: 0,
			FirstPlayer: true,
		})
	}

	return positions
}

// func (c *Calendar) Rotate(g *Game) {
// 	c.Rotation = rotation
// 	for i := 0; i < len(c.Wheels); i++ {
// 		c.Wheels[i].SetRotation(rotation)
// 	}
// }

// todo rework when days are implemented?

func (c *Calendar) Rotate(g *Game) {
	// c.SetRotation(c.Rotation + 1);
	for i := 0; i < len(c.Wheels); i++ {
		c.Wheels[i].Rotate(g)
	}
}

func (c *Calendar) String(workers []*Worker) string {
	var br strings.Builder

	fmt.Fprintf(&br, "----Calendar------------\n")
	for _, wheel := range c.Wheels {
		fmt.Fprintf(&br, "| %s: ", wheel.Name)

		out := make([]string, wheel.Size)

		for k, v := range wheel.Occupied {
			out[k] = workers[v].Color.String()
		}

		for _, o := range out {
			if len(o) > 0 {
				fmt.Fprintf(&br, "%s", o)
			} else {
				fmt.Fprintf(&br, "_")
			}
		}
		fmt.Fprintf(&br, "(%v)\n", wheel.Occupied)
	}
	fmt.Fprintf(&br, "------------------------\n")

	return br.String()
}