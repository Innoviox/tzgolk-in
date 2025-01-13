package disp

import (
    "fmt"
    "strings"
)

type Screen struct {
    width int
    height int

    grid [][]byte
}

// -- MARK -- Basic methods
func MakeScreen(width int, height int) *Screen {
    grid := make([][]byte, height)

    for i := 0; i < height; i++ {
        grid[i] = make([]byte, width)
        for j := 0; j < width; j++ {
            grid[i][j] = ' '
        }
    }

    return &Screen {
        width: width,
        height: height,
        grid: grid,
    }
}

func (s *Screen) String() string {
    var br strings.Builder

    for i := 0; i < s.width + 2; i++ {
        fmt.Fprintf(&br, "-")
    }
    fmt.Fprintf(&br, "\n")

    for i := 0; i < s.height; i++ {
        fmt.Fprintf(&br, "|")
        for j := 0; j < s.width; j++ {
            fmt.Fprintf(&br, "%c", s.grid[i][j])
        }
        fmt.Fprintf(&br, "|\n")
    }

    for i := 0; i < s.width + 2; i++ {
        fmt.Fprintf(&br, "-")
    }
    fmt.Fprintf(&br, "\n")

    return br.String()
}

// -- MARK -- Unique methods
func (s *Screen) Set(x, y int, c byte) {
    s.grid[y][x] = c
}

func (s *Screen) Put(x, y int, grid [][]byte) {
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            s.grid[y + i][x + j] = grid[i][j]
        }
    }
}

func (s *Screen) Clear() {
    for i := 0; i < s.height; i++ {
        for j := 0; j < s.width; j++ {
            s.grid[i][j] = ' '
        }
    }
}