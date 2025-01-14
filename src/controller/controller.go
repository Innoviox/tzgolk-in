package controller

import (
    // "fmt"
    // "os"
    "math/rand"
    . "tzgolkin/model"
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

    age1Buildings := MakeAge1Buildings(rand)
    age2Buildings := MakeAge2Buildings(rand)
    monuments := MakeMonuments(rand)
    tiles := MakeWealthTiles(rand)

    game := &Game {
        Calendar: MakeCalendar(wheels),
        Temples: MakeTemples(temples),
        Age1Buildings: age1Buildings,
        Age2Buildings: age2Buildings,
        AllMonuments: monuments,
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
    c.game.Run(MarkStep, false, nil)
}

func (c *Controller) IsOver() bool {
    return c.game.Over
}

func (c *Controller) GetGame() *Game {
    return c.game
}