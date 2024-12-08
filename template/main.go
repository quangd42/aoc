package main

import (
	"bufio"
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
