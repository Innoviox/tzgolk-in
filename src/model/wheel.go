package model

type Wheel struct {
	// this is all todo
	Id int
	Size int

	// map from position => worker id
	Occupied map[int]int

	Positions []*Position
	Name string
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
	}
}

func (w *Wheel) Clone() *Wheel {
	new_occupied := make(map[int]int)
	for k, v := range w.Occupied {
		new_occupied[k] = v
	}

	return &Wheel {
		Id: w.Id,
		Size: w.Size,
		Occupied: new_occupied,
		Positions: w.Positions,
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