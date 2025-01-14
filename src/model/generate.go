package model

import (
    "fmt"
    "os"
)

func (g *Game) GenerateMoves(p *Player) []Move {
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
			if w.Wheel_id != -1 {
				retrieval = append(retrieval, w.Id)
			} else if w.Available {
				placement = append(placement, w.Id)
			}
		}
	}

	fmt.Fprintf(os.Stdout, "\t%s R %v P %v\n", p.Color, retrieval, placement)

	retrieval_moves := append(make([]Move, 0), MakeEmptyRetrievalMove())
	moves = append(moves, g.MakeRetrievalMoves(retrieval_moves, retrieval)...)
	
	placement_moves := append(make([]Move, 0), MakeEmptyPlacementMove())
	moves = append(moves, g.MakePlacementMoves(placement_moves, placement)...)

	// todo find filter method
	out := make([]Move, 0)
	for _, move := range moves {
		if len(move.Workers) > 0 && move.Corn <= p.Corn {
			out = append(out, move)
		}
	}

	// todo filter by Corn cost

	return out
}

func (g *Game) GetOptions(worker *Worker) []Option {
	wheel := g.Calendar.Wheels[worker.Wheel_id]
	position := wheel.Positions[worker.Position]
	player := g.GetPlayerByColor(worker.Color)

	options := position.GetOptions(g, player)

	// fmt.Fprintf(os.Stdout, "\tOptions for worker on wheel %s position %d: %v\n", wheel.Name, worker.Position, len(options))

	return options
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
func (g *Game) MakeRetrievalMoves(moves []Move, retrieval []int) []Move {
	if len(retrieval) == 0 {
		return moves
	}

	out := make([]Move, 0)
	out = append(out, moves...)
	
	for _, w := range retrieval {
		worker := g.GetWorker(w)

		m := make([]Move, 0)
		m = append(m, moves...)

		// fmt.Fprintf(os.Stdout, "\t\tR %v W %v\n", retrieval, w)
		rest := except(retrieval, w)
		// fmt.Fprintf(os.Stdout, "\t\tRest %v\n", rest)

		for i := 0; i < len(moves); i++ {
			new_game := g.Clone()
			new_game.Calendar.Execute(moves[i], new_game, func(s string){})
			for _, option := range new_game.GetOptions(worker) {
				// if worker.Wheel_id != 4 {
				// 	for j := 1; j < worker.Position; j++ {

				// 		m = append(m, moves[i].Retrieve(w, &SpecificPosition {
				// 			Wheel_id: worker.Wheel_id,
				// 			Corn: j,
				// 			Execute: option,
				// 		}, worker.Position - j))
				// 	}
				// }
				m = append(m, moves[i].Retrieve(w, &SpecificPosition {
					Wheel_id: worker.Wheel_id,
					Corn: worker.Position,
					Execute: option,
				}, 0))
			}
		}

		out = append(out, g.MakeRetrievalMoves(m, rest)...)
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
func (g *Game) MakePlacementMoves(moves []Move, placement []int) []Move {
	// fmt.Fprintf(os.Stdout, "\nMakePlacementMoves %v %v\n", len(moves), placement)
	if len(placement) == 0 {
		return moves
	}

	worker := placement[0]
	rest := placement[1:]

	l := len(moves)
	for i := 0; i < l; i++ {
		new_game := g.Clone()
		new_game.Calendar.Execute(moves[i], new_game, func(s string){})

		for _, position := range new_game.Calendar.LegalPositions() {
			moves = append(moves, moves[i].Place(worker, position))
		}
	}

	return g.MakePlacementMoves(moves, rest)
}