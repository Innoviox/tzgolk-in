package model

func Palenque0(g *Game, c Color, d int) {
	return
}

func Palenque1(g *Game, c Color, d int) {

}

func MakePalenque() Wheel {
	positions := make([]*Position, 0)
	positions = append(positions, &Position{
		wheel_id: 1,
		corn: 0,
		Execute: Palenque0,
		decisions: 0, 
	})


	return Wheel {
		id: 1,
		size: 7,
		occupied: make([]int, 0),
		workers: make([]int, 0),
		positions: positions,
		rotation: 0,
		name: "Palenque",
	}
}