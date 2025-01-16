package engine

import (
	"fmt"
)

type Worker struct {
	Id int

	Color Color

	Available bool
	Wheel_id int // use -1's for off wheel
	Position int 
}

// -- MARK -- Basic methods
func MakeWorker(id int, color Color) *Worker {
	return &Worker {
		Id: id,
		Color: color,
		Available: true,
		Wheel_id: -1,
		Position: -1,
	}
}

func (w *Worker) Clone() *Worker {
	return &Worker {
		Id: w.Id,
		Color: w.Color,
		Available: w.Available,
		Wheel_id: w.Wheel_id,
		Position: w.Position,
	}
}

func (w *Worker) Copy(other *Worker) {
	w.Id = other.Id
	w.Color = other.Color
	w.Available = other.Available
	w.Wheel_id = other.Wheel_id
	w.Position = other.Position
}

func (w *Worker) AddDelta(delta WorkerDelta, mul int) {
	// fmt.Println("Worker AddDelta", w, delta, mul)
	w.Available = Bool(delta.Available, mul, w.Available)
	w.Wheel_id += delta.Wheel_id * mul
	w.Position += delta.Position * mul
	// fmt.Println("Worker Added", w)
}

func (w *Worker) Exact(other *Worker) bool {
	b := w.Id == other.Id &&
		w.Color == other.Color &&
		w.Available == other.Available &&
		w.Wheel_id == other.Wheel_id &&
		w.Position == other.Position

	if !b {
		fmt.Println("Worker", w, other)
	}

	return b
}

// -- MARK -- Unique methods
func (w *Worker) PlaceOn(wheel_id int, corn int) *Delta {
	return &Delta{WorkerDeltas: map[int]WorkerDelta{w.Id: WorkerDelta{
		Available: -1,
		Wheel_id: wheel_id - w.Wheel_id,
		Position: corn - w.Position,
	}}}
}

func (w *Worker) ReturnFrom(wheel *Wheel) *Delta {
	d := &Delta{WorkerDeltas: map[int]WorkerDelta{w.Id: WorkerDelta{
		Available: 1,
		Wheel_id: -1 - w.Wheel_id,
		Position: -1 - w.Position,
	}}}

	d.Add(wheel.RemoveWorker(w.Id))

	return d
}