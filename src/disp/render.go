package disp

import (
    "fmt"
    "strings"
    . "tzgolkin/model"
)

func (d *Display) Render(step string) {
    g := d.controller.GetGame()

    for i, p := range g.Players {
        d.screen.Put(i * 25, 0, d.RenderPlayer(p, g))
    }

    d.screen.Put(0, 13, d.RenderGame(g))
    d.screen.Put(40, 13, d.RenderCalendar(g.Calendar))
    d.screen.Put(d.screen.width - 40, 0, d.RenderTemples(g.Temples))
    d.screen.Put(d.screen.width - 40, 15, d.RenderResearch(g.Research))
    for i, b := range g.CurrentBuildings {
        d.screen.Put(d.screen.width - 8 - 6 * i, 25, d.RenderBuilding(b))
    }

    d.screen.Put(10, 25, Convert(step))
}

func (d *Display) RenderPlayer(p *Player, g *Game) [][]rune {
    return Convert(p.String(g))
}

func (d *Display) RenderCalendar(c *Calendar) [][]rune {
    return Convert(c.String(d.controller.GetGame().Workers))
}

func (d *Display) RenderTemples(t *Temples) [][]rune {
    return Convert(t.String())
}

func (d *Display) RenderResearch(r *Research) [][]rune {
    return Convert(r.String())
}

func (d *Display) RenderBuilding(b Building) [][]rune {
    return Convert(b.String())
}

func (d *Display) RenderGame(g *Game) [][]rune {
    var br strings.Builder

    days := make([]string, 27)

    for i := 0; i < 27; i++ {
        days[i] = "_"
    }
    for _, d := range g.ResDays {
        days[d] = "R"
    }

    for _, d := range g.PointDays {
        days[d] = "P"
    }

    if g.Day < 27 {
        days[g.Day] = "X"
    }
    
    fmt.Fprintf(&br, "-----Game-----\n")
    fmt.Fprintf(&br, "|Age: %d\n", g.Age)
    fmt.Fprintf(&br, "|Days: ")
    for _, d := range days {
        fmt.Fprintf(&br, "%s", d)
    }
    fmt.Fprintf(&br, "\n|Accumulated Corn: %d\n", g.AccumulatedCorn)
    fmt.Fprintf(&br, "|First Player: %d\n", g.FirstPlayer)
    fmt.Fprintf(&br, "|Current Player: %d\n", g.CurrPlayer)
    fmt.Fprintf(&br, "--------------\n")

    return Convert(br.String())
}

// func (d *Display) RenderMonuments(m *Monuments) [][]byte {

// }

func (d *Display) RenderMove(m *Move) [][]rune {
    // todo make sure retrieval moves show the retrieval being chosen
    return Convert(m.String())
}