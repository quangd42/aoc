package main

import (
	"bufio"
	_ "embed"
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

func parseLine(input string) []int {
	nStr := strings.Fields(input)
	out := make([]int, len(nStr))
	for i, n := range nStr {
		out[i] = parse.Int(n)
	}
	return out
}

func isSafe1(l []int) bool {
	var inc bool
	if l[1]-l[0] == 0 {
		return false
	} else if l[1]-l[0] < 0 {
		inc = false
	} else {
		inc = true
	}
	for i := 0; i < len(l)-1; i++ {
		dif := l[i+1] - l[i]
		if abs(dif) < 1 || abs(dif) > 3 || (dif < 0 && inc) || (dif > 0 && !inc) {
			return false
		}
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	out := 0
	for scanner.Scan() {
		if isSafe1(parseLine(scanner.Text())) {
			out++
		}
	}
	return out
}

func isSafe2(l []int) bool {
	o := slices.Clone(l)
	t := struct {
		tol int
		inc int
	}{
		tol: 1,
	}
	for i := 0; i < len(o)-1; i++ {
		dif := o[i+1] - o[i]
		if t.inc == 0 {
			t.inc = dif
		}
		if abs(dif) < 1 || abs(dif) > 3 || (dif < 0 && t.inc > 0) || (dif > 0 && t.inc < 0) {
			if t.tol < 1 {
				return false
			}
			t.tol--
			o[i+1] = l[i]
		}
	}
	return true
}

func part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	out := 0
	for scanner.Scan() {
		f := parseLine(scanner.Text())
		r := slices.Clone(f)
		slices.Reverse(r)
		if isSafe2(f) || isSafe2(r) {
			out++
		}
	}
	return out
}
