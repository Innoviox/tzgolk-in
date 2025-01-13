package types

type Option struct {
	Execute func(*Game, *Player)
	description string
	buildingNum int
}

type Options func(*Game, *Player) []Option



const TempleDebug = "BYG"

type Resource int
const (
	Wood Resource = iota
	Stone
	Gold
	Skull
)
const ResourceDebug = "WSGC"

type Color int
const (
	Red Color = iota
	Green
	Blue
	Yellow
)

type Science int
const (
	Agriculture Science = iota
	Resources
	Construction
	Theology
)
const ResearchDebug = "ARCT"
