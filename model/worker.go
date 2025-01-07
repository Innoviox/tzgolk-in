package model

type Worker struct {
	id int

	color string // todo Color type

	available bool
	wheel_id int// | nil // ?
	position int// | nil
}