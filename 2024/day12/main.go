package main

import (
	"bufio"
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

type region struct {
	val       map[grid.Pos]bool
	perimeter int
}

// TODO:
func (r region) cost() int { return 0 }

type finder struct {
	v grid.Grid
}

func (f finder) findRegion(g grid.Grid, p grid.Pos) (region, bool) {
	// if this pos is already marked in the finder, return false

	// start recursive find and return region
	return region{}, false
}

func part1(input string) int {
	out := 0
	g := grid.NewGrid(input)
	f := finder{grid.NewGrid(input)}
	for y, row := range g {
		for x := range row {
			re, found := f.findRegion(g, grid.Pos{X: x, Y: y})
			if found {
				out += re.cost()
			}
		}
	}
	return out
}

func part2(input string) int {
	out := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
	}
	return out
}
