package model

type Worker struct {
	id int

	color string // todo Color type

	available bool
	wheel_id int// | nil // use -1's
	position int// | nil
}

func (w *Worker) ReturnFrom(wheel *Wheel) {
	w.available = true
	w.wheel_id = -1
	w.position = -1

	
	wheel.RemoveWorker(w.id)
}