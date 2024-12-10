// Convenient types and funcs for solving puzzles with grid
package grid

import "strings"

// Alias for [][]rune
type Grid [][]rune

func NewGrid(s string) Grid {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	grid := make([][]rune, len(lines))
	for i, l := range lines {
		grid[i] = []rune(l)
	}
	return grid
}

func (g Grid) ValueAt(p Pos) rune {
	return g[p.Y][p.X]
}

// Direction aka Vector(?)
type Dir struct {
	X int
	Y int
}

// Position aka Coordinate
type Pos struct {
	X int
	Y int
}

func (p Pos) IsInbound(g Grid) bool {
	maxX := len(g[0]) - 1
	maxY := len(g) - 1
	if 0 <= p.X && p.X <= maxX && 0 <= p.Y && p.Y <= maxY {
		return true
	}
	return false
}

// Move returns the new position after taking 'steps' at specified direction
func (p Pos) Move(d Dir, steps int) Pos {
	return Pos{
		p.X + d.X*steps,
		p.Y + d.Y*steps,
	}
}

// Move1 returns the new position after taking 1 step at specified direction
func (p Pos) Move1(d Dir) Pos {
	return p.Move(d, 1)
}

func Around4(p Pos) []Pos {
	out := make([]Pos, 4)
	for i, d := range FourDirs {
		out[i] = p.Move1(d)
	}
	return out
}

func Around8(p Pos) []Pos {
	out := make([]Pos, 8)
	for i, d := range EightDirs {
		out[i] = p.Move1(d)
	}
	return out
}

var (
	Left      = Dir{-1, 0}
	Right     = Dir{1, 0}
	Up        = Dir{0, -1}
	Down      = Dir{0, 1}
	UpLeft    = Dir{-1, -1}
	UpRight   = Dir{1, -1}
	DownLeft  = Dir{-1, 1}
	DownRight = Dir{1, 1}
	FourDirs  = []Dir{Up, Right, Down, Left}
	EightDirs = []Dir{Up, Right, Down, Left, UpLeft, UpRight, DownLeft, DownRight}
)
