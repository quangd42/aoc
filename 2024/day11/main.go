package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/quangd42/aoc/parse"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimSpace(input)
	st1 := time.Now()
	fmt.Printf("Part 1: %d     %v\n", part1(input, 25), time.Since(st1))
	st2 := time.Now()
	fmt.Printf("Part 2: %d     %v\n", part1(input, 75), time.Since(st2))
}

func parseInput(s string) []int {
	nums := strings.Fields(strings.TrimSpace(s))
	out := make([]int, len(nums))
	for i, n := range nums {
		out[i] = parse.Int(n)
	}
	return out
}

func morph(i int) []int {
	s := strconv.Itoa(i)
	switch {
	case i == 0:
		return []int{1}
	case len(s)%2 == 0:
		mid := len(s) / 2
		return []int{parse.Int(s[:mid]), parse.Int(s[mid:])}
	default:
		return []int{i * 2024}
	}
}

func part1(input string, count int) int {
	rocks := parseInput(input)
	for range count {
		temp := []int{}
		for _, r := range rocks {
			temp = append(temp, morph(r)...)
		}
		rocks = temp
	}
	return len(rocks)
}
