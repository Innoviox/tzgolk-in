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
    for !c.game.IsOver() {
        MarkStep("Start of round\n")
    
        c.game.CurrPlayer = c.game.FirstPlayer
        for i := 0; i < len(c.game.Players); i++ {
            c.game.TakeTurn(MarkStep)
            c.game.CurrPlayer = (c.game.CurrPlayer + 1) % len(c.game.Players)
        }
    
        if c.game.Calendar.FirstPlayer != -1 {
            c.game.FirstPlayerSpace(MarkStep)
        }
    
        c.game.Rotate(MarkStep)

        MarkStep("End of round\n")
    
        // fmt.Fprintf(os.Stdout, "End of round\n")
        // fmt.Fprintf(os.Stdout, "%s", c.game.String())
    }

    // MarkStep("Game over\n")
}

func (c *Controller) IsOver() bool {
    return c.game.Over
}

func (c *Controller) GetGame() *Game {
    return c.game
}