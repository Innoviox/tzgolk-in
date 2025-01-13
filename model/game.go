package model

import (
	"fmt"
	"os"
	"strings"
	"math/rand"
	"github.com/innoviox/tzgolkin/model/types"
)

func (g *Game) Init(rand *rand.Rand) {
	g.rand = rand
	g.players = make([]*Player, 4)
	for i, color := range [...]Color{Red, Green, Blue, Yellow} {
		g.players[i] = &Player{
			resources: [...]int{0, 0, 0, 0},
			corn: 0,
			color: color,
			points: 0,
			freeWorkers: 0,
			workerDeduction: 0,
			lightSide: true,
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

	g.age1Buildings = MakeAge1Buildings(g.rand)
	g.age2Buildings = MakeAge2Buildings(g.rand)

	g.nBuildings = 6
	g.currentBuildings = make([]Building, 0)
	g.DealBuildings()

	g.nMonuments = 6
	g.allMonuments = MakeMonuments(g.rand)
	g.currentMonuments = make([]Monument, 0)
	g.DealMonuments()

	g.TileSetup()

	g.currPlayer = 0
	g.firstPlayer = 0
	g.accumulatedCorn = 0

	g.age = 1
	g.day = 0
	g.resDays = []int{7, 20}
	g.pointDays = []int{13, 26}

	g.over = false

	fmt.Fprintf(os.Stdout, "%s", g.String())
}

func (g *Game) TileSetup() {
	// todo: 4 choose 2

	tiles := MakeWealthTiles(g.rand)
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
	if g.over {
		return 
	}

	g.currPlayer = g.firstPlayer
	for i := 0; i < len(g.players); i++ {
		g.TakeTurn()
		g.currPlayer = (g.currPlayer + 1) % len(g.players)
	}

	if g.calendar.firstPlayer != -1 {
		g.FirstPlayer()
	}

	// todo food days
	g.Rotate()

	fmt.Fprintf(os.Stdout, "End of round\n")
	fmt.Fprintf(os.Stdout, "%s", g.String())
}

func (g *Game) FirstPlayer() {
	fmt.Fprintf(os.Stdout, "Firstplayering");
	worker := g.GetWorker(g.calendar.firstPlayer)
	fmt.Fprintf(os.Stdout, "Firstplayering %s\n", worker.color)
	worker.available = true
	worker.wheel_id = -1
	worker.position = -1
	g.calendar.firstPlayer = -1
	player := g.GetPlayerByColor(worker.color)
	fmt.Fprintf(os.Stdout, "giving player %d corn", player.corn)
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

	if !nextDayIsFoodDay && player.lightSide && g.rand.Intn(2) == 0 {
		// todo actually decide
		player.lightSide = false
		fmt.Fprintf(os.Stdout, "Player %s has gone dark\n", player.color)
		g.Rotate()
	}
}

func (g *Game) Rotate() {
	fmt.Fprintf(os.Stdout, "Rotating Calendar\n")
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
				g.currentBuildings = nil
				g.DealBuildings()
			} else {
				g.EndGame()
			}
			return
		}
	}
}

func (g *Game) EndGame() {
	for _, p := range g.players {
		p.points += TotalCorn(p) / 4
		p.points += p.resources[Skull] * 3

		for _, m := range p.monuments {
			p.points += m.GetPoints(g, p)
		}
	}

	fmt.Fprintf(os.Stdout, "%s", g.String())
	g.over = true
}

func (g *Game) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "Calendar State: \n%s\n", g.calendar.String(g.workers))

	for i := 0; i < len(g.players); i++ {
		fmt.Fprintf(&br, "%s", g.players[i].String())
	}
	fmt.Fprintf(&br, "Accumulated corn: %d\n", g.accumulatedCorn)

	fmt.Fprintf(&br, "%s\n%s", g.research.String(), g.temples.String())

	return br.String()
}

