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
	
	currentBuildings []Building
	age1Buildings []Building
	age2Buildings []Building 

	currPlayer int
	firstPlayer int 

	accumulatedCorn int

	age int
	day int
	resDays []int
	pointDays []int
}

func (g *Game) Init() {
	g.players = make([]*Player, 4)
	for i, color := range [...]Color{Red, Green, Blue, Yellow} {
		g.players[i] = &Player{
			resources: [...]int{0, 0, 0, 0},
			corn: 0,
			color: color,
			points: 0,
			freeWorkers: 0,
			workerDeduction: 0,
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
	g.research = MakeResearch()

	g.age1Buildings = MakeAge1Buildings()
	g.age2Buildings = MakeAge2Buildings()

	// todo currentbuildings
	// todo monuments


	g.TileSetup()

	g.currPlayer = 0
	g.firstPlayer = 0
	g.accumulatedCorn = 0

	g.age = 1
	g.day = 0
	g.resDays = []int{7, 20}
	g.pointDays = []int{13, 26}
}

func (g *Game) TileSetup() {
	// todo: 4 choose 2

	tiles := MakeWealthTiles()
	t := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Fprintf(os.Stdout, "Placing tile %d for player %s\n", tiles[t].n, g.GetPlayer(i).color.String())
			tiles[t].Execute(g, g.GetPlayer(i))
			t++
		}
	}
}

func (g *Game) Round() {
	g.currPlayer = g.firstPlayer
	for i := 0; i < len(g.players); i++ {
		g.TakeTurn()
		g.currPlayer = (g.currPlayer + 1) % len(g.players)
	}

	if g.calendar.firstPlayer != -1 {
		g.FirstPlayer()
	}

	// todo food days
	fmt.Fprintf(os.Stdout, "Rotating Calendar\n")
	g.Rotate()

	fmt.Fprintf(os.Stdout, "Calendar State: \n%s\n", g.calendar.String(g.workers))

	for i := 0; i < len(g.players); i++ {
		fmt.Fprintf(os.Stdout, "%s", g.players[i].String())
	}
	fmt.Fprintf(os.Stdout, "Accumulated corn: %d\n", g.accumulatedCorn)
}

func (g *Game) FirstPlayer() {
	worker := g.GetWorker(g.calendar.firstPlayer)
	worker.available = true
	g.calendar.firstPlayer = -1
	player := g.GetPlayerByColor(worker.color)
	player.corn += g.accumulatedCorn
	g.accumulatedCorn = 0

	playerIdx := 0
	for i := 0; i < len(g.players); i++ {
		if g.players[i].color == player.color {
			playerIdx = i
			break
		}
	}

	if g.firstPlayer == playerIdx {
		g.firstPlayer = (g.firstPlayer + 1) % len(g.players)
	} else {
		g.firstPlayer = playerIdx
	}

	nextDayIsFoodDay := false
	for _, day := range g.resDays {
		if g.day == day - 1 {
			nextDayIsFoodDay = true
			break
		}
	}
	for _, day := range g.pointDays {
		if g.day == day - 1 {
			nextDayIsFoodDay = true
			break
		}
	}

	if !nextDayIsFoodDay && player.lightSide && rand.Intn(2) == 0 {
		// todo actually decide
		player.lightSide = false
		fmt.Fprintf(os.Stdout, "Player %s has gone dark\n", player.color)
		g.Rotate()
	}
}

func (g *Game) Rotate() {
	g.calendar.Rotate(g)
	g.accumulatedCorn += 1
	g.day += 1
	g.CheckDay()
}

func (g *Game) CheckDay() {
	for _, day := range g.resDays {
		if g.day == day {
			g.FoodDay()

			for _, player := range g.players {
				g.temples.GainResources(player)
			}
			return
		}
	}

	for _, day := range g.pointDays {
		if g.day == day {
			g.FoodDay()
			
			for _, player := range g.players {
				g.temples.GainPoints(player, g.age)
			}

			g.age += 1
			if g.age == 2 {
				// todo deal age 2 buildings
			} else {
				// todo end game
			}
			return
		}
	}
}

func (g *Game) FoodDay() {
	for _, player := range g.players {
		paid := 0
		unpaid := 0
		for _, w := range g.workers {
			if w.color == player.color {
				if w.wheel_id != -1 || w.available {
					if player.corn >= 2 {
						player.corn -= 2
						paid += 1
					} else {
						player.points -= 3
						unpaid += 1
					}
				}
			}
		}
		fmt.Fprintf(os.Stdout, "Player %s paid %d workers, didn't pay %d workers\n", player.color.String(), paid, unpaid)
	}
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

	// fmt.Fprintf(os.Stdout, "\t%s R %v P %v\n", p.color, retrieval, placement)

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

func (g *Game) GetOptions(worker *Worker) []Option {
	wheel := g.calendar.wheels[worker.wheel_id]
	position := wheel.positions[worker.position]
	player := g.GetPlayerByColor(worker.color)

	options := position.GetOptions(g, player)

	// fmt.Fprintf(os.Stdout, "\tOptions for worker on wheel %s position %d: %v\n", wheel.name, worker.position, len(options))

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
			m.append(j + i)
		return make_retrieval_moves(m, r)
*/
func (g *Game) MakeRetrievalMoves(moves []Move, retrieval []int) []Move {
	// fmt.Fprintf(os.Stdout, "\tMakeRetrievalMoves %v %v\n", len(moves), retrieval)

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
			for _, option := range g.GetOptions(worker) {
				// fmt.Fprintf(os.Stdout, "\t\tOption %v\n", option)
				// todo pay corn to go lower
				m = append(m, moves[i].Retrieve(w, &SpecificPosition {
					wheel_id: worker.wheel_id,
					corn: worker.position,
					Execute: option,
				}))
			}
		}

		out = append(out, g.MakeRetrievalMoves(m, rest)...)
	}
	return out // g.MakeRetrievalMoves(out, rest)
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