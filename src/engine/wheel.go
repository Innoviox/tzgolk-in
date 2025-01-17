package engine

import (
	"fmt"
	"strings"
	"reflect"
)

// '⁰'

type Wheel struct {
	// this is all todo
	Id int
	Size int

	// map from position to worker
	Occupied map[int]int

	Positions []*Position
	Name string
	String func(*Wheel, []*Worker) string
}

// -- MARK -- Basic methods
func MakeWheel(options []Options, Wheel_id int, wheel_name string) *Wheel {
	positions := make([]*Position, 0)
	occupied := make(map[int]int)

	for i := 0; i < len(options); i++ {
		positions = append(positions, &Position{
			Wheel_id: Wheel_id,
			Corn: i,
			GetOptions: options[i],
		})
		occupied[i] = -1
	}

	for i := 6; i < 8; i++ {
		positions = append(positions, &Position{
			Wheel_id: Wheel_id,
			Corn: i,
			GetOptions: Flatten(options),
		})
		occupied[i] = -1
	}

	return &Wheel{
		Id: Wheel_id,
		Size: len(positions),
		Occupied: occupied,
		Positions: positions, 
		Name: wheel_name,
		String: func (wheel *Wheel, workers []*Worker) string {
			var br strings.Builder
		
			fmt.Fprintf(&br, "| %-12s: ", wheel.Name)
		
			out := make([]string, wheel.Size)
		
			for k, v := range wheel.Occupied {
				if v >= 0 {
					out[k] = workers[v].Color.String()
				}
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

func (w *Wheel) Copy(other *Wheel) {
	w.Id = other.Id
	w.Size = other.Size
	w.Name = other.Name

	w.Occupied = make(map[int]int)
	for k, v := range other.Occupied {
		w.Occupied[k] = v
	}

	for i, p := range other.Positions {
		w.Positions[i].Copy(p)
	}
}

func (w *Wheel) Exact(other *Wheel) bool {
	if w.Id != other.Id || w.Size != other.Size || w.Name != other.Name {
		fmt.Println("gerijo")
		return false
	}

	if !reflect.DeepEqual(w.Occupied, other.Occupied) {
		fmt.Println("\nocc", w.Occupied, other.Occupied)
		return false
	}

	for i, p := range w.Positions {
		if !p.Exact(other.Positions[i]) {
			fmt.Println("pos", i)
			return false
		}
	}

	return true
}

func (w *Wheel) AddDelta(delta WheelDelta, mul int) {
	// todo how should this work?
	if delta.Sign != 0 {
		if mul * delta.Sign < 0 {
			w.Occupied = CopyMap(delta.OldOccupied)
		} else {
			w.Occupied = CopyMap(delta.NewOccupied)
		}
	}

	for i, p := range delta.PositionDeltas {
		w.Positions[i].AddDelta(p, mul)
	}
}

// -- MARK -- Unique methods
func (w *Wheel) AddWorker(position int, worker int) *Delta {
	newOccupied := CopyMap(w.Occupied)
	newOccupied[position] = worker

	return w.MakeDelta(newOccupied, 1)
}

// func (w *Wheel) RemoveWorker(worker int) *Delta {
// 	newOccupied := CopyMap(w.Occupied)
	
// 	for k, v := range w.Occupied {
// 		if v == worker {
// 			delete(newOccupied, k)
// 			return w.MakeDelta(newOccupied, 1)
// 		}
// 	}

// 	return w.MakeDelta(newOccupied, 1)
// }

func (w *Wheel) MakeDelta(Occupied map[int]int, Sign int) *Delta {
	return &Delta{CalendarDelta: CalendarDelta{WheelDeltas: map[int]WheelDelta{w.Id: WheelDelta{
		OldOccupied: CopyMap(w.Occupied),
		NewOccupied: Occupied,
		Sign: Sign,
	}}}}
}

func (w *Wheel) Rotate(g *Game) *Delta {
	d := &Delta{}

	// workerToRemove := -1 // only one worker per wheel can fall off
	new_occupied := make(map[int]int)
	for k, v := range w.Occupied {
		if v >= w.Size - 1 || v == -1 {
			// workerToRemove = k
		} else {
			new_occupied[k] = -1
			new_occupied[k + 1] = v
			// worker := g.GetWorker(v)
			// d.Add(&Delta{WorkerDeltas: map[int]WorkerDelta{worker.Id: WorkerDelta{
			// 	// Position: 1,
			// }}})
		}
	} 

	// if workerToRemove != -1 {
	// 	d.Add(g.GetWorker(workerToRemove).ReturnFrom(w))
	// }

	d.Add(w.MakeDelta(new_occupied, 1))

	return d
}

func (w *Wheel) LowestUnoccupied() int {
	for i := 0; i < w.Size; i++ {
		v := w.Occupied[i]
		if v == -1 {
			return i
		}
	}

	return -1
}