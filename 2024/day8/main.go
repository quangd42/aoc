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

func markAntinodes(a, b grid.Pos, g grid.Grid, marked map[grid.Pos]bool) {
	diff := grid.Dir{X: a.X - b.X, Y: a.Y - b.Y}
	a1 := grid.Pos{X: b.X - diff.X, Y: b.Y - diff.Y}
	a2 := grid.Pos{X: a.X + diff.X, Y: a.Y + diff.Y}

	if a1.IsInbound(g) {
		marked[a1] = true
	}
	if a2.IsInbound(g) {
		marked[a2] = true
	}
}

func part1(input string) int {
	g := grid.NewGrid(input)
	antennas := map[rune][]grid.Pos{}

	for y, row := range g {
		for x, r := range row {
			if r != '.' {
				antennas[r] = append(antennas[r], grid.Pos{X: x, Y: y})
			}
		}
	}

	marked := map[grid.Pos]bool{}
	for _, ps := range antennas {
		for i := 0; i < len(ps)-1; i++ {
			for j := i + 1; j < len(ps); j++ {
				markAntinodes(ps[i], ps[j], g, marked)
			}
		}
	}

	return len(marked)
}

func markAntinodes2(a, b grid.Pos, g grid.Grid, marked map[grid.Pos]bool) {
	marked[a] = true
	marked[b] = true
	diff := grid.Dir{X: a.X - b.X, Y: a.Y - b.Y}

	b1 := b.Move(diff, -1)
	i := 1
	for b1.IsInbound(g) {
		marked[b1] = true
		i++
		b1 = b.Move(diff, -i)
	}
	a1 := a.Move(diff, 1)
	j := 1
	for a1.IsInbound(g) {
		marked[a1] = true
		j++
		a1 = b.Move(diff, j)
	}
}

func part2(input string) int {
	g := grid.NewGrid(input)
	antennas := map[rune][]grid.Pos{}

	for y, row := range g {
		for x, r := range row {
			if r != '.' {
				antennas[r] = append(antennas[r], grid.Pos{X: x, Y: y})
			}
		}
	}

	marked := map[grid.Pos]bool{}
	for _, ps := range antennas {
		for i := 0; i < len(ps)-1; i++ {
			for j := i + 1; j < len(ps); j++ {
				markAntinodes2(ps[i], ps[j], g, marked)
			}
		}
	}

	return len(marked)
}
