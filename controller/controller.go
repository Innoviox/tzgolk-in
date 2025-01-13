package controller

import (
    "math/rand"
    . "tzgolkin/model"
    . "tzgolkin/model/impl"
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
        []*Temple{Brown(), YellowT(), GreenT()},
    }

    age1Buildings := MakeAge1Buildings(g.rand)
    age2Buildings := MakeAge2Buildings(g.rand)
    monuments := MakeMonuments(g.rand)
    tiles := MakeWealthTiles(g.rand)

    game := &Game {
        calendar: MakeCalendar(wheels),
        temples: MakeTemples(temples),
        age1Buildings: age1Buildings,
        age2Buildings: age2Buildings,
        monuments: monuments,
        tiles: tiles,
    }

    game.Init()

    return &Controller {
        rand: rand,
        game: game,
    }
}