package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/quangd42/aoc/grid"
	"github.com/quangd42/aoc/parse"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimSpace(input)
	st1 := time.Now()
	fmt.Printf("Part 1: %d     %v\n", part1(input, floor{101, 103}), time.Since(st1))
	st2 := time.Now()
	fmt.Printf("Part 2: %d     %v\n", part2(input), time.Since(st2))
}

type robot struct {
	grid.Pos
	vel grid.Dir
}

// returns the number of quadrant after move
func (rb *robot) move(f floor, t int) int {
	rb.X = (rb.X + rb.vel.X*t) % f.x
	rb.X = (rb.X + f.x) % f.x
	rb.Y = (rb.Y + rb.vel.Y*t) % f.y
	rb.Y = (rb.Y + f.y) % f.y

	midX := (f.x - 1) / 2
	midY := (f.y - 1) / 2

	switch {
	case rb.X > midX && rb.Y < midY:
		return 1
	case rb.X > midX && rb.Y > midY:
		return 2
	case rb.X < midX && rb.Y > midY:
		return 3
	case rb.X < midX && rb.Y < midY:
		return 4
	default:
		return 0
	}
}

type floor struct {
	x, y int
}

func parseInput(s string) []robot {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	out := make([]robot, len(lines))
	for i, l := range lines {
		p, v, f := strings.Cut(l, " ")
		if !f {
			log.Fatalf("malformed input: %s\n", l)
		}
		rb := robot{}
		px, py, f := strings.Cut(strings.TrimLeft(p, "p="), ",")
		if !f {
			log.Fatalf("malformed input: %s\n", l)
		}
		rb.Pos = grid.Pos{X: parse.Int(px), Y: parse.Int(py)}
		vx, vy, f := strings.Cut(strings.TrimLeft(v, "v="), ",")
		if !f {
			log.Fatalf("malformed input: %s\n", l)
		}
		rb.vel = grid.Dir{X: parse.Int(vx), Y: parse.Int(vy)}
		out[i] = rb
	}
	return out
}

func part1(input string, fl floor) int {
	out := 1
	rbs := parseInput(input)
	tracker := map[int]int{}
	for _, rb := range rbs {
		qr := rb.move(fl, 100)
		tracker[qr] = tracker[qr] + 1
	}
	for i := range 4 {
		out *= tracker[i+1]
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
