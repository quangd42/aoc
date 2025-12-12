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

var dirs = directions(grid.FourDirs)

type directions []grid.Dir

func (ds directions) prevDir(i int) int {
	return (i - 1 + len(ds)) % len(ds)
}

func (ds directions) nextDir(i int) int {
	return (i + 1 + len(ds)) % len(ds)
}

func solve(g grid.Grid, start, end grid.Pos) int {
	return 0
}

func solveR(g grid.Grid, cur grid.Pos, curDirIdx int, score int) (int, bool) {
	// look forward, left, right
	// if found end, return score, true
	// if found wall, return 0, false
	// default: call recursive on next pos
	// if true, get min of out and score
	sc := make([]int, 0)
	dir := dirs[curDirIdx]
	next := cur.Move1(dir)
	switch g.ValueAt(next) {
	case '#':
	// continue to next dir
	case 'E':
		return score, true
	default:
		sf, found := solveR(g, next, curDirIdx, score+1)
		if found {
			sc = append(sc, sf)
		}
	}
	return 0, false
}

func part1(input string) int {
	out := 0
	g := grid.NewGrid(input)
	var startPos, endPos grid.Pos
	for y, row := range g {
		for x, r := range row {
			switch r {
			case 'S':
				startPos = grid.Pos{X: x, Y: y}
			case 'E':
				endPos = grid.Pos{X: x, Y: y}
			}
		}
	}
	out = solve(g, startPos, endPos)
	return out
}

func part2(input string) int {
	out := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
	}
	return out
}
