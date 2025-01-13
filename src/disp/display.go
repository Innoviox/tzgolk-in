package disp

import (
    "fmt"
    "bufio"
    "os"
    . "tzgolkin/controller"
)

type Display struct {
    screen *Screen
    controller *Controller
}

// -- MARK -- Basic methods
func MakeDisplay(controller *Controller) *Display {
    return &Display {
        screen: MakeScreen(150, 30),
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

func (d *Display) Run() {
    for !d.controller.IsOver() {
        d.controller.Step()
        fmt.Println(d.String())
        d.Hang()
    }
}

func (d *Display) Hang() {
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
}