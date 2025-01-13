package model

import (
	"fmt"
	"os"
	"strings"
	"math/rand"
)

type Game struct {
	// todo multiskull edge case
	Players []*Player
	Workers []*Worker

	Calendar *Calendar 
	Temples *Temples
	Research *Research 
	
	NMonuments int
	CurrentMonuments []Monument
	AllMonuments []Monument 
	
	NBuildings int
	CurrentBuildings []Building
	Age1Buildings []Building
	Age2Buildings []Building 

	CurrPlayer int
	FirstPlayer int 

	AccumulatedCorn int

	Age int
	Day int
	ResDays []int
	PointDays []int

	Over bool

	Rand *rand.Rand

	Tiles []Tile
}

func (g *Game) Init() {
	g.Players = make([]*Player, 4)
	g.Workers = make([]*Worker, 0)
	for i, color := range [...]Color{Red, Green, Blue, Yellow} {
		g.Players[i] = &Player{
			Resources: [...]int{0, 0, 0, 0},
			Corn: 0,
			Color: color,
			Points: 0,
			FreeWorkers: 0,
			WorkerDeduction: 0,
			LightSide: true,
		}

		for j := 0; j < 6; j++ {
			g.Workers = append(g.Workers, &Worker{
				Id: i * 6 + j,
				Color: color,
				Available: j < 3,
				Wheel_id: -1,
				Position: -1,
			})
		}
	}

	g.Research = MakeResearch()

	g.NBuildings = 6
	g.CurrentBuildings = make([]Building, 0)
	g.DealBuildings()

	g.NMonuments = 6
	g.CurrentMonuments = make([]Monument, 0)
	g.DealMonuments()

	g.TileSetup()

	g.CurrPlayer = 0
	g.FirstPlayer = 0
	g.AccumulatedCorn = 0

	g.Age = 1
	g.Day = 0
	g.ResDays = []int{7, 20}
	g.PointDays = []int{13, 26}

	g.Over = false

	fmt.Fprintf(os.Stdout, "%s", g.String())
}

func (g *Game) TileSetup() {
	// todo: 4 choose 2
	t := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Fprintf(os.Stdout, "Placing tile %d for player %s\n", g.Tiles[t].N, g.GetPlayer(i).Color.String())
			g.Tiles[t].Execute(g, g.GetPlayer(i))
			t++
		}
	}
}

func (g *Game) Round() {
	if g.Over {
		return 
	}

	g.CurrPlayer = g.FirstPlayer
	for i := 0; i < len(g.Players); i++ {
		g.TakeTurn()
		g.CurrPlayer = (g.CurrPlayer + 1) % len(g.Players)
	}

	if g.Calendar.FirstPlayer != -1 {
		g.FirstPlayerSpace()
	}

	// todo food days
	g.Rotate()

	fmt.Fprintf(os.Stdout, "End of round\n")
	fmt.Fprintf(os.Stdout, "%s", g.String())
}

func (g *Game) FirstPlayerSpace() {
	fmt.Fprintf(os.Stdout, "Firstplayering");
	worker := g.GetWorker(g.Calendar.FirstPlayer)
	fmt.Fprintf(os.Stdout, "Firstplayering %s\n", worker.Color)
	worker.Available = true
	worker.Wheel_id = -1
	worker.Position = -1
	g.Calendar.FirstPlayer = -1
	player := g.GetPlayerByColor(worker.Color)
	fmt.Fprintf(os.Stdout, "giving player %d Corn", player.Corn)
	player.Corn += g.AccumulatedCorn
	g.AccumulatedCorn = 0

	playerIdx := 0
	for i := 0; i < len(g.Players); i++ {
		if g.Players[i].Color == player.Color {
			playerIdx = i
			break
		}
	}

	if g.FirstPlayer == playerIdx {
		g.FirstPlayer = (g.FirstPlayer + 1) % len(g.Players)
	} else {
		g.FirstPlayer = playerIdx
	}

	nextDayIsFoodDay := false
	for _, day := range g.ResDays {
		if g.Day == day - 1 {
			nextDayIsFoodDay = true
			break
		}
	}
	for _, day := range g.PointDays {
		if g.Day == day - 1 {
			nextDayIsFoodDay = true
			break
		}
	}

	if !nextDayIsFoodDay && player.LightSide && g.Rand.Intn(2) == 0 {
		// todo actually decide
		player.LightSide = false
		fmt.Fprintf(os.Stdout, "Player %s has gone dark\n", player.Color)
		g.Rotate()
	}
}

func (g *Game) Rotate() {
	fmt.Fprintf(os.Stdout, "Rotating Calendar\n")
	g.Calendar.Rotate(g)
	g.AccumulatedCorn += 1
	g.Day += 1
	g.CheckDay()
}

func (g *Game) CheckDay() {
	for _, day := range g.ResDays {
		if g.Day == day {
			g.FoodDay()

			for _, player := range g.Players {
				g.Temples.GainResources(player)
			}
			return
		}
	}

	for _, day := range g.PointDays {
		if g.Day == day {
			g.FoodDay()
			
			for _, player := range g.Players {
				g.Temples.GainPoints(player, g.Age)
			}

			g.Age += 1
			if g.Age == 2 {
				g.CurrentBuildings = nil
				g.DealBuildings()
			} else {
				g.EndGame()
			}
			return
		}
	}
}

func (g *Game) EndGame() {
	for _, p := range g.Players {
		p.Points += TotalCorn(p) / 4
		p.Points += p.Resources[Skull] * 3

		for _, m := range p.Monuments {
			p.Points += m.GetPoints(g, p)
		}
	}

	fmt.Fprintf(os.Stdout, "%s", g.String())
	g.Over = true
}

func (g *Game) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "Calendar State: \n%s\n", g.Calendar.String(g.Workers))

	for i := 0; i < len(g.Players); i++ {
		fmt.Fprintf(&br, "%s", g.Players[i].String())
	}
	fmt.Fprintf(&br, "Accumulated Corn: %d\n", g.AccumulatedCorn)

	fmt.Fprintf(&br, "%s\n%s", g.Research.String(), g.Temples.String())

	return br.String()
}

func (g *Game) FoodDay() {
	for _, player := range g.Players {
		paid := 0
		unpaid := 0
		for _, w := range g.Workers {
			if w.Color == player.Color {
				if w.Wheel_id != -1 || w.Available {
					if player.Corn >= 2 - player.WorkerDeduction {
						player.Corn -= 2 - player.WorkerDeduction
						paid += 1
					} else if unpaid < player.FreeWorkers{
						player.Points -= 3
						unpaid += 1
					}
				}
			}
		}
		fmt.Fprintf(os.Stdout, "Player %s paid %d workers, didn't pay %d workers\n", player.Color.String(), paid, unpaid)
	}
}

func (g *Game) TakeTurn() {
	player := g.Players[g.CurrPlayer]
	if player.Corn < 3 {
		player.Corn = 3 // todo actually have begging
	}
	moves := g.GenerateMoves(g.Players[g.CurrPlayer])
	move := moves[g.Rand.Intn(len(moves))]

	fmt.Fprintf(os.Stdout, "Playing move %s for %s\n", move.String(), player.Color)
	
	g.Calendar.Execute(move, g)
	player.Corn -= move.Corn
	g.DealBuildings()
}

func (g *Game) DealBuildings() {
	for len(g.CurrentBuildings) < g.NBuildings {
		g.CurrentBuildings = append(g.CurrentBuildings, g.DealBuilding())
	}

	for i := 0; i < g.NBuildings; i++ {
		b := g.CurrentBuildings[i]
		for _, p := range g.Players {
			for _, b2 := range p.Buildings {
				if b2.Id == b.Id {
					g.CurrentBuildings[i] = g.DealBuilding()
				}
			}
		}
	}

	// for len(g.CurrentBuildings) < g.NBuildings {
	// 	var b Building
	// 	if g.Age == 1 {
	// 		b, g.Age1Buildings = g.Age1Buildings[0], g.Age1Buildings[1:]
	// 		g.CurrentBuildings = append(g.CurrentBuildings, b)
	// 	} else {
	// 		b, g.Age2Buildings = g.Age2Buildings[0], g.Age2Buildings[1:]
	// 		g.CurrentBuildings = append(g.CurrentBuildings, b)
	// 	}
	// }
}

func (g *Game) DealBuilding() Building {
	var b Building
	if g.Age == 1 {
		b, g.Age1Buildings = g.Age1Buildings[0], g.Age1Buildings[1:]
	} else {
		b, g.Age2Buildings = g.Age2Buildings[0], g.Age2Buildings[1:]
	}

	return b
}

func (g *Game) DealMonuments() {
	for len(g.CurrentMonuments) < g.NMonuments {
		var m Monument
		m, g.AllMonuments = g.AllMonuments[0], g.AllMonuments[1:]
		g.CurrentMonuments = append(g.CurrentMonuments, m)
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
			new_game.Calendar.Execute(moves[i], new_game)
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
		new_game.Calendar.Execute(moves[i], new_game)

		for _, position := range new_game.Calendar.LegalPositions() {
			moves = append(moves, moves[i].Place(worker, position))
		}
	}

	return g.MakePlacementMoves(moves, rest)
}

func (g *Game) Clone() *Game {
	players := make([]*Player, 0)
	for _, p := range g.Players {
		players = append(players, p.Clone())
	}

	workers := make([]*Worker, 0)
	for _, w := range g.Workers {
		workers = append(workers, w.Clone())
	}

	new_calendar := g.Calendar.Clone()
	new_temples := g.Temples.Clone()
	new_research := g.Research.Clone()

	currentBuildings := make([]Building, 0)
	currentBuildings = append(currentBuildings, g.CurrentBuildings...)

	currentMonuments := make([]Monument, 0)
	currentMonuments = append(currentMonuments, g.CurrentMonuments...)

	return &Game {
		Players: players,
		Workers: workers,
		Calendar: new_calendar,
		Temples: new_temples,
		Research: new_research,
		NMonuments: g.NMonuments,
		CurrentMonuments: currentMonuments,
		AllMonuments: g.AllMonuments,
		NBuildings: g.NBuildings,
		CurrentBuildings: currentBuildings,
		Age1Buildings: g.Age1Buildings,
		Age2Buildings: g.Age2Buildings,
		CurrPlayer: g.CurrPlayer,
		FirstPlayer: g.FirstPlayer,
		AccumulatedCorn: g.AccumulatedCorn,
		Age: g.Age,
		Day: g.Day,
		ResDays: g.ResDays,
		PointDays: g.PointDays,
		Over: g.Over,
	}
}

func (g *Game) GetPlayer(num int) *Player {
	return g.Players[num]
}

func (g *Game) GetPlayerByColor(color Color) *Player {
	for _, player := range g.Players {
		if player.Color == color {
			return player
		}
	}

	return nil
}

func (g *Game) GetWorker(num int) *Worker {
	return g.Workers[num]
}

func (g *Game) UnlockWorker(color Color) {
	for _, w := range g.Workers {
		if w.Color == color {
			if !w.Available && w.Wheel_id == -1 {
				w.Available = true
				break
			}
		}
	}
}

func (g *Game) RemoveBuilding(b Building) {
	i := 0
	for _, b2 := range g.CurrentBuildings {
		if b2.Id == b.Id {
			break
		}
		i++
	}

	g.CurrentBuildings = remove(g.CurrentBuildings, i)
}

func (g *Game) RemoveMonument(m Monument) {
	i := 0
	for _, m2 := range g.CurrentMonuments {
		if m2.Id == m.Id {
			break
		}
		i++
	}

	g.CurrentMonuments = remove(g.CurrentMonuments, i)
}

func (g *Game) IsOver() bool {
	return g.Over
}