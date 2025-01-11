package model

func Yaxchilan0(p *Player) {
	return
}

func Yaxchilan1(p *Player) {
	p.resources[Wood] += 1
}

func Yaxchilan2(p *Player) {
	p.resources[Stone] += 1
	p.corn += 1
}

func Yaxchilan3(p *Player) {
	p.resources[Gold] += 1
	p.corn += 2
}

func Yaxchilan4(p *Player) {
	p.resources[Skull] += 1
}

func Yaxchilan5(p *Player) {
	p.resources[Gold] += 1
	p.resources[Stone] += 1
	p.corn += 2
}

func MakeYaxchilan() *Wheel {
	positions := make([]*Position, 0)

	options := [][]Option{
		[]Option{PlayerOption(Yaxchilan0)},
		[]Option{PlayerOption(Yaxchilan1)},
		[]Option{PlayerOption(Yaxchilan2)},
		[]Option{PlayerOption(Yaxchilan3)},
		[]Option{PlayerOption(Yaxchilan4)},
		[]Option{PlayerOption(Yaxchilan5)},
	}

	for i := 0; i < 6; i++ {
		positions = append(positions, &Position{
			wheel_id: 2,
			corn: i,
			options: options[i],
		})
	}

	for i := 6; i < 8; i++ {
		positions = append(positions, &Position{
			wheel_id: 2,
			corn: i,
			options: flatten(options),
		})
	}

	return &Wheel{
		id: 2,
		size: len(positions),
		occupied: make([]int, 0),
		workers: make([]int, 0),
		positions: positions, 
		rotation: 0,
		name: "Yaxchilan",
	}
}