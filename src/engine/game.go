package engine

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
	// 0 => in deck
	// 1 => dealt (showing now)
	// 2 => built
	CurrentMonuments map[int]int
	Monuments map[int]Monument
	
	NBuildings int
	CurrentBuildings map[int]int
	Buildings map[int]Building
	Age1Cutoff int

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
	g.CurrentBuildings = make(map[int]int)
	for i := 0; i < len(g.Buildings); i++ {
		g.CurrentBuildings[i] = 0
	}
	g.CurrentMonuments = make(map[int]int)
	for i := 0; i < len(g.Monuments); i++ {
		g.CurrentMonuments[i] = 0
	}

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
			Buildings: CopyMap(g.CurrentBuildings),
			Monuments: CopyMap(g.CurrentMonuments),
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
	g.AddDelta(g.DealBuildings(), 1)

	g.NMonuments = 6
	g.AddDelta(g.DealMonuments(), 1)

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

	currentBuidings := CopyMap(g.CurrentBuildings)

	currentMonuments := CopyMap(g.CurrentMonuments)

	return &Game {
		Players: players,
		Workers: workers,
		Calendar: new_calendar,
		Temples: new_temples,
		Research: new_research,
		CurrentMonuments: currentMonuments,
		NMonuments: g.NMonuments,
		Monuments: g.Monuments,
		NBuildings: g.NBuildings,
		CurrentBuildings: currentBuidings,
		Buildings: g.Buildings,
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
	g.CurrentMonuments = CopyMap(other.CurrentMonuments)

	g.NBuildings = other.NBuildings
	g.CurrentBuildings = CopyMap(other.CurrentBuildings)


	g.CurrPlayer = other.CurrPlayer
	g.FirstPlayer = other.FirstPlayer
	g.AccumulatedCorn = other.AccumulatedCorn
	g.Age = other.Age
	g.Day = other.Day
	g.ResDays = other.ResDays
	g.PointDays = other.PointDays
	g.Over = other.Over
}

func (g *Game) AddDelta(delta *Delta, mul int) {
	for _, p := range g.Players {
		res, ok := delta.PlayerDeltas[p.Color]
		if ok {
			p.AddDelta(res, mul)
		}
	}

	for _, w := range g.Workers {
		w.AddDelta(delta.WorkerDeltas[w.Id], mul)
	}

	g.Calendar.AddDelta(delta.CalendarDelta, mul)
	g.Temples.AddDelta(delta.TemplesDelta, mul)
	g.Research.AddDelta(delta.ResearchDelta, mul)

	for k, v := range delta.Buildings {
		g.CurrentBuildings[k] += v * mul 
	}

	for k, v := range delta.Monuments {
		g.CurrentMonuments[k] += v * mul
	}

	g.CurrPlayer += delta.CurrPlayer * mul
	g.FirstPlayer += delta.FirstPlayer * mul
	g.AccumulatedCorn += delta.AccumulatedCorn * mul
	g.Age += delta.Age * mul
	g.Day += delta.Day * mul
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
func (g *Game) FirstPlayerSpace(MarkStep func(string)) *Delta {
	worker := g.GetWorker(g.Calendar.FirstPlayer)
	player := g.GetPlayerByColor(worker.Color)

	d := &Delta{
		PlayerDeltas: map[Color]PlayerDelta{player.Color: PlayerDelta{
			Corn: g.AccumulatedCorn,
		}},
		WorkerDeltas: map[int]WorkerDelta{worker.Id: WorkerDelta{
			Available: 1,
			Wheel_id: -1 - worker.Wheel_id, 
			Position: -1 - worker.Position,
		}},
		CalendarDelta: CalendarDelta{FirstPlayer: -1 - g.Calendar.FirstPlayer},
		AccumulatedCorn: -g.AccumulatedCorn,
	}

	MarkStep(fmt.Sprintf("FirstPlayer for %s, +%d Corn", player.Color.String(), g.AccumulatedCorn))
	
	playerIdx := 0
	for i := 0; i < len(g.Players); i++ {
		if g.Players[i].Color == player.Color {
			playerIdx = i
			break
		}
	}

	newFirstPlayer := playerIdx
	if g.FirstPlayer == playerIdx {
		newFirstPlayer = (g.FirstPlayer + 1) % len(g.Players)
	}

	d.FirstPlayer = newFirstPlayer - g.FirstPlayer

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
		d.Add(PlayerDeltaWrapper(player.Color, PlayerDelta{
			LightSide: -1,
		}))
		// fmt.Fprintf(os.Stdout, "Player %s has gone dark\n", player.Color)
		MarkStep(fmt.Sprintf("Player %s has gone dark", player.Color.String()))
		d.Add(g.Rotate(MarkStep))
	}

	return d
}

