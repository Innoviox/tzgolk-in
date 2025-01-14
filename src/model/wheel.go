package model

import (
	"fmt"
	"strings"
)

// '⁰'

type Wheel struct {
	// this is all todo
	Id int
	Size int

	// map from position => worker id
	Occupied map[int]int

	Positions []*Position
	Name string
	String func(*Wheel, []*Worker) string
}

// -- MARK -- Basic methods
func MakeWheel(options []Options, Wheel_id int, wheel_name string) *Wheel {
	positions := make([]*Position, 0)

	for i := 0; i < len(options); i++ {
		positions = append(positions, &Position{
			Wheel_id: Wheel_id,
			Corn: i,
			GetOptions: options[i],
		})
	}

	for i := 6; i < 8; i++ {
		positions = append(positions, &Position{
			Wheel_id: Wheel_id,
			Corn: i,
			GetOptions: Flatten(options),
		})
	}

	return &Wheel{
		Id: Wheel_id,
		Size: len(positions),
		Occupied: make(map[int]int),
		Positions: positions, 
		Name: wheel_name,
		String: func (wheel *Wheel, workers []*Worker) string {
			var br strings.Builder
		
			fmt.Fprintf(&br, "| %-12s: ", wheel.Name)
		
			out := make([]string, wheel.Size)
		
			for k, v := range wheel.Occupied {
				out[k] = workers[v].Color.String()
			}
		
			for k, o := range out {
				if len(o) > 0 {
					fmt.Fprintf(&br, "  %s", o)
				} else {
					fmt.Fprintf(&br, "%3d", k)
				}
			}
			fmt.Fprintf(&br, "\n")
		
			return br.String()
		},
	}
}

func (w *Wheel) Clone() *Wheel {
	new_occupied := make(map[int]int)
	for k, v := range w.Occupied {
		new_occupied[k] = v
	}

	new_positions := make([]*Position, 0)
	for _, p := range w.Positions {
		new_positions = append(new_positions, p.Clone())
	}

	return &Wheel {
		Id: w.Id,
		Size: w.Size,
		Occupied: new_occupied,
		Positions: new_positions,
		Name: w.Name,
	}
}

// -- MARK -- Unique methods
func (w *Wheel) AddWorker(position int, worker int) {
	w.Occupied[position] = worker
}

func (w *Wheel) RemoveWorker(worker int) {
	for k, v := range w.Occupied {
		if v == worker {
			delete(w.Occupied, k)
			return
		}
	}
}

func (w *Wheel) Rotate(g *Game) {
	workerToRemove := -1 // only one worker per wheel can fall off
	new_occupied := make(map[int]int)
	for k, v := range w.Occupied {
		if k >= w.Size - 1 {
			workerToRemove = v
		} else {
			new_occupied[k + 1] = v
			worker := g.GetWorker(v)
			worker.Position++
		}
	} 

	if workerToRemove != -1 {
		g.GetWorker(workerToRemove).ReturnFrom(w)
	}

	w.Occupied = new_occupied
}

func (w *Wheel) LowestUnoccupied() int {
	for i := 0; i < w.Size; i++ {
		_, ok := w.Occupied[i]
		if !ok {
			return i
		}
	}

	return -1
}