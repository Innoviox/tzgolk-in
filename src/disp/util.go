package disp

import (
    "strings"
)

func Convert(s string) [][]rune {
    var res [][]rune
    for _, r := range strings.Split(s, "\n") {
        // res = append(res, []rune(r))
        row := make([]rune, 0)
        for _, c := range r {
            row = append(row, rune(c))
        }
        res = append(res, row)
    }
    return res
}