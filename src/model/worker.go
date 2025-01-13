package model

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

// -- MARK -- Unique methods
func (w *Worker) ReturnFrom(wheel *Wheel) {
	w.Available = true
	w.Wheel_id = -1
	w.Position = -1

	wheel.RemoveWorker(w.Id)
}