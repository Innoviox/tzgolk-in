package controller

import (
    "math/rand"
    . "tzgolkin/model"
    . "tzgolkin/model/impl"
    . "tzgolkin/model/impl/wheels"
    . "tzgolkin/model/impl/buildings"
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
    }

    game.Init()

    return &Controller {
        rand: rand,
        game: game,
    }
}

func (c *Controller) RunGame() {
    for !c.game.IsOver() {
        c.game.Round()
    }
}