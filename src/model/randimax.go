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
)

func ComputeMove(g *Game, p *Player, ply int, rec bool) (Move, int) {
    if ply == 0 || g.Over {
        // return MakeEmptyMove(p.Color)
        return Move{}, p.Evaluate(g)
    }

    moves := g.GenerateMoves(p)

    var bar *progressbar.ProgressBar
    if (!rec) {
        bar = progressbar.Default(int64(len(moves)))
    }

    best := -100
    var best_move Move
    for _, m := range moves {
        new_game := g.Clone()
        new_game.Calendar.Execute(m, new_game, func(s string){})
        new_game.Run(func(s string){}, true, &p.Color)

        _, score := ComputeMove(new_game, p, ply - 1, true)
        if score > best {
            best = score
            best_move = m
        }

        if bar != nil {
            bar.Add(1)
        }
    }   

    return best_move, best
}