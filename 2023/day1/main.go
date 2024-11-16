package main

import _ "embed"

//go:embed input.txt
var input string

func main() {
	println("Part 1: ", part1(input))
	println("Part 2: ", part2(input))
}
