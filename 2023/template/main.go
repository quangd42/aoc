package main

import (
	"bufio"
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	println("Part 1: ", part1(input))
	println("Part 2: ", part2(input))
}

func part1(input string) int {
	out := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
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
