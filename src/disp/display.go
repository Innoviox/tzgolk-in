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

func (d *Display) String(step string) string {
    d.Blit(step)
    return d.screen.String()
}

// -- MARK -- Unique methods
func (d *Display) Blit(step string) {
    d.screen.Clear()
    d.Render(step)
}

func (d *Display) Run() {
    d.controller.Run(d.MarkStep)
}

func (d *Display) MarkStep(step string) {
    fmt.Println(d.String(step))
    // d.Hang()
}

func (d *Display) Hang() {
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
}