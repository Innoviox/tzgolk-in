package wheel

type Calendar struct {
	wheels []*Wheel
	rotation int
	// todo food days & such
}

func (c *Calendar) Init() {
	c.wheels = make([]*Wheel, 0)
	c.rotation = 0

	for i := 0; i < 5; i++ {
		c.addWheel(&Wheel {
			id: i,
			size: 5,
			occupied: make([]int, 0)
		})
	}
}

func (c *Calendar) AddWheel(w *Wheel) {
	c.wheels = append(c.wheels, w)
}

func (c *Calendar) setRotation(rotation int) {
	c.rotation = rotation
	for i := 0; i < len(c.wheels); i++ {
		c.wheels[i].setRotation(rotation)
	}
}

func (c *Calendar) rotate() {
	c.setRotation(c.rotation + 1);
}