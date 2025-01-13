package types

type Temple struct {
	steps int
	playerLocations map[Color]int
	age1Prize int
	age2Prize int
	points []int
	resources map[int]Resource
}

type Temples struct {
	temples []*Temple
}
