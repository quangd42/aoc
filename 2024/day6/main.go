package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func main() {
	println("Part 1: ", part1(input))
	println("Part 2: ", part2(input))
}

type direction struct {
	x, y int
}

type position struct {
	x, y int
}

var (
	up    = direction{0, -1}
	right = direction{1, 0}
	down  = direction{0, 1}
	left  = direction{-1, 0}
	dirs  = []direction{up, right, down, left}
)

func getStartPos(grid []string) position {
	pos := position{}
	for y, row := range grid {
		for x := range row {
			if row[x] == '^' {
				pos.x = x
				pos.y = y
			}
		}
	}
	return pos
}

func isInbound(p position, grid []string) bool {
	maxX := len(grid[0]) - 1
	maxY := len(grid) - 1
	if 0 <= p.x && p.x <= maxX && 0 <= p.y && p.y <= maxY {
		return true
	}
	return false
}

func part1(input string) int {
	grid := strings.Split(strings.TrimSpace(input), "\n")
	guardPos := getStartPos(grid)

	curDirIdx := 0

	steps, exited := walkGrid(guardPos, curDirIdx, grid)
	if !exited {
		log.Fatal("stuck in a loop somewhere")
	}
	return steps
}

// WARNING: This sadly does not give correct answer even though it passes the example test
// I'm missing a corner case somewhere

type state struct {
	pos position
	dir direction
}

func NewState(p position, d direction) state {
	return state{
		pos: position{p.x, p.y},
		dir: direction{d.x, d.y},
	}
}

func walkGrid(cur position, curDirIdx int, grid []string) (int, bool) {
	walked := map[position]bool{}
	wps := map[state]bool{
		NewState(cur, dirs[curDirIdx]): true,
	}
	var next position
	for {
		next = position{
			cur.x + dirs[curDirIdx].x,
			cur.y + dirs[curDirIdx].y,
		}
		if !isInbound(next, grid) {
			return len(walked), true
		}

		if grid[next.y][next.x] == '#' {
			curDirIdx = (curDirIdx + 1 + len(dirs)) % len(dirs)
			continue
		}

		// check if this next step is already made
		wp := NewState(next, dirs[curDirIdx])
		if _, ok := wps[wp]; ok {
			return len(walked), false
		}
		wps[wp] = true

		// actually make the step
		cur.x = next.x
		cur.y = next.y
		walked[cur] = true
	}
}

func part2(input string) int {
	startTime := time.Now()
	grid := strings.Split(strings.TrimSpace(input), "\n")
	startPos := getStartPos(grid)
	cur := startPos

	curDirIdx := 0

	loopPos := map[position]bool{}
	wps := map[state]bool{}

	for {
		next := position{
			cur.x + dirs[curDirIdx].x,
			cur.y + dirs[curDirIdx].y,
		}
		// If next is out of bound, can's go there or place obstacle there
		if !isInbound(next, grid) {
			break
		}

		// If next is a block already, just change direction
		if grid[next.y][next.x] == '#' {
			curDirIdx = (curDirIdx + 1 + len(dirs)) % len(dirs)
			continue
		}

		// Check whether we will create a loop if we block next pos
		// can't place the blocker at the starting position
		// skip checking if it's already confirmed
		if next != startPos || loopPos[next] == false {
			gridClone := slices.Clone(grid)
			// --all this to modify a char in row...
			rowCloneRune := []rune(gridClone[next.y])
			rowCloneRune[next.x] = '#'
			gridClone[next.y] = string(rowCloneRune)
			// --
			_, exited := walkGrid(startPos, 0, gridClone)
			if !exited {
				loopPos[next] = true
			}
		}

		// Actually make the move
		cur.x = next.x
		cur.y = next.y
		wps[NewState(cur, dirs[curDirIdx])] = true
	}

	fmt.Println("part 2 took ", time.Since(startTime))

	// Can't place obstacle at the starting position
	if loopPos[startPos] {
		return len(loopPos) - 1
	}
	return len(loopPos)
}
