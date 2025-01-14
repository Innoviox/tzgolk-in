package model

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

    moves := g.GenerateMoves(p)

    var bar *progressbar.ProgressBar
    if (!rec) {
        bar = progressbar.Default(int64(len(moves)))
    }

    prev := g.Clone()

    best := float64(-100)
    var best_move Move
    for _, m := range moves {
        g.Calendar.Execute(m, g, func(s string){})
        g.CurrPlayer = (g.CurrPlayer + 1) % len(g.Players)
        g.RunStop(func(s string){/*fmt.Println(s)*/}, p)
        
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
        
        g.Copy(prev)
    }   

    return &best_move, best
}