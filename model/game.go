package model

import (
	"tzgolkin/model/wheel"
)

type Game struct {
	players []*Player
	workers []*Worker

	calendar Calendar 
	// temples []Temple 
	// research []Research 
	

	// currentMonuments []Monument
	// allMonuments []Monument 

	// currentBuildings []Building
	// age1Buildings []Building
	// age2Buildings []Building 

	currPlayer int
	firstPlayer int 
}

func (g *Game) Init() {
	g.players := make([]Player, 4)
	for i, color := range [...]string{"red", "blue", "green", "yellow"} {
		g.players[i] = &Player{
			resources: [...]int{0, 0, 0},
			corn: 0,
			color: color
		}

		for j := 0, j < 6; i++ {
			g.workers = append(g.workers, &Worker{
				id: i * 6 + j,
				color: color,
				available: j < 3,
				wheel_id: nil,
				position: nil
			})
		}
	}

	g.calendar = new(wheel.Calendar)
	g.calendar.Init()

	g.currPlayer = 0
	g.firstPlayer = 0
}

func (g *Game) round() {
	g.currPlayer = g.firstPlayer
	for i := 0; i < len(g.players); i++ {
		g.players[g.currPlayer].play()
		g.currPlayer = (g.currPlayer + 1) % len(g.players)
	}

	// todo first player nonsense
	// todo food days
	g.calendar.rotate()
}

func (g *Game) generateMoves(p *Player) []Move {
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

	for _, w := range g.workers {
		if w.player.id == p.id {
			if w.wheel != nil {
				retrieval = append(retrieval, w.id)
			} else if w.available {
				placement = append(placement, w.id)
			}
		}
	}

	// todo find extend method
	for _, m := range make_retrieval_moves([make_empty_retrieval_move()], retrieval) {
		moves = append(moves, m)
	}
	
	for _, m := range make_placement_moves([make_empty_placement_move()], placement) {
		moves = append(moves, m)
	}

	return moves
}

/*
	make_retrieval_moves(moves, retrieval)
		if retrieval = []:
			return moves
		i = retrieval[0]
		r = retrieval[1:]
		m = moves[:]
		for j in moves:
			m.append(j + i)
		return make_retrieval_moves(m, r)
*/
func (g *Game) make_retrieval_moves(moves []Move, retrieval []int) []Move {
	if len(retrieval) == 0 {
		return moves
	}

	worker = retrieval[0]
	rest = retrieval[1:]
	out = copy_moves(moves)

	for _, move := range moves {
		out = append(out, move.retrieve(worker))
	}
	return make_retrieval_moves(out, rest)
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
			ws.execute(j)
			for p in ws.legal_places()
				m.append(j + place(p, i))

		return mpm(m, p, wheels)
*/
func (g *Game) make_placement_moves(moves []Move, placement []int) []Move {
	// todo first player nonsense
	if len(placement) == 0 {
		return moves
	}

	worker = placement[0]
	rest = placement[1:]
	out = copy_moves(moves)

	for _, move := range moves {
		new_wheels = copy_wheels(wheels)
		new_wheels.execute(move)

		for _, position := range new_wheels.legal_positions() {
			out = append(out, move.place(worker, position))
		}
	}

	return make_placement_moves(out, rest)
}