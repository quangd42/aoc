package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
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

type position struct {
	x, y int
}

func (p position) isInbound(grid []string) bool {
	maxX := len(grid[0]) - 1
	maxY := len(grid) - 1
	if 0 <= p.x && p.x <= maxX && 0 <= p.y && p.y <= maxY {
		return true
	}
	return false
}

func markAntinodes(a, b position, grid []string, marked map[position]bool) {
	diff := position{a.x - b.x, a.y - b.y}
	a1 := position{b.x - diff.x, b.y - diff.y}
	a2 := position{a.x + diff.x, a.y + diff.y}

	if a1.isInbound(grid) {
		marked[a1] = true
	}
	if a2.isInbound(grid) {
		marked[a2] = true
	}
}

func part1(input string) int {
	grid := strings.Split(input, "\n")
	antennas := map[rune][]position{}

	for y, row := range grid {
		for x, r := range row {
			if r != '.' {
				antennas[r] = append(antennas[r], position{x, y})
			}
		}
	}

	marked := map[position]bool{}
	for _, ps := range antennas {
		for i := 0; i < len(ps)-1; i++ {
			for j := i + 1; j < len(ps); j++ {
				markAntinodes(ps[i], ps[j], grid, marked)
			}
		}
	}

	return len(marked)
}

func markAntinodes2(a, b position, grid []string, marked map[position]bool) {
	diff := position{a.x - b.x, a.y - b.y}
	a1 := position{b.x - diff.x, b.y - diff.y}
	a2 := position{a.x + diff.x, a.y + diff.y}
	marked[a] = true
	marked[b] = true

	i := 1
	for a1.isInbound(grid) {
		marked[a1] = true
		i++
		a1 = position{b.x - diff.x*i, b.y - diff.y*i}
	}
	j := 1
	for a2.isInbound(grid) {
		marked[a2] = true
		j++
		a2 = position{a.x + diff.x*j, a.y + diff.y*j}
	}
}

func part2(input string) int {
	grid := strings.Split(input, "\n")
	antennas := map[rune][]position{}

	for y, row := range grid {
		for x, r := range row {
			if r != '.' {
				antennas[r] = append(antennas[r], position{x, y})
			}
		}
	}

	marked := map[position]bool{}
	for _, ps := range antennas {
		for i := 0; i < len(ps)-1; i++ {
			for j := i + 1; j < len(ps); j++ {
				markAntinodes2(ps[i], ps[j], grid, marked)
			}
		}
	}

	return len(marked)
}
