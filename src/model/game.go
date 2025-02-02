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

	Freeze map[int]*Game
}

// -- MARK -- Basic methods
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

	g.Freeze = make(map[int]*Game)

	fmt.Fprintf(os.Stdout, "%s", g.String())
}

func (g *Game) Clone() *Game {
	fmt.Println("Cloning game!")
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
		Rand: g.Rand,
		Freeze: g.Freeze,
	}
}

func (g *Game) Copy(other *Game) {
	for i := 0; i < len(g.Players); i++ {
		g.Players[i].Copy(other.Players[i])
	}

	for i := 0; i < len(g.Workers); i++ {
		g.Workers[i].Copy(other.Workers[i])
	}

	g.Calendar.Copy(other.Calendar)
	g.Temples.Copy(other.Temples)
	g.Research.Copy(other.Research)

	g.NMonuments = other.NMonuments
	g.CurrentMonuments = nil
	g.CurrentMonuments = append(g.CurrentMonuments, other.CurrentMonuments...)

	g.AllMonuments = nil
	g.AllMonuments = append(g.AllMonuments, other.AllMonuments...)

	g.NBuildings = other.NBuildings
	g.CurrentBuildings = nil
	g.CurrentBuildings = append(g.CurrentBuildings, other.CurrentBuildings...)

	g.Age1Buildings = nil
	g.Age1Buildings = append(g.Age1Buildings, other.Age1Buildings...)

	g.Age2Buildings = nil
	g.Age2Buildings = append(g.Age2Buildings, other.Age2Buildings...)

	g.CurrPlayer = other.CurrPlayer
	g.FirstPlayer = other.FirstPlayer
	g.AccumulatedCorn = other.AccumulatedCorn
	g.Age = other.Age
	g.Day = other.Day
	g.ResDays = other.ResDays
	g.PointDays = other.PointDays
	g.Over = other.Over
}

func (g *Game) Save(key int) {
	res, ok := g.Freeze[key]
	if ok {
		res.Copy(g)
	} else {
		g.Freeze[key] = g.Clone()
	}
}

func (g *Game) Load(key int) {
	g.Copy(g.Freeze[key])
}

func (g *Game) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "Calendar State: \n%s\n", g.Calendar.String(g.Workers))

	for i := 0; i < len(g.Players); i++ {
		fmt.Fprintf(&br, "%s", g.Players[i].String(g))
	}
	fmt.Fprintf(&br, "Accumulated Corn: %d\n", g.AccumulatedCorn)

	fmt.Fprintf(&br, "%s\n%s", g.Research.String(), g.Temples.String())

	return br.String()
}

// -- MARK -- Setup methods
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

// -- MARK -- Flow methods
func (g *Game) FirstPlayerSpace(MarkStep func(string)) {
	// fmt.Fprintf(os.Stdout, "Firstplayering");
	worker := g.GetWorker(g.Calendar.FirstPlayer)
	// fmt.Fprintf(os.Stdout, "Firstplayering %s\n", worker.Color)
	worker.Available = true
	worker.Wheel_id = -1
	worker.Position = -1
	g.Calendar.FirstPlayer = -1
	player := g.GetPlayerByColor(worker.Color)
	// fmt.Fprintf(os.Stdout, "giving player %d Corn", player.Corn)
	player.Corn += g.AccumulatedCorn
	MarkStep(fmt.Sprintf("FirstPlayer for %s, +%d Corn", player.Color.String(), g.AccumulatedCorn))

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
		// fmt.Fprintf(os.Stdout, "Player %s has gone dark\n", player.Color)
		MarkStep(fmt.Sprintf("Player %s has gone dark", player.Color.String()))
		g.Rotate(MarkStep)
	}
}

func (g *Game) Rotate(MarkStep func(string)) {
	// fmt.Fprintf(os.Stdout, "Rotating Calendar\n")
	MarkStep("Rotating Calendar")
	g.Calendar.Rotate(g)
	g.AccumulatedCorn += 1
	MarkStep("Rotated Calendar")
	g.CheckDay(MarkStep)
	g.Day += 1
}

func (g *Game) CheckDay(MarkStep func(string)) {
	for _, day := range g.ResDays {
		if g.Day == day {
			g.FoodDay(MarkStep)

			for _, player := range g.Players {
				g.Temples.GainResources(player)
			}

			MarkStep("Gained resources")

			return
		}
	}

	for _, day := range g.PointDays {
		if g.Day == day {
			g.FoodDay(MarkStep)
			
			for _, player := range g.Players {
				g.Temples.GainPoints(player, g.Age)
			}

			MarkStep("Gained points")

			g.Age += 1
			if g.Age == 2 {
				g.CurrentBuildings = make([]Building, 0)
				g.DealBuildings()
				MarkStep("Dealt new buildings")
			} else {
				g.EndGame(MarkStep)
			}
			return
		}
	}
}

