package model

import (
	// "fmt"
	// "os"
	// "sort"
)

type Wheel struct {
	// this is all todo
	Id int
	Size int

	// occupied []int
	// workers []int
	// map from position => worker id
	Occupied map[int]int


	Positions []*Position
	Name string
}

func (w *Wheel) Clone() *Wheel {
	// new_occupied := make([]int, 0)
	// new_occupied = append(new_occupied, w.occupied...)

	// new_workers := make([]int, 0)
	// new_workers = append(new_workers, w.workers...)
	new_occupied := make(map[int]int)
	for k, v := range w.occupied {
		new_occupied[k] = v
	}

	return &Wheel {
		id: w.id,
		size: w.size,
		occupied: new_occupied,
		positions: w.positions,
		name: w.name,
	}
}

func (w *Wheel) AddWorker(position int, worker int) {
	// fmt.Fprintf(os.Stdout, "Adding worker %d to %s position %d\n", worker, w.name, position)
	// w.occupied = append(w.occupied, position)
	// w.workers = append(w.workers, worker)
	w.occupied[position] = worker
}

func (w *Wheel) Rotate(g *Game) {
	workerToRemove := -1
	new_occupied := make(map[int]int)
	for k, v := range w.occupied {
		if k >= w.size - 1 {
			workerToRemove = v
		} else {
			new_occupied[k + 1] = v
			worker := g.GetWorker(v)
			worker.position++
		}
	} 

	if workerToRemove != -1 {
		g.GetWorker(workerToRemove).ReturnFrom(w)
	}

	w.occupied = new_occupied
}

func (w *Wheel) RemoveWorker(worker int) {
	for k, v := range w.occupied {
		if v == worker {
			delete(w.occupied, k)
			return
		}
	}
}

func MakeWheel(options []Options, wheel_id int, wheel_name string) *Wheel {
	positions := make([]*Position, 0)

	for i := 0; i < len(options); i++ {
		positions = append(positions, &Position{
			wheel_id: wheel_id,
			corn: i,
			GetOptions: options[i],
		})
	}

	for i := 6; i < 8; i++ {
		positions = append(positions, &Position{
			wheel_id: wheel_id,
			corn: i,
			GetOptions: flatten(options),
		})
	}

	return &Wheel{
		id: wheel_id,
		size: len(positions),
		occupied: make(map[int]int),
		positions: positions, 
		name: wheel_name,
	}
}

func (w *Wheel) LowestUnoccupied() int {
	for i := 0; i < w.size; i++ {
		_, ok := w.occupied[i]
		if !ok {
			return i
		}
	}

	return -1
}