package disp

import (
    "fmt"
    "strings"
)

type Screen struct {
    width int
    height int

    grid [][]rune
}

// -- MARK -- Basic methods
func MakeScreen(width int, height int) *Screen {
    grid := make([][]rune, height)

    for i := 0; i < height; i++ {
        grid[i] = make([]rune, width)
        for j := 0; j < width; j++ {
            grid[i] = append(grid[i], ' ')
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
            // fmt.Fprintf(&br, "%q", s.grid[i][j])
            br.WriteRune(s.grid[i][j])
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
func (s *Screen) Set(x, y int, c rune) {
    s.grid[y][x] = c
}

func (s *Screen) Put(x, y int, grid [][]rune) {
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