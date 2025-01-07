package model

type Temple struct {
	points []int
	resources []Resource // todo differentiate block vs skull
	startingStep int 
	ageRewards []int
}