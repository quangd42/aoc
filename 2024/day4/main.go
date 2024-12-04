package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	println("Part 1: ", part1(input))
	println("Part 2: ", part2(input))
}

type coord struct {
	x, y int
}

type direction struct {
	x, y int
}

func isWord(cur int, target []rune, c coord, input []string, d direction) int {
	maxX := len(input[0]) - 1
	maxY := len(input) - 1

	if d.y+c.y < 0 || d.y+c.y > maxY || d.x+c.x < 0 || d.x+c.x > maxX {
		return 0
	}
	if rune(input[d.y+c.y][d.x+c.x]) == target[cur] {
		if cur == len(target)-1 {
			return 1
		}
		return isWord(cur+1, target, coord{c.x + d.x, c.y + d.y}, input, d)

	}
	return 0
}

func part1(input string) int {
	out := 0
	chars := []rune{'X', 'M', 'A', 'S'}
	rows := strings.Split(strings.TrimSpace(input), "\n")

	dirs := []direction{}
	for i := range 3 {
		for j := range 3 {
			if i == 1 && j == 1 {
				continue
			}
			dirs = append(dirs, direction{
				x: i - 1,
				y: j - 1,
			})
		}
	}

	for y, row := range rows {
		for x := range row {
			if rune(row[x]) == 'X' {
				for _, dir := range dirs {
					out += isWord(1, chars, coord{x, y}, rows, dir)
				}
			}
		}
	}

	return out
}

func isX(c coord, input []string) int {
	if c.x == 0 || c.x == len(input[0])-1 || c.y == 0 || c.y == len(input)-1 {
		return 0
	}
	dirs := []direction{
		{-1, -1}, {1, -1},
	}
	charMap := map[byte]byte{
		'M': 'S',
		'S': 'M',
	}
	for _, d := range dirs {
		char, ok := charMap[input[c.y+d.y][c.x+d.x]]
		if !ok || char != input[c.y-d.y][c.x-d.x] {
			return 0
		}
	}
	return 1
}

func part2(input string) int {
	out := 0
	rows := strings.Split(strings.TrimSpace(input), "\n")

	for y, row := range rows {
		for x := range row {
			if rune(row[x]) == 'A' {
				out += isX(coord{x, y}, rows)
			}
		}
	}

	return out
}