func (g *Game) FoodDay() {
	for _, player := range g.players {
		paid := 0
		unpaid := 0
		for _, w := range g.workers {
			if w.color == player.color {
				if w.wheel_id != -1 || w.available {
					if player.corn >= 2 - player.workerDeduction {
						player.corn -= 2 - player.workerDeduction
						paid += 1
					} else if unpaid < player.freeWorkers{
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
	if player.corn < 3 {
		player.corn = 3 // todo actually have begging
	}
	moves := g.GenerateMoves(g.players[g.currPlayer])
	move := moves[g.rand.Intn(len(moves))]

	fmt.Fprintf(os.Stdout, "Playing move %s for %s\n", move.String(), player.color)
	
	g.calendar.Execute(move, g)
	player.corn -= move.corn
	g.DealBuildings()
}

func (g *Game) DealBuildings() {
	for len(g.currentBuildings) < g.nBuildings {
		g.currentBuildings = append(g.currentBuildings, g.DealBuilding())
	}

	for i := 0; i < g.nBuildings; i++ {
		b := g.currentBuildings[i]
		for _, p := range g.players {
			for _, b2 := range p.buildings {
				if b2.id == b.id {
					g.currentBuildings[i] = g.DealBuilding()
				}
			}
		}
	}

	// for len(g.currentBuildings) < g.nBuildings {
	// 	var b Building
	// 	if g.age == 1 {
	// 		b, g.age1Buildings = g.age1Buildings[0], g.age1Buildings[1:]
	// 		g.currentBuildings = append(g.currentBuildings, b)
	// 	} else {
	// 		b, g.age2Buildings = g.age2Buildings[0], g.age2Buildings[1:]
	// 		g.currentBuildings = append(g.currentBuildings, b)
	// 	}
	// }
}

func (g *Game) DealBuilding() Building {
	var b Building
	if g.age == 1 {
		b, g.age1Buildings = g.age1Buildings[0], g.age1Buildings[1:]
	} else {
		b, g.age2Buildings = g.age2Buildings[0], g.age2Buildings[1:]
	}

	return b
}

func (g *Game) DealMonuments() {
	for len(g.currentMonuments) < g.nMonuments {
		var m Monument
		m, g.allMonuments = g.allMonuments[0], g.allMonuments[1:]
		g.currentMonuments = append(g.currentMonuments, m)
	}
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

	fmt.Fprintf(os.Stdout, "\t%s R %v P %v\n", p.color, retrieval, placement)

	retrieval_moves := append(make([]Move, 0), MakeEmptyRetrievalMove())
	moves = append(moves, g.MakeRetrievalMoves(retrieval_moves, retrieval)...)
	
	placement_moves := append(make([]Move, 0), MakeEmptyPlacementMove())
	moves = append(moves, g.MakePlacementMoves(placement_moves, placement)...)

	// todo find filter method
	out := make([]Move, 0)
	for _, move := range moves {
		if len(move.workers) > 0 && move.corn <= p.corn {
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
	// fmt.Fprintf(os.Stdout, "MakeRetrievalMoves %v %v\n", len(moves), retrieval)
// 
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
			new_game.calendar.Execute(moves[i], new_game)
			for _, option := range new_game.GetOptions(worker) {
				// if worker.wheel_id != 4 {
				// 	for j := 1; j < worker.position; j++ {

				// 		m = append(m, moves[i].Retrieve(w, &SpecificPosition {
				// 			wheel_id: worker.wheel_id,
				// 			corn: j,
				// 			Execute: option,
				// 		}, worker.position - j))
				// 	}
				// }
				m = append(m, moves[i].Retrieve(w, &SpecificPosition {
					wheel_id: worker.wheel_id,
					corn: worker.position,
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
			ws.execute(j)
			for p in ws.legal_places()
				m.append(j + place(p, i))

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
		new_game.calendar.Execute(moves[i], new_game)

		for _, position := range new_game.calendar.LegalPositions() {
			moves = append(moves, moves[i].Place(worker, position))
		}
	}

	return g.MakePlacementMoves(moves, rest)
}

func (g *Game) Clone() *Game {
	players := make([]*Player, 0)
	for _, p := range g.players {
		players = append(players, p.Clone())
	}

	workers := make([]*Worker, 0)
	for _, w := range g.workers {
		workers = append(workers, w.Clone())
	}

	new_calendar := g.calendar.Clone()
	new_temples := g.temples.Clone()
	new_research := g.research.Clone()

	currentBuildings := make([]Building, 0)
	for _, b := range g.currentBuildings {
		currentBuildings = append(currentBuildings, b)
	}

	currentMonuments := make([]Monument, 0)
	for _, m := range g.currentMonuments {
		currentMonuments = append(currentMonuments, m)
	}

	return &Game {
		players: players,
		workers: workers,
		calendar: new_calendar,
		temples: new_temples,
		research: new_research,
		nMonuments: g.nMonuments,
		currentMonuments: currentMonuments,
		allMonuments: g.allMonuments,
		nBuildings: g.nBuildings,
		currentBuildings: currentBuildings,
		age1Buildings: g.age1Buildings,
		age2Buildings: g.age2Buildings,
		currPlayer: g.currPlayer,
		firstPlayer: g.firstPlayer,
		accumulatedCorn: g.accumulatedCorn,
		age: g.age,
		day: g.day,
		resDays: g.resDays,
		pointDays: g.pointDays,
		over: g.over,
	}
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

func (g *Game) RemoveBuilding(b Building) {
	i := 0
	for _, b2 := range g.currentBuildings {
		if b2.id == b.id {
			break
		}
		i++
	}

	g.currentBuildings = remove(g.currentBuildings, i)
}

func (g *Game) RemoveMonument(m Monument) {
	i := 0
	for _, m2 := range g.currentMonuments {
		if m2.id == m.id {
			break
		}
		i++
	}

	g.currentMonuments = remove(g.currentMonuments, i)
}

func (g *Game) Over() bool {
	return g.over
}