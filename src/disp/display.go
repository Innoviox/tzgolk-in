package disp

import (
    "fmt"
    . "tzgolkin/controller"
)

type Display struct {
    screen *Screen
    controller *Controller
}

// -- MARK -- Basic methods
func MakeDisplay(controller *Controller) *Display {
    return &Display {
        screen: MakeScreen(150, 100),
        controller: controller,
    }
}

func (d *Display) String() string {
    d.Blit()
    return d.screen.String()
}

// -- MARK -- Unique methods
func (d *Display) Blit() {
    d.screen.Clear()
    d.Render()
}

func (d *Display) RunGame() {
    for !d.controller.IsOver() {
        d.controller.Step()
        fmt.Println(d.String())
    }
}