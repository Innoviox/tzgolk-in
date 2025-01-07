package model

type Calendar struct {
	wheels []*Wheel
	rotation int
	// todo food days & such
}

func (c *Calendar) Init() {
	c.wheels = make([]*Wheel, 0)
	c.rotation = 0

	for i := 0; i < 5; i++ {
		c.AddWheel(&Wheel {
			id: i,
			size: 5,
			occupied: make([]int, 0),
		})
	}
}

func (c *Calendar) AddWheel(w *Wheel) {
	c.wheels = append(c.wheels, w)
}

func (c *Calendar) Clone() Calendar {
	new_wheels := make([]*Wheel, 0) // todo map method?
	for _, wheel := range c.wheels {
		new_wheels = append(new_wheels, wheel.Clone())
	}

	return Calendar {
		wheels: new_wheels,
		rotation: c.rotation,
	}
}

func (c *Calendar) Execute(move Move) {
	if (move.placing) {
		for i := 0; i < len(move.workers); i++ {
			p := move.positions[i]
			c.wheels[p.wheel_id].AddWorker(p.position_num, move.workers[i])
		}
	} else {
		// todo
	}
}

func (c *Calendar) LegalPositions() []*Position {
	positions := make([]*Position, 0)

	for _, wheel := range c.wheels {
		i := 0
		for j := 0; j < len(wheel.occupied); j++ {
			if wheel.occupied[j] > i {
				break
			} else {
				i++
			}
		}
		
		positions = append(positions, &Position{
			wheel_id: wheel.id,
			position_num : i,
		})
	}

	return positions
}

// func (c *Calendar) setRotation(rotation int) {
// 	c.rotation = rotation
// 	for i := 0; i < len(c.wheels); i++ {
// 		c.wheels[i].setRotation(rotation)
// 	}
// }

// func (c *Calendar) rotate() {
// 	c.setRotation(c.rotation + 1);
// }