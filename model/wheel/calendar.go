package wheel

type Calendar struct {
	wheels []*Wheel
	rotation int
}

func (c *Calendar) AddWheel(w Wheel) {
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