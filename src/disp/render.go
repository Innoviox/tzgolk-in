package disp

import (
    . "tzgolkin/model"
)

func (d *Display) Render() {
    // d.controller.Render(d.screen)
}

func (d *Display) RenderPlayer(p *Player) [][]byte {
    return Convert(p.String())
}

func (d *Display) RenderCalendar(c *Calendar) [][]byte {
    return Convert(c.String(d.controller.GetGame().Workers))
}

func (d *Display) RenderTemples(t *Temples) [][]byte {
    return Convert(t.String())
}

func (d *Display) RenderResearch(r *Research) [][]byte {
    return Convert(r.String())
}

func (d *Display) RenderBuildings(b *Buildings) [][]byte {

}

func (d *Display) RenderMonuments(m *Monuments) [][]byte {

}

func (d *Display) RenderMove(m *Move) [][]byte {
    // todo make sure retrieval moves show the retrieval being chosen
    return Convert(m.String())
}