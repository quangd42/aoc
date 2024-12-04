package main

import (
	_ "embed"
	"regexp"
	"strings"

	"github.com/quangd42/aoc/parse"
)

//go:embed input.txt
var input string

func main() {
	println("Part 1: ", part1(input))
	println("Part 2: ", part2(input))
}

func getMult(matches [][]string) int {
	out := 0
	for _, match := range matches {
		out += parse.Int(match[1]) * parse.Int(match[2])
	}
	return out
}

func part1(input string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	return getMult(matches)
}

func part2(input string) int {
	out := 0
	ins := strings.Split(input, "don't()")
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	out += getMult(re.FindAllStringSubmatch(ins[0], -1))
	for i := 1; i < len(ins); i++ {
		_, after, found := strings.Cut(ins[i], "do()")
		if found {
			matches := re.FindAllStringSubmatch(after, -1)
			out += getMult(matches)
		}
	}
	return out
}
