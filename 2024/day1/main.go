package main

import (
	"bufio"
	_ "embed"
	"log"
	"math"
	"slices"
	"strings"

	"github.com/quangd42/aoc/parse"
)

//go:embed input.txt
var input string

func main() {
	println("Part 1: ", part1(input))
	println("Part 2: ", part2(input))
}

type lists struct {
	left  []int
	right []int
}

func parser(input string) lists {
	scanner := bufio.NewScanner(strings.NewReader(input))
	left := []int{}
	right := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		ns := strings.Fields(line)
		if len(ns) < 2 {
			log.Fatal("not enough info", line)
		}

		left = append(left, parse.Int(ns[0]))
		right = append(right, parse.Int(ns[1]))
	}
	return lists{
		left:  left,
		right: right,
	}
}

func part1(input string) int {
	lists := parser(input)
	slices.Sort(lists.left)
	slices.Sort(lists.right)
	out := 0
	for i := range lists.left {
		out += int(math.Abs(float64(lists.left[i] - lists.right[i])))
	}
	return out
}

func part2(input string) int {
	lists := parser(input)
	cache := map[int]int{}
	out := 0
	for _, n := range lists.right {
		cache[n]++
	}
	for _, n := range lists.left {
		count := cache[n]
		out += n * count
	}
	return out
}
