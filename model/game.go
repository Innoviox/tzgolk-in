package model

type Game struct {
	players []Player
	calendar Calendar 
	temples []Temple 
	research []Research 

	currentMonuments []Monument
	allMonuments []Monument 

	currentBuildings []Building
	age1Buildings []Building
	age2Buildings []Building 

	currPlayer int
	firstPlayer int 
}

func (g *Game) round() {
	g.currPlayer = g.firstPlayer
	for i := 0; i < len(g.players); i++ {
		g.players[g.currPlayer].play()
		g.currPlayer = (g.currPlayer + 1) % len(g.players)
	}

	// todo first player nonsense
}