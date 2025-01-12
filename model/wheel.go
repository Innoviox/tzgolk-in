package model

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
	new_occupied = append(new_occupied, w.occupied...)

	new_workers := make([]int, 0)
	new_workers = append(new_workers, w.workers...)

	return &Wheel {
		id: w.id,
		size: w.size,
		occupied: new_occupied,
		workers: new_workers,
		positions: w.positions,
		rotation: w.rotation,
		name: w.name,
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
		occupied: make([]int, 0),
		workers: make([]int, 0),
		positions: positions, 
		rotation: 0,
		name: wheel_name,
	}
}