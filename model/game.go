package model

import (
	"fmt"
	"os"
	"math/rand"
)

type Game struct {
	// todo multiskull edge case
	players []*Player
	workers []*Worker

	calendar *Calendar 
	temples *Temples
	research *Research 
	

	// currentMonuments []Monument
	// allMonuments []Monument 

	age int
	currentBuildings []*Building
	age1Buildings []*Building
	age2Buildings []*Building 

	currPlayer int
	firstPlayer int 
}

func (g *Game) Init() {
	g.players = make([]*Player, 4)
	for i, color := range [...]Color{Red, Green, Blue, Yellow} {
		g.players[i] = &Player{
			resources: [...]int{0, 0, 0, 0},
			corn: 10, // todo wealth tiles
			color: color,
			points: 0,
		}

		for j := 0; j < 6; j++ {
			g.workers = append(g.workers, &Worker{
				id: i * 6 + j,
				color: color,
				available: j < 3,
				wheel_id: -1,
				position: -1,
			})
		}
	}

	g.calendar = new(Calendar)
	g.calendar.Init()

	g.temples = MakeTemples()

	g.currPlayer = 0
	g.firstPlayer = 0
}

func (g *Game) Round() {
	g.currPlayer = g.firstPlayer
	for i := 0; i < len(g.players); i++ {
		// g.players[g.currPlayer].play()
		g.TakeTurn()
		g.currPlayer = (g.currPlayer + 1) % len(g.players)
	}

	// todo first player nonsense
	// todo food days
	fmt.Fprintf(os.Stdout, "Rotating Calendar\n")
	g.calendar.Rotate()

	fmt.Fprintf(os.Stdout, "Calendar State: \n%s\n", g.calendar.String(g.workers))
}

func (g *Game) TakeTurn() {
	player := g.players[g.currPlayer]
	moves := g.GenerateMoves(g.players[g.currPlayer])
	move := moves[rand.Intn(len(moves))]

	fmt.Fprintf(os.Stdout, "Playing move %s for %s\n", move.String(), player.color)
	
	g.calendar.Execute(move, g)
}

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

	for _, w := range g.workers {
		if w.color == p.color {
			if w.wheel_id != -1 {
				retrieval = append(retrieval, w.id)
			} else if w.available {
				placement = append(placement, w.id)
			}
		}
	}

	retrieval_moves := append(make([]Move, 0), MakeEmptyRetrievalMove())
	moves = append(moves, g.MakeRetrievalMoves(retrieval_moves, retrieval)...)
	
	placement_moves := append(make([]Move, 0), MakeEmptyPlacementMove())
	moves = append(moves, g.MakePlacementMoves(placement_moves, placement)...)

	// todo find filter method
	out := make([]Move, 0)
	for _, move := range moves {
		if len(move.workers) > 0 {
			out = append(out, move)
		}
	}

	// todo filter by corn cost

	return out
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
func (g *Game) MakeRetrievalMoves(moves []Move, retrieval []int) []Move {
	if len(retrieval) == 0 {
		return moves
	}

	worker := retrieval[0]
	rest := retrieval[1:]
	out := make([]Move, 0)

	l := len(moves)

	for i := 0; i < l; i++ {
		// todo: generate choices for retrieval
		// todo: order matters
		moves = append(moves, moves[i].Retrieve(worker))
	}
	return g.MakeRetrievalMoves(out, rest)
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
func (g *Game) MakePlacementMoves(moves []Move, placement []int) []Move {
	// todo first player nonsense
	if len(placement) == 0 {
		return moves
	}

	worker := placement[0]
	rest := placement[1:]

	l := len(moves)
	for i := 0; i < l; i++ {
		new_calendar := g.calendar.Clone()
		new_calendar.Execute(moves[i], g)

		for _, position := range new_calendar.LegalPositions() {
			moves = append(moves, moves[i].Place(worker, position))
		}
	}

	return g.MakePlacementMoves(moves, rest)
}

func (g *Game) GetPlayer(num int) *Player {
	return g.players[num]
}

func (g *Game) GetPlayerByColor(color Color) *Player {
	for _, player := range g.players {
		if player.color == color {
			return player
		}
	}

	return nil
}

func (g *Game) GetWorker(num int) *Worker {
	return g.workers[num]
}

func (g *Game) UnlockWorker(color Color) {
	for _, w := range g.workers {
		if w.color == color {
			if !w.available && w.wheel_id == -1 {
				w.available = true
				break
			}
		}
	}
}