package types

type Calendar struct {
	wheels []*Wheel
	rotation int
	firstPlayer int
	clone bool
	// tojm9 do food days & such
}