func (g *Game) Rotate(MarkStep func(string)) *Delta {
	MarkStep("Rotating Calendar")
	delta := g.Calendar.Rotate(g)
	delta.AccumulatedCorn = 1
	MarkStep("Rotated Calendar")
	delta.Add(g.CheckDay(MarkStep))
	delta.Day = 1

	return delta
}

func (g *Game) CheckDay(MarkStep func(string)) *Delta {
	d := &Delta{}
	for _, day := range g.ResDays {
		if g.Day == day {
			d.Add(g.FoodDay(MarkStep))

			for _, player := range g.Players {
				d.Add(g.Temples.GainResources(player))
			}

			MarkStep("Gained resources")

			return d
		}
	}

	for _, day := range g.PointDays {
		if g.Day == day {
			d.Add(g.FoodDay(MarkStep))
			
			for _, player := range g.Players {
				d.Add(g.Temples.GainPoints(player, g.Age))
			}

			MarkStep("Gained points")

			d.Add(&Delta{Age: 1})
			if g.Age == 1 {
				// todo buildings
				for k, v := range g.CurrentBuildings {
					if v == 1 {
						d.Buildings[k] = -1
					}
				}
				g.DealBuildings()
				MarkStep("Dealt new buildings")
			} else {
				d.Add(g.EndGame(MarkStep))
			}
			return d
		}
	}
	return d
}

func (g *Game) EndGame(MarkStep func(string)) *Delta {
	d := &Delta{PlayerDeltas: map[Color]PlayerDelta{}}
	for _, p := range g.Players {
		i := 0
		i += p.TotalCorn() / 4
		i += p.Resources[Skull] * 3

		for k, v := range p.Monuments {
			if v == 1 {
				i += g.Monuments[k].GetPoints(g, p)
			}
		}
		MarkStep(fmt.Sprintf("Gained endgame points for %s", p.Color.String()))
		d.PlayerDeltas[p.Color] = PlayerDelta{Points: i}
	}

	d.Over = 1
	return d
}

func (g *Game) FoodDay(MarkStep func(string)) *Delta {
	d := &Delta{PlayerDeltas: map[Color]PlayerDelta{}}
	for _, player := range g.Players {
		paid := 0
		unpaid := 0
		pd := PlayerDelta{}
		for _, w := range g.Workers {
			if w.Color == player.Color {
				if w.Wheel_id != -1 || w.Available {
					if player.Corn >= 2 - player.WorkerDeduction {
						pd.Corn -= 2 - player.WorkerDeduction
						paid += 1
					} else if unpaid >= player.FreeWorkers {
						pd.Points -= 3
						unpaid += 1
					} else {
						unpaid += 1
					}
				}
			}
		}
		d.PlayerDeltas[player.Color] = pd
		// fmt.Fprintf(os.Stdout, "Player %s paid %d workers, didn't pay %d workers\n", player.Color.String(), paid, unpaid)
		MarkStep(fmt.Sprintf("Player %s paid %d workers, didn't pay %d workers", player.Color.String(), paid, unpaid))
	}
	return d
}

