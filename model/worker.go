package model

import (
	"github.com/innoviox/tzgolkin/model/types"
)

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