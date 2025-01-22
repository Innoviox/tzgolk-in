package controller

import (
    // "fmt"
    // "os"
    "math/rand"
    . "tzgolkin/engine"
    . "tzgolkin/impl"
    . "tzgolkin/impl/wheels"
    . "tzgolkin/impl/buildings"
)

type Controller struct {
    rand *rand.Rand

    game *Game
}

func MakeController(rand *rand.Rand) *Controller {
    // todo make functions in impl for these
    wheels := []*Wheel{
        MakePalenque(), MakeYaxchilan(), MakeTikal(), MakeUxmal(), MakeChichen(),
    }

    temples := []*Temple{
        Brown(), YellowT(), GreenT(),
    }

    buildings := map[int]Building{}
    for k, v := range MakeAge1Buildings() {
        buildings[k] = v
    }
    for k, v := range MakeAge2Buildings() {
        buildings[k] = v
    }

    monuments := MakeMonuments()
    tiles := MakeWealthTiles(rand)

    game := &Game {
        Calendar: MakeCalendar(wheels),
        Temples: MakeTemples(temples),
        Buildings: buildings,
        Age1Cutoff: 15,
        Monuments: monuments,
        Tiles: tiles,
        Rand: rand,
    }

    game.Init()

    return &Controller {
        rand: rand,
        game: game,
    }
}

func (c *Controller) Run(MarkStep func(string)) {
    c.game.Run(MarkStep, true)
}

func (c *Controller) IsOver() bool {
    return c.game.Over
}

func (c *Controller) GetGame() *Game {
    return c.game
}