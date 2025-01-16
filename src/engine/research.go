package engine

import (
	"fmt"
	"strings"
)

type Science int
type Levels map[Science]int

func MakeLevels() Levels {
	return Levels{
		Agriculture: 0,
		Resources: 0,
		Construction: 0,
		Theology: 0,
	}
}

type Research struct {
	Levels map[Color]Levels
}

func MakeResearch() *Research {
	return &Research{
		Levels: map[Color]Levels{
			Red: MakeLevels(),
			Green: MakeLevels(),
			Blue: MakeLevels(),
			Yellow: MakeLevels(),
		},
	}
}

func (r *Research) Clone() *Research {
	var newLevels = make(map[Color]Levels)

	for k, v := range r.Levels {
		newSci := make(Levels)
		for k2, v2 := range v {
			newSci[k2] = v2
		}
		newLevels[k] = newSci
	}

	return &Research {
		Levels: newLevels,
	}
}

func (r *Research) Copy(other *Research) {
	for k, v := range other.Levels {
		for k2, v2 := range v {
			r.Levels[k][k2] = v2
		}
	}
}

func (r *Research) AddDelta(delta ResearchDelta, mul int) {
	for c, l := range delta.Levels {
		for s, n := range l {
			r.Levels[c][s] += n * mul
		}
	}
}

func (r *Research) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "----Research------------\n")

	var matrix [][]strings.Builder

	for j := 0; j < 4; j++ {
		row := make([]strings.Builder, 0)
		for k := 0; k < 4; k++ {
			row = append(row, strings.Builder{})
		}
		matrix = append(matrix, row)
	}

	for c, l := range r.Levels {
		for s, n := range l {
			fmt.Fprintf(&matrix[int(s)][n], "%s", c.String())
		}
	}
	
	for i := 0; i < 4; i++ {
		fmt.Fprintf(&br, "|%s |", string(ResearchDebug[i]))

		for j := 0; j < 4; j++ {
			fmt.Fprintf(&br, "%-4s|", matrix[i][j].String())
		}

		fmt.Fprintf(&br, "\n")
	}

	fmt.Fprintf(&br, "------------------------\n")

	matrix = nil

	return br.String()
}

func (r *Research) HasLevel(c Color, s Science, level int) bool {
	return r.Levels[c][s] >= level
}

func (r *Research) Devout(c Color) bool {
	return r.HasLevel(c, Theology, 2)
}

func (r *Research) Foresight(c Color) bool {
	return r.HasLevel(c, Theology, 1)
}

func (r *Research) CornBonus(player Color, positionColor Color) int {
	if positionColor == Blue {
		if r.HasLevel(player, Agriculture, 2) {
			return 1
		} else {
			return 0
		}
	} else if positionColor == Green {
		if r.HasLevel(player, Agriculture, 1) {
			return 1
		} else if r.HasLevel(player, Agriculture, 3) {
			return 3
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func (r *Research) Irrigation(c Color) bool {
	return r.HasLevel(c, Agriculture, 2)
}

func (r *Research) ResourceBonus(c Color, res Resource) int {
	if r.HasLevel(c, Resources, int(res) + 1) {
		return 1
	} else {
		return 0
	}
}

func (r *Research) Built(p *Player) *Delta {
	if r.HasLevel(p.Color, Construction, 1) {
		p.Corn += 1
	}

	if r.HasLevel(p.Color, Construction, 2) {
		p.Points += 2
	}
}

func (r *Research) BuiltString(p *Player) string {
	var br strings.Builder

	if r.HasLevel(p.Color, Construction, 1) {
		fmt.Fprintf(&br, "1 Corn ")
	}

	if r.HasLevel(p.Color, Construction, 2) {
		fmt.Fprintf(&br, "2 points")
	}

	return br.String()
}

func (r *Research) Builder(c Color) bool {
	return r.HasLevel(c, Construction, 3)
}