package model

import (
	// "tzgolkin/model/util"
)

type Wheel struct {
	// this is all todo
	id int
	size int
	occupied []int
	workers []int


	positions []*Position
	// workers []int
	rotation int
	name string
}

func (w *Wheel) Clone() *Wheel {
	new_occupied := make([]int, 0)
	for _, o := range w.occupied {
		new_occupied = append(new_occupied, o)
	}

	return &Wheel {
		id: w.id,
		size: w.size,
		occupied: new_occupied,
	}
}

func (w *Wheel) AddWorker(position int, worker int) {
	w.occupied = append(w.occupied, position)
	w.workers = append(w.workers, worker)
}

func (w *Wheel) SetRotation(rotation int) {
	w.rotation = rotation
}

func (w *Wheel) RemoveWorker(worker int) {
	// todo use indexof method?
	j := 0
	for i := 0; i < len(w.workers); i++ {
		if w.workers[i] == worker {
			j = i
			break
		}
	}

	w.workers = remove(w.workers, j)
	w.occupied = remove(w.occupied, j)
}

