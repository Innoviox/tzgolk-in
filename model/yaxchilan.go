package model

// todo func type
func Yaxchilan(i int) func(*Game, Color, int) {
	if i == 0 {
		return func(g *Game, c Color, d int) {
			return 
		}
	} else if i == 1 {
		return func(g *Game, c Color, d int) {
			p := g.GetPlayerByColor(c)

			p.resources[Wood] += 1
		}
	} else if i == 2 {
		return func(g *Game, c Color, d int) {
			p := g.GetPlayerByColor(c)

			p.resources[Stone] += 1
			p.corn += 1
		}
	} else if i == 3 {
		return func(g *Game, c Color, d int) {
			p := g.GetPlayerByColor(c)

			p.resources[Gold] += 1
			p.corn += 2
		}
	} else if i == 4 {
		return func(g *Game, c Color, d int) {
			p := g.GetPlayerByColor(c)

			p.resources[Skull] += 1
		}
	} else if i == 5 {
		return func(g *Game, c Color, d int) {
			p := g.GetPlayerByColor(c)

			p.resources[Gold] += 1
			p.resources[Stone] += 1
			p.corn += 2
		}
	} else if i == 6 || i == 7{
		return func(g *Game, c Color, d int) {
			Yaxchilan(d - 1)(g, c, d)
		}
	}
}

func MakeYaxchilan() Wheel {
	positions := make([]*Position, 0)
}