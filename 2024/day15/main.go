package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/quangd42/aoc/grid"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimSpace(input)
	st1 := time.Now()
	fmt.Printf("Part 1: %d     %v\n", part1(input), time.Since(st1))
	st2 := time.Now()
	fmt.Printf("Part 2: %d     %v\n", part2(input), time.Since(st2))
}

type object struct {
	grid.Pos
	val rune
}

func objectAt(g grid.Grid, p grid.Pos) object {
	return object{
		Pos: p,
		val: g.ValueAt(p),
	}
}

func move(p grid.Pos, g grid.Grid, d grid.Dir) bool {
	// find dest pos
	dest := p.Move1(d)
	// if dest pos == '#', return, move nothing
	if g.ValueAt(dest) == '#' {
		return false
	}
	// if dest pos == 'O'
	if g.ValueAt(dest) == 'O' {
		//   try to move dest pos, if false move nothing
		if !move(dest, g, d) {
			return false
		}
		//   if true, move self to dest pos
	}

	// if dest pos == '.' ( base case), move self to dest pos
	g[dest.Y][dest.X] = g.ValueAt(p)
	g[p.Y][p.X] = '.'
	return true
}

func getInstructions(s string) []grid.Dir {
	out := make([]grid.Dir, len(s))
	for i, r := range s {
		switch r {
		case '<':
			out[i] = grid.DirLeft
		case '>':
			out[i] = grid.DirRight
		case '^':
			out[i] = grid.DirUp
		case 'v':
			out[i] = grid.DirDown
		}
	}
	return out
}

func part1(input string) int {
	out := 0
	gStr, iStr, found := strings.Cut(strings.TrimSpace(input), "\n\n")
	if !found {
		log.Fatal("bad input, no empty line")
	}
	g := grid.NewGrid(gStr)
	ins := getInstructions(iStr)

	var rb grid.Pos
	for y, row := range g {
		for x, r := range row {
			if r == '@' {
				rb = grid.Pos{X: x, Y: y}
			}
		}
	}
	for _, in := range ins {
		if move(rb, g, in) {
			rb = rb.Move1(in)
		}
	}

	for y, row := range g {
		for x, r := range row {
			if r == 'O' {
				out += x + 100*y
			}
		}
	}
	return out
}

func newGridWide(g grid.Grid) grid.Grid {
	out := make([][]rune, len(g))
	for y, row := range g {
		for _, r := range row {
			switch r {
			case '#':
				out[y] = append(out[y], '#', '#')
			case 'O':
				out[y] = append(out[y], '[', ']')
			case '.':
				out[y] = append(out[y], '.', '.')
			case '@':
				out[y] = append(out[y], '@', '.')
			}
		}
	}
	return out
}

func moveWide(p grid.Pos, g grid.Grid, d grid.Dir) bool {
	// find dest pos
	dest := p.Move1(d)
	if g.ValueAt(dest) == '#' {
		// if dest pos == '#', return, move nothing
		return false
	}
	// if move robot or move boxes sideway, use previous move logic
	if (d == grid.DirLeft || d == grid.DirRight) && (g.ValueAt(dest) == '[' || g.ValueAt(dest) == ']') {
		if !moveWide(dest, g, d) {
			return false
		}
		g[dest.Y][dest.X] = g.ValueAt(p)
		g[p.Y][p.X] = '.'
		return true
	}
	return false
}

func part2(input string) int {
	out := 0
	gStr, iStr, found := strings.Cut(strings.TrimSpace(input), "\n\n")
	if !found {
		log.Fatal("bad input, no empty line")
	}
	g := newGridWide(grid.NewGrid(gStr))
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("part2 start state:\n %v", g)) // __AUTO_GENERATED_PRINT_VAR_END__
	dirs := getInstructions(iStr)

	var rb grid.Pos
	for y, row := range g {
		for x, r := range row {
			if r == '@' {
				rb = grid.Pos{X: x, Y: y}
			}
		}
	}
	count := 0
	for _, dir := range dirs {
		if moveWide(rb, g, dir) {
			rb = rb.Move1(dir)
		}
		// __AUTO_GENERATED_PRINT_VAR_START__
		fmt.Println(fmt.Sprintf("part2 g:\n %v", g)) // __AUTO_GENERATED_PRINT_VAR_END__
		count++
		if count == 2 {
			break
		}
	}

	for y, row := range g {
		for x, r := range row {
			if r == '[' {
				out += x + 100*y
			}
		}
	}
	return out
}
