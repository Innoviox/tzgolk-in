package model

type Science int

// todo get actual names
const (
	Agriculture Science = iota
	Resources
	Construction
	Theology
)


type Research struct {

}

func (r *Research) HasLevel(c Color, s Science, level int) bool {
	return false
}