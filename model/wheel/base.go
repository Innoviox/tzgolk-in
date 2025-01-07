package wheel

import (
	"tzgolkin/model/util"
)

type Wheel struct {
	// this is all todo
	id int
	size int
	occupied []int


	// positions []Position
	// workers []int
	// rotation int
}

// func (w *Wheel) AddPosition(p Position) {
// 	w.positions = append(w.positions, p)
// }

// func (w *Wheel) AddWorker(worker int) {
// 	w.workers = append(w.workers, worker)
// }

// func (w *Wheel) rotate() {
// 	i := 0
// 	for i < len(w.workers) {
// 		w.workers[i] = w.workers[i] + 1;
// 		if w.workers[i] >= len(w.positions) {
// 			remove(w.workers, i)
// 			// todo: return worker to player
// 		} else {
// 		i = i + 1
// 		}
// 	}
// }
