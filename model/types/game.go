package types

import (
    "math/rand"
)

type Game struct {
	// todo multiskull edge case
	players []*Player
	workers []*Worker

	calendar *Calendar 
	temples *Temples
	research *Research 
	
	nMonuments int
	currentMonuments []Monument
	allMonuments []Monument 
	
	nBuildings int
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

	over bool

	rand *rand.Rand
}