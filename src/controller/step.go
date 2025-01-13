package controller

// import (
//     . "tzgolkin/model"
// )

// type Step struct {
//     move *Move
//     description string
// }

func (c *Controller) Step() {
    c.Round()
}