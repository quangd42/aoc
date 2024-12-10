package main

import (
	_ "embed"
	"fmt"
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

var directions = grid.FourDirs

func score(g grid.Grid, cur grid.Pos, peaks map[grid.Pos]bool) {
	v := int(g.ValueAt(cur) - '0')
	if v == 9 {
		peaks[cur] = true
		return
	}
	for _, dir := range directions {
		np := cur.Move(dir, 1)
		if np.IsInbound(g) && int(g.ValueAt(np)-'0') == v+1 {
			score(g, np, peaks)
		}
	}
}

func part1(input string) int {
	out := 0
	g := grid.NewGrid(input)
	for y, row := range g {
		for x, r := range row {
			if r == '0' {
				peaks := map[grid.Pos]bool{}
				score(g, grid.Pos{X: x, Y: y}, peaks)
				out += len(peaks)
			}
		}
	}
	return out
}

func rate(g grid.Grid, cur grid.Pos) int {
	out := 0
	v := int(g.ValueAt(cur) - '0')
	if v == 9 {
		return 1
	}

	for _, np := range grid.Around4(cur) {
		if np.IsInbound(g) && int(g.ValueAt(np)-'0') == v+1 {
			out += rate(g, np)
		}
	}

	return out
}

func part2(input string) int {
	out := 0
	gr := grid.NewGrid(input)

	for y, row := range gr {
		for x, r := range row {
			if r == '0' {
				out += rate(gr, grid.Pos{X: x, Y: y})
			}
		}
	}
	return out
}
