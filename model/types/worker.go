package types

type Worker struct {
	id int

	color Color

	available bool
	wheel_id int // use -1's
	position int
}