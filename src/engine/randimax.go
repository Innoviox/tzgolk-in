package engine

import (
    "fmt"
)
/*


compute(game_state, ply):
    if ply = 0:
        return evaluate(game_state)

    moves = generate_moves(game_state)

    best = -inf
    best_move = None
    for m in moves:
        new_game = game_state.clone()
        new_game.execute(m)
        # todo: new_game.execute_until_next_move(color)

        score = compute(new_game, ply - 1)
        if score > best:
            best = score
            best_move = m

    return best_move
*/

// import (
//     "fmt"
//     "os"
// )

import (
    "github.com/schollz/progressbar/v3"
    // "fmt"
    // "os"
)

func ComputeMove(g *Game, p *Player, ply int, rec bool) (*Move, float64) {
    if ply == 0 || g.Over {
        // return MakeEmptyMove(p.Color)
        return nil, p.Evaluate(g)
    }

    moves := g.GenerateMoves(p, ply)

    var bar *progressbar.ProgressBar
    if (!rec) {
        bar = progressbar.Default(int64(len(moves)))
    }

    ccp := g.CurrPlayer

    g2 := g.Clone()

    best := float64(-100)
    var best_move Move
    for _, m := range moves {
        // d := &Delta{}
        
        g.CurrPlayer = ccp
        // fmt.Println("TESTING MOVE", m, g.CurrPlayer, g.FirstPlayer)
        // fmt.Println("a", g.CurrPlayer, ccp)
        
        // fmt.Println("EXECUTING", m)
        d1 := g.Calendar.Execute(m, g, func(s string){/*fmt.Println(s)*/})
        g.AddDelta(d1, 1)
        // fmt.Println("aaa", d1.WorkerDeltas)
        // d.Add(d1)

        g.CurrPlayer = (ccp + 1) % len(g.Players)
        // fmt.Println("b", g.CurrPlayer, ccp)

        d2 := g.RunStop(func(s string){/*fmt.Println(s)*/}, p)
        // fmt.Println("bbb", d2.WorkerDeltas)
        // g.AddDelta(d2, 1)
        // d.Add(d2)
        
        _, score := ComputeMove(g, p, ply - 1, true)
        // if !rec {
        //     fmt.Fprintf(os.Stdout, "Score: %f for move %s\n", score, m.String())
        // }
        if score > best {
            best = score
            best_move = m
        }

        if bar != nil {
            bar.Add(1)
        }
        
        // g.Load(ply)
        d3 := Combine(d1, d2)
        g.AddDelta(d3, -1)
        // g.AddDelta(d2, -1)
        // g.AddDelta(d, -1)
        g.CurrPlayer = ccp
        // fmt.Println(d3, d3.FirstPlayer)
        if !g.Exact(g2) {
            fmt.Println("PLATO ERROR 2")
            fmt.Println("ccc", d3.ResearchDelta)

            fmt.Println([]int{}[1])
        }
        // g.AddDelta(d, -1)
    }   

    g.CurrPlayer = ccp

    return &best_move, best
}