func (g *Game) EndGame(MarkStep func(string)) {
	for _, p := range g.Players {
		p.Points += p.TotalCorn() / 4
		p.Points += p.Resources[Skull] * 3

		for _, m := range p.Monuments {
			p.Points += m.GetPoints(g, p)
		}
		MarkStep(fmt.Sprintf("Gained endgame points for %s", p.Color.String()))
	}

	g.Over = true
}

func (g *Game) FoodDay(MarkStep func(string)) {
	for _, player := range g.Players {
		paid := 0
		unpaid := 0
		for _, w := range g.Workers {
			if w.Color == player.Color {
				if w.Wheel_id != -1 || w.Available {
					if player.Corn >= 2 - player.WorkerDeduction {
						player.Corn -= 2 - player.WorkerDeduction
						paid += 1
					} else if unpaid >= player.FreeWorkers {
						player.Points -= 3
						unpaid += 1
					} else {
						unpaid += 1
					}
				}
			}
		}
		// fmt.Fprintf(os.Stdout, "Player %s paid %d workers, didn't pay %d workers\n", player.Color.String(), paid, unpaid)
		MarkStep(fmt.Sprintf("Player %s paid %d workers, didn't pay %d workers", player.Color.String(), paid, unpaid))
	}
}

func (g *Game) TakeTurn(MarkStep func(string), random bool) {
	player := g.Players[g.CurrPlayer]

	var move *Move
	if random {
		// this number is equal to ply + 1
		moves := g.GenerateMoves(g.Players[g.CurrPlayer], 3)
		if len(moves) > 0 {
			move = &moves[g.Rand.Intn(len(moves))]
		}
	} else {
		move, _ = ComputeMove(g, player, 2, false)
	}

	// fmt.Fprintf(os.Stdout, "Playing move %s for %s\n", move.String(), player.Color)
	
	if move != nil {
		MarkStep(fmt.Sprintf("Playing move %s for %s", move.String(), player.Color.String()))
		g.Calendar.Execute(*move, g, MarkStep)
	} else {
		MarkStep(fmt.Sprintf("[FATAL ERROR] No move for %s", player.Color.String()))
	}
	// player.Corn -= move.Corn

	g.DealBuildings()
}

func (g *Game) Run(MarkStep func(string), random bool, stopPlayer *Player) {
	for !g.IsOver() && g.Day < 4 {
		if stopPlayer == nil {
			g.CurrPlayer = g.FirstPlayer
		}
		for i := 0; i < len(g.Players); i++ {
			if stopPlayer != nil && g.Players[g.CurrPlayer].Color == stopPlayer.Color {
				MarkStep("Stopped by color")
				return
			}
			g.TakeTurn(MarkStep, random)
			g.CurrPlayer = (g.CurrPlayer + 1) % len(g.Players)
		}

		if g.Calendar.FirstPlayer != -1 {
			g.FirstPlayerSpace(MarkStep)
		}

		g.Rotate(MarkStep)
	}
}

func mod(a, b int) int {
    return (a % b + b) % b
}

func (g *Game) RunStop(MarkStep func(string), stopPlayer *Player) {
	// current player is set to stopPlayer + 1
	run1 := mod(g.FirstPlayer - int(stopPlayer.Color) + 3, 4)
	MarkStep(fmt.Sprintf("Running for %d players (%d %d)", run1, g.FirstPlayer, int(stopPlayer.Color)))
	for i := 0; i < run1; i++ {
		g.TakeTurn(MarkStep, true)
		g.CurrPlayer = (g.CurrPlayer + 1) % len(g.Players)
	}

	if g.Calendar.FirstPlayer != -1 {
		g.FirstPlayerSpace(MarkStep)
	}

	g.Rotate(MarkStep)

	run2 := 3 - run1
	for i := 0; i < run2; i++ {
		g.TakeTurn(MarkStep, true)
		g.CurrPlayer = (g.CurrPlayer + 1) % len(g.Players)
	}
	MarkStep(fmt.Sprintf("Stopped by color %s", stopPlayer.Color.String()))
	/*
	fp sp 0 1 2 3
	0     3 2 1 0
	1     0 3 2 1
	2     1 0 3 2
	3     2 1 0 3
	*/
	// run for the next n players, n = mod(4 - fp - sp - 1, 4)

	
}

// -- MARK -- Getters
func (g *Game) DealBuildings() {
	for len(g.CurrentBuildings) < g.NBuildings && g.CanDealBuilding() {
		g.CurrentBuildings = append(g.CurrentBuildings, g.DealBuilding())
	}

	for i := 0; i < g.NBuildings; i++ {
		b := g.CurrentBuildings[i]
		for _, p := range g.Players {
			for _, b2 := range p.Buildings {
				if b2.Id == b.Id && g.CanDealBuilding() {
					g.CurrentBuildings[i] = g.DealBuilding()
				}
			}
		}
	}
}

func (g *Game) CanDealBuilding() bool {
	if g.Age == 1 {
		return len(g.Age1Buildings) > 0 
	} else {
		return len(g.Age2Buildings) > 0
	}
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