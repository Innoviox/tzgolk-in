package engine

import (
	"fmt"
	// "os"
	"strings"
)

type Calendar struct {
	Wheels []*Wheel
	Rotation int
	FirstPlayer int
	IsClone bool
}

// -- MARK -- Basic methods
func MakeCalendar(wheels []*Wheel) *Calendar {
	return &Calendar {
		Rotation: 0,
		Wheels: wheels,
		FirstPlayer: -1,
		IsClone: false,
	}
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

func (c *Calendar) Copy(other *Calendar) {
	c.Rotation = other.Rotation
	c.FirstPlayer = other.FirstPlayer
	c.IsClone = other.IsClone

	for i, wheel := range other.Wheels {
		c.Wheels[i].Copy(wheel)
	}
}

func (c *Calendar) AddDelta(delta CalendarDelta, mul int) {
	c.Rotation += delta.Rotation * mul
	c.FirstPlayer += delta.FirstPlayer * mul

	for i, wheel := range delta.WheelDeltas {
		c.Wheels[i].AddDelta(wheel, mul)
	}
}

func (c *Calendar) String(workers []*Worker) string {
	var br strings.Builder

	fmt.Fprintf(&br, "----Calendar------------\n")
	for _, wheel := range c.Wheels {
		fmt.Fprintf(&br, "%s", wheel.String(wheel, workers))
	}
	fmt.Fprintf(&br, "| First Player Spot (0): ")
	if c.FirstPlayer != -1 {
		fmt.Fprintf(&br, "%s\n", workers[c.FirstPlayer].Color.String())
	} else {
		fmt.Fprintf(&br, "None\n")
	}
	fmt.Fprintf(&br, "------------------------\n")

	return br.String()
}


// -- MARK -- Unique methods
/* 
func (c *Calendar) Execute(move Move, game *Game, MarkStep func(string)) {
	player := game.GetPlayerByColor(move.Player)

	if (move.Begged != -1) {
		player.Corn = 3
		game.Temples.Step(player, move.Begged, -1)
		if !c.IsClone { MarkStep(fmt.Sprintf("Begging for corn on temple %s", string(TempleDebug[move.Begged]))) }
	}

	player.Corn -= move.Corn
	if !c.IsClone { MarkStep(fmt.Sprintf("Executing move %s for %s (-%d corn)", move.String(), player.Color.String(), move.Corn)) }
	
	if (move.Placing) {
		// if !c.IsClone { fmt.Fprintf(os.Stdout, "Placing workers %s\n", move.String()) }
		for i := 0; i < len(move.Workers); i++ {
			p := move.Positions[i]
			worker := game.GetWorker(move.Workers[i])

			if p.FirstPlayer {
				// if !c.IsClone { fmt.Fprintf(os.Stdout, "First playering!\n") }
				if !c.IsClone { MarkStep(fmt.Sprintf("First playering %s", worker.Color.String())) }
				c.FirstPlayer = move.Workers[i]
				worker.Available = false
			} else {
				// if !c.IsClone { fmt.Fprintf(os.Stdout, "Placing worker %d on %s position %d\n", move.Workers[i], c.Wheels[p.Wheel_id].Name, p.Corn) }
				if !c.IsClone { MarkStep(fmt.Sprintf("Placing worker %s on %s position %d", worker.Color.String(), c.Wheels[p.Wheel_id].Name, p.Corn)) }
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

			// if !c.IsClone { fmt.Fprintf(os.Stdout, "Retrieving worker %d from %s position %d, executing %s\n", 
						// w.Id, c.Wheels[p.Wheel_id].Name, p.Corn, p.Execute.Description) }
			if !c.IsClone { MarkStep(fmt.Sprintf("Retrieving worker %s from %s position %d, executing %s",
				w.Color.String(), c.Wheels[p.Wheel_id].Name, p.Corn, p.Execute.Description)) }
			p.Execute.Execute(game, player)
			w.ReturnFrom(c.Wheels[p.Wheel_id])
		}
	}
}
*/

func (c *Calendar) Execute(move Move, game *Game, MarkStep func(string)) *Delta {
	player := game.GetPlayerByColor(move.Player)

	d := &Delta{}

	pd := PlayerDelta{}
	if move.Begged != -1 {
		pd.Corn = 3 - player.Corn
		// todo use actual Step for lightsiding
		d.Add(game.Temples.Step(player, move.Begged, -1))
	}
	pd.Corn -= move.Corn
	d.Add(&Delta{PlayerDeltas: map[Color]PlayerDelta{move.Player: pd}})

	for i := 0; i < len(move.Workers); i++ {
		p := move.Positions[i]
		w := game.GetWorker(move.Workers[i])
		if (move.Placing) {
			if p.FirstPlayer {
				d.Add(&Delta{
					WorkerDeltas: map[int]WorkerDelta{move.Workers[i]: WorkerDelta{
						Available: -1,
					}},
					CalendarDelta: CalendarDelta{
						FirstPlayer: move.Workers[i] - c.FirstPlayer,
					},
				})
			} else {
				d.Add(c.Wheels[p.Wheel_id].AddWorker(p.Corn, move.Workers[i]))
				d.Add(w.PlaceOn(p.Wheel_id, p.Corn))
			}
		} else {
			d.Add(p.Execute)
			d.Add(w.ReturnFrom(c.Wheels[p.Wheel_id]))
		}
	}
	
	return d
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

func (c *Calendar) Rotate(g *Game) *Delta {
	d := &Delta{}
	for i := 0; i < len(c.Wheels); i++ {
		d.Add(c.Wheels[i].Rotate(g))
	}
	return d
}