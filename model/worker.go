package model

type Worker struct {
	Id int

	Color Color

	Available bool
	Wheel_id int// | nil // use -1's
	Position int// | nil
}

func MakeWorker(id int, color Color) *Worker {
	return &Worker {
		id: id,
		color: color,
		available: true,
		wheel_id: -1,
		position: -1,
	}
}

func (w *Worker) ReturnFrom(wheel *Wheel) {
	w.available = true
	w.wheel_id = -1
	w.position = -1

	
	wheel.RemoveWorker(w.id)
}

func (w *Worker) Clone() *Worker {
	return &Worker {
		id: w.id,
		color: w.color,
		available: w.available,
		wheel_id: w.wheel_id,
		position: w.position,
	}
}