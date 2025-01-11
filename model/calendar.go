package model

import (
	"fmt"
	"strings"
)

type Calendar struct {
	wheels []*Wheel
	rotation int
	// tojm9 do food days & such
}

func (c *Calendar) Init() {
	c.wheels = make([]*Wheel, 0)
	c.rotation = 0

	for i := 0; i < 5; i++ {
		c.AddWheel(&Wheel {
			id: i,
			size: 5,
			occupied: make([]int, 0),
		})
	}
}

func (c *Calendar) AddWheel(w *Wheel) {
	c.wheels = append(c.wheels, w)
}

func (c *Calendar) Clone() Calendar {
	new_wheels := make([]*Wheel, 0) // todo map method?
	for _, wheel := range c.wheels {
		new_wheels = append(new_wheels, wheel.Clone())
	}

	return Calendar {
		wheels: new_wheels,
		rotation: c.rotation,
	}
}

func (c *Calendar) Execute(move Move, game *Game) {
	if (move.placing) {
		for i := 0; i < len(move.workers); i++ {
			p := move.positions[i]
			c.wheels[p.wheel_id].AddWorker(p.corn, move.workers[i])
		}
	} else {
		for i := 0; i < len(move.workers); i++ {
			// steps:
			// - call the position's function on the game & player id
			// - return the worker to the player
			w := game.GetWorker(move.workers[i])
			p := move.positions[i]

			p.Execute(game, w.color)
			w.ReturnFrom(c.wheels[p.wheel_id])
		}
	}
}

func (c *Calendar) LegalPositions() []*Position {
	positions := make([]*Position, 0)

	for _, wheel := range c.wheels {
		i := 0
		for j := 0; j < len(wheel.occupied); j++ {
			if (wheel.occupied[j] + wheel.rotation) > i {
				break
			} else {
				i++
			}
		}
		
		positions = append(positions, &Position{
			wheel_id: wheel.id,
			position_num: i,
			corn: i,
		})
	}

	return positions
}

func (c *Calendar) SetRotation(rotation int) {
	c.rotation = rotation
	for i := 0; i < len(c.wheels); i++ {
		c.wheels[i].SetRotation(rotation)
	}
}

func (c *Calendar) Rotate() {
	c.SetRotation(c.rotation + 1);
}

func (c *Calendar) String(workers []*Worker) string {
	var br strings.Builder

	for i, wheel := range c.wheels {
		fmt.Fprintf(&br, "%d: ", i)

		out := make([]string, wheel.size)

		for i := 0; i < len(wheel.occupied); i++ {
			if wheel.occupied[i] + wheel.rotation < wheel.size {
				out[wheel.occupied[i] + wheel.rotation] = workers[wheel.workers[i]].color
			}
		}

		for _, o := range out {
			if len(o) > 0 {
				fmt.Fprintf(&br, "%s", o)
			} else {
				fmt.Fprintf(&br, "_")
			}
		}
		fmt.Fprintf(&br, "\n")
	}


	return br.String()
}