package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/quangd42/aoc/parse"
)

//go:embed input.txt
var input string

func main() {
	println("Part 1: ", part1(input))
	println("Part 1R: ", part1R(input))
	println("Part 2: ", part2(input))
}

type equation struct {
	test int
	nums []int
}

func parseInput(s string) []equation {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	out := make([]equation, len(lines))
	for i, l := range lines {
		var e equation
		bef, aft, found := strings.Cut(l, ":")
		if !found {
			log.Fatalf("bad line %s\n", l)
		}
		e.test = parse.Int(bef)

		nums := []int{}
		for _, n := range strings.Fields(strings.TrimSpace(aft)) {
			nums = append(nums, parse.Int(n))
		}
		e.nums = nums
		out[i] = e
	}
	return out
}

func isPossible(e equation, curIdx, curSum int) bool {
	if curIdx == len(e.nums) {
		if curSum != e.test {
			return false
		}
		return true
	}

	return isPossible(e, curIdx+1, curSum+e.nums[curIdx]) || isPossible(e, curIdx+1, curSum*e.nums[curIdx])
}

func isPossibleR(e equation, idx, val int) bool {
	if idx == 0 {
		if val != e.nums[0] {
			return false
		}
		return true
	}
	if val%e.nums[idx] != 0 {
		return isPossibleR(e, idx-1, val-e.nums[idx])
	}
	return isPossibleR(e, idx-1, val-e.nums[idx]) || isPossibleR(e, idx-1, int(val/e.nums[idx]))
}

func part1(input string) int {
	startTime := time.Now()
	out := 0
	es := parseInput(input)
	for _, e := range es {
		if isPossible(e, 1, e.nums[0]) {
			out += e.test
		}
	}
	fmt.Printf("Part 1 took %v\n", time.Since(startTime))
	return out
}

func part1R(input string) int {
	startTime := time.Now()
	out := 0
	es := parseInput(input)
	for _, e := range es {
		if isPossibleR(e, len(e.nums)-1, e.test) {
			out += e.test
		}
	}
	fmt.Printf("Part 1R took %v\n", time.Since(startTime))
	return out
}

func concatInt(a, b int) int64 {
	as := strconv.Itoa(a)
	bs := strconv.Itoa(b)
	cc := fmt.Sprintf("%s%s", as, bs)
	out, err := strconv.ParseInt(cc, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func isPossible2(e equation, curIdx, curSum int) bool {
	if curIdx == len(e.nums) {
		if curSum != e.test {
			return false
		}
		return true
	}

	return (isPossible2(e, curIdx+1, curSum+e.nums[curIdx]) ||
		isPossible2(e, curIdx+1, curSum*e.nums[curIdx]) ||
		isPossible2(e, curIdx+1, int(concatInt(curSum, e.nums[curIdx]))))
}

func part2(input string) int {
	startTime := time.Now()
	out := 0
	es := parseInput(input)
	for _, e := range es {
		if isPossible2(e, 1, e.nums[0]) {
			out += e.test
		}
	}
	fmt.Printf("Part 2 took %v\n", time.Since(startTime))
	return out
}
