package main

import (
	_ "embed"
	"log"
	"strings"
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

func (p position) canReach(o position, m []string) bool {
	if p.x == o.x && p.y == o.y {
		return true
	}
	cur := position{p.x, p.y}
	switch {
	case cur.x == o.x && o.y < cur.y:
		for o.y < cur.y {
			cur.y--
			if m[cur.y][cur.x] == '#' {
				return false
			}
		}
		return true
	case cur.x == o.x && o.y > cur.y:
		for o.y > cur.y {
			cur.y++
			if m[cur.y][cur.x] == '#' {
				return false
			}
		}
		return true
	case cur.y == o.y && o.x > cur.x:
		for o.x > cur.x {
			cur.x++
			if m[cur.y][cur.x] == '#' {
				return false
			}
		}
		return true
	case cur.y == o.y && o.x < cur.x:
		for o.x < cur.x {
			cur.x--
			if m[cur.y][cur.x] == '#' {
				return false
			}
		}
		return true
	}
	log.Fatal("something went wrong with canReach")
	return false
}

var (
	up    = direction{0, -1}
	right = direction{1, 0}
	down  = direction{0, 1}
	left  = direction{-1, 0}
	dirs  = []direction{up, right, down, left}
)

func getStartPos(m []string) position {
	pos := position{}
	for y, row := range m {
		for x := range row {
			if row[x] == '^' {
				pos.x = x
				pos.y = y
			}
		}
	}
	return pos
}

func isInbound(p position, m []string) bool {
	maxX := len(m[0]) - 1
	maxY := len(m) - 1
	if 0 <= p.x && p.x <= maxX && 0 <= p.y && p.y <= maxY {
		return true
	}
	return false
}

func part1(input string) int {
	gameMap := strings.Split(strings.TrimSpace(input), "\n")
	guardPos := getStartPos(gameMap)

	curDirIdx := 0
	walked := map[position]bool{}

	for {
		next := position{
			guardPos.x + dirs[curDirIdx].x,
			guardPos.y + dirs[curDirIdx].y,
		}
		if !isInbound(next, gameMap) {
			break
		}

		if gameMap[next.y][next.x] == '#' {
			curDirIdx = (curDirIdx + 1 + len(dirs)) % len(dirs)
			continue
		}
		guardPos.x = next.x
		guardPos.y = next.y
		walked[guardPos] = true
	}
	return len(walked)
}

// WARNING: This sadly does not give correct answer even though it passes the example test

func getNextDir(idx int) direction {
	idx = (idx + 1 + len(dirs)) % len(dirs)
	return dirs[idx]
}

func part2(input string) int {
	gameMap := strings.Split(strings.TrimSpace(input), "\n")
	guardPos := getStartPos(gameMap)

	curDirIdx := 0
	var nextDir direction
	lineWalked := map[direction][]position{
		up:    {guardPos},
		right: {},
		down:  {},
		left:  {},
	}

	placedPos := map[position]bool{}
	for {
		next := position{
			guardPos.x + dirs[curDirIdx].x,
			guardPos.y + dirs[curDirIdx].y,
		}
		if !isInbound(next, gameMap) {
			break
		}

		nextDir = getNextDir(curDirIdx)
		walkedStartPoints := lineWalked[nextDir]

		for _, wsp := range walkedStartPoints {
			switch nextDir {
			case up:
				if guardPos.x == wsp.x && (guardPos.y <= wsp.y || (guardPos.y > wsp.y && guardPos.canReach(wsp, gameMap))) {
					placedPos[next] = true
				}
			case down:
				if guardPos.x == wsp.x && (guardPos.y >= wsp.y || (guardPos.y <= wsp.y && guardPos.canReach(wsp, gameMap))) {
					placedPos[next] = true
				}
			case left:
				if guardPos.y == wsp.y && (guardPos.x <= wsp.x || (guardPos.x > wsp.x && guardPos.canReach(wsp, gameMap))) {
					placedPos[next] = true
				}
			case right:
				if guardPos.y == wsp.y && (guardPos.x >= wsp.x || (guardPos.x < wsp.x && guardPos.canReach(wsp, gameMap))) {
					placedPos[next] = true
				}
			}
		}

		if gameMap[next.y][next.x] == '#' {
			curDirIdx = (curDirIdx + 1 + len(dirs)) % len(dirs)
			lineWalked[dirs[curDirIdx]] = append(lineWalked[dirs[curDirIdx]], position{
				guardPos.x,
				guardPos.y,
			})
			continue
		}

		guardPos.x = next.x
		guardPos.y = next.y
	}

	return len(placedPos)
}
