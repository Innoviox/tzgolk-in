package engine

import (
    "fmt"
    // "os"
)

func (g *Game) GenerateMoves(p *Player, key int) []Move {
	// all possible moves are:
	// - retrieve any combination of workers
	// - place any legal combination of workers
	
	/*
	how to store workers?
	- option 1: game has a list of workers[]
		each worker stores its position and its player, as well as if it's off the board or not
		to retrieve workers, we go

		moves = []
		make a list for retrieval
		make a list for placement
		for each worker
			if that worker is on a wheel
				add its id (?) to the retrieval list
			else
				add its id to the placement list

		moves += call mrm([empty_move], retrieval)

		moves += call mpm([empty_move], placement, wheels[:])
		moves -= empty_move
		return moves
	*/

	moves := make([]Move, 0)

	// worker ids
	retrieval := make([]int, 0)
	placement := make([]int, 0)

	for _, w := range g.Workers {
		if w.Color == p.Color {
			if g.Calendar.WheelFor(w.Id) != nil {
				retrieval = append(retrieval, w.Id)
			} else if w.Unlocked {
				placement = append(placement, w.Id)
			}
		}
	}

	retrieval_moves := g.AddBegging(MakeEmptyRetrievalMove(p.Color), p)
	moves = append(moves, g.MakeRetrievalMoves(retrieval_moves, retrieval, 10000 * key)...)
	
	placement_moves := g.AddBegging(MakeEmptyPlacementMove(p.Color), p)
	moves = append(moves, g.MakePlacementMoves(placement_moves, placement, 20000 * key)...)

	// todo find filter method
	out := make([]Move, 0)
	for _, move := range moves {
		if len(move.Workers) > 0 && move.Corn <= p.Corn {
			out = append(out, move)
		}
	}

	// fmt.Fprintf(os.Stdout, "Generated %d moves for %s\n", len(out), p.Color.String())

	return out
}

func (g *Game) AddBegging(move Move, player *Player) []Move {
	moves := []Move{move}

	if player.Corn >= 3 {
		return moves
	}
	
	for i := 0; i < 3; i++ {
		if g.Temples.CanStep(player, i, -1) {
			moves = append(moves, move.Beg(i))
		}
	}

	return moves
}

func (g *Game) GetOptions(worker *Worker) []*SpecificPosition {
	// fmt.Printf("%v\n", worker)
	positions := make([]*SpecificPosition, 0)
	// if worker.Wheel_id < -1 || worker.Position < 0 || worker.Position >= g.Calendar.Wheels[worker.Wheel_id].Size {
	// 	fmt.Printf("invalid worker %v\n", worker)
	// 	return []*Delta{}
	// }
	wheel := g.Calendar.WheelFor(worker.Id)
	if (wheel == nil ) {
		fmt.Printf("[FATAL ERROR] wheel not found for worker %v\n", worker)
		return positions
	}
	// fmt.Printf("\twheel %s worker %v\n", wheel.Name, worker)
	position := wheel.Positions[wheel.Occupied[worker.Id]]
	player := g.GetPlayerByColor(worker.Color)
	// fmt.Printf("\tpostion %v player %v\n", position, player)
	options := position.GetOptions(g, player)

	// fmt.Fprintf(os.Stdout, "\tOptions for worker on wheel %s position %d: %v\n", wheel.Name, worker.Position, len(options))
	// fmt.Println("aergaergPosition", position)
	// return options
	for _, option := range options {
		positions = append(positions, &SpecificPosition {
			Wheel_id: wheel.Id,
			Corn: position.Corn,
			Execute: option,
		})
	}

	return positions
}

/*
	make_retrieval_moves(moves, retrieval)
		if retrieval = []:
			return moves
		i = retrieval[0]
		r = retrieval[1:]
		m = moves[:]
		for j in moves:
			m.Append(j + i)
		return make_retrieval_moves(m, r)
*/
func (g *Game) MakeRetrievalMoves(moves []Move, retrieval []int, key int) []Move {
	if len(retrieval) == 0 {
		return moves
	}

	out := make([]Move, 0)
	out = append(out, moves...)
	
	for _, w := range retrieval {
		worker := g.GetWorker(w)
		wOrig := worker.Clone()

		m := make([]Move, 0)
		m = append(m, moves...)

		// fmt.Fprintf(os.Stdout, "\t\tR %v W %v\n", retrieval, w)
		// fmt.Fprintf(os.Stdout, "%v %v\n", prev.GetWorker(w), g.GetWorker(w))
		rest := except(retrieval, w)
		// fmt.Fprintf(os.Stdout, "\t\tRest %v\n", rest)

		for i := 0; i < len(moves); i++ {
			// g2 := g.Clone()
			d := g.Calendar.Execute(moves[i], g, func(s string){})
			
			
			g.AddDelta(d, 1)
			for _, position := range g.GetOptions(wOrig) {
				// if worker.Wheel_id != 4 {
				// 	for j := 1; j < worker.Position; j++ {

				// 		m = append(m, moves[i].Retrieve(w, &SpecificPosition {
				// 			Wheel_id: worker.Wheel_id,
				// 			Corn: j,
				// 			Execute: option,
				// 		}, worker.Position - j))
				// 	}
				// }
				m = append(m, moves[i].Retrieve(w, position, 0))
			}
			g.AddDelta(d, -1)
			// if !g.Exact(g2) {
			// 	fmt.Println("PLATO ERROR 0")
			// 	fmt.Println(d)
			// 	fmt.Println(d.CalendarDelta)
			// 	fmt.Println([]int{}[1])
			// }
		}

		out = append(out, g.MakeRetrievalMoves(m, rest, key + 1)...)
	}
	return out
}

/*
	make_placement_moves(moves, placement, wheels)
		if placement = []:
			return moves
		
		i = placement[0]
		p = placement[1:]
		m = moves[:]

		for j in moves:
			ws = wheels[:]
			ws.Execute(j)
			for p in ws.Legal_places()
				m.Append(j + place(p, i))

		return mpm(m, p, wheels)
*/
func (g *Game) MakePlacementMoves(moves []Move, placement []int, key int) []Move {
	// fmt.Fprintf(os.Stdout, "\nMakePlacementMoves %v %v\n", len(moves), placement)
	if len(placement) == 0 {
		return moves
	}

	worker := placement[0]
	rest := placement[1:]

	l := len(moves)
	for i := 0; i < l; i++ {
		d := g.Calendar.Execute(moves[i], g, func(s string){})
		// g2 := g.Clone()
		g.AddDelta(d, 1)

		for _, position := range g.Calendar.LegalPositions() {
			moves = append(moves, moves[i].Place(worker, position))
		}
		g.AddDelta(d, -1)
		// if !g.Exact(g2) {
		// 	fmt.Println("PLATO ERROR 1")
		// 	fmt.Println(d)
		// 	fmt.Println([]int{}[1])
		// }
	}

	return g.MakePlacementMoves(moves, rest, key + 1)
}