func (g *Game) TakeTurn(MarkStep func(string), random bool) *Delta {
	d := &Delta{}
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
		d.Add(g.Calendar.Execute(*move, g, MarkStep))
	} else {
		MarkStep(fmt.Sprintf("[FATAL ERROR] No move for %s", player.Color.String()))
	}
	// player.Corn -= move.Corn

	d.Add(g.DealBuildings())
	return d
}

func (g *Game) Run(MarkStep func(string), random bool) {
	for !g.IsOver() /* && g.Day < 4 */ {
		for i := 0; i < len(g.Players); i++ {
			g.AddDelta(g.TakeTurn(MarkStep, random), 1)
			MarkStep("Test")
			g.CurrPlayer = (g.CurrPlayer + 1) % len(g.Players)
		}

		if g.Calendar.FirstPlayer != -1 {
			g.AddDelta(g.FirstPlayerSpace(MarkStep), 1)
		}

		g.AddDelta(g.Rotate(MarkStep), 1)
	}
}

func mod(a, b int) int {
    return (a % b + b) % b
}

func (g *Game) RunStop(MarkStep func(string), stopPlayer *Player) *Delta {
	d := &Delta{}
	// current player is set to stopPlayer + 1
	run1 := mod(g.FirstPlayer - int(stopPlayer.Color) + 3, 4)
	MarkStep(fmt.Sprintf("Running for %d players (%d %d)", run1, g.FirstPlayer, int(stopPlayer.Color)))
	for i := 0; i < run1; i++ {
		d1 := g.TakeTurn(MarkStep, true)
		d.Add(d1)
		g.AddDelta(d1, 1)

		g.CurrPlayer = (g.CurrPlayer + 1) % len(g.Players)
	}

	if g.Calendar.FirstPlayer != -1 {
		d1 := g.FirstPlayerSpace(MarkStep)
		d.Add(d1)
		g.AddDelta(d1, 1)
	}

	d1 := g.Rotate(MarkStep)
	d.Add(d1)
	g.AddDelta(d1, 1)

	run2 := 3 - run1
	for i := 0; i < run2; i++ {
		d2 := g.TakeTurn(MarkStep, true)
		d.Add(d2)
		g.AddDelta(d2, 1)

		g.CurrPlayer = (g.CurrPlayer + 1) % len(g.Players)
	}
	MarkStep(fmt.Sprintf("Stopped by color %s", stopPlayer.Color.String()))
	return d

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
func (g *Game) DealBuildings() *Delta {
	d := &Delta{
		Buildings: map[int]int{},
	}
	n := 0
	for _, v := range g.CurrentBuildings {
		if v == 1 {
			n += 1
		}
	}
	lo, hi := 0, g.Age1Cutoff
	if g.Age == 2 {
		lo, hi = g.Age1Cutoff, len(g.Buildings)
	}

	r := RandZeros(g.CurrentBuildings, lo, hi, g.NBuildings - n)
	for _, i := range r {
		d.Buildings[i] = 1
	}

	return d
}

func (g *Game) DealMonuments() *Delta {
	d := &Delta{
		Monuments: map[int]int{},
	}
	r := RandZeros(g.CurrentMonuments, 0, len(g.Monuments), g.NMonuments)
	for _, n := range r {
		d.Monuments[n] = 1
	}
	return d
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

func (g *Game) UnlockWorker(color Color) *Delta {
	d := &Delta{WorkerDeltas: map[int]WorkerDelta{}}
	for _, w := range g.Workers {
		if w.Color == color {
			if !w.Available && w.Wheel_id == -1 {
				d.WorkerDeltas[w.Id] = WorkerDelta{Available: 1}
				// w.Available = true
				break
			}
		}
	}
	return d
}

func (g *Game) RemoveBuilding(b int) *Delta {
	d := &Delta{}
	d.Buildings[b] = 1
	return d
}

func (g *Game) RemoveMonument(m int) *Delta {
	d := &Delta{}
	d.Buildings[m] = 1
	return d
}

func (g *Game) IsOver() bool {
	return g.Over
}