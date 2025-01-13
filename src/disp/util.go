package disp

import (
    "strings"
)

func Convert(s string) [][]byte {
    var res [][]byte
    for _, r := range strings.Split(s, "\n") {
        res = append(res, []byte(r))
    }
    return res
}