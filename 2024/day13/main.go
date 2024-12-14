package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/quangd42/aoc/parse"
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

var (
	costA = 3
	costB = 1
)

type (
	coord  struct{ x, y int }
	vector struct{ x, y int }
)

type machine struct {
	bA    vector
	bB    vector
	prize coord
}

// func (m machine) solve() (int, bool) {
// 	out := math.MaxInt
//
// 	for countA := m.prize.x / m.bA.x; countA >= 0; countA-- {
// 		if (m.prize.x-m.bA.x*countA)%m.bB.x == 0 && (m.prize.y-m.bA.y*countA)%m.bB.y == 0 {
// 			countB := (m.prize.x - m.bA.x*countA) / m.bB.x
// 			if countB <= 100 && countA <= 100 {
// 				out = min(out, countA*3+countB)
// 			}
// 		}
// 	}
// 	if out == math.MaxInt {
// 		return out, false
// 	}
// 	return out, true
// }

func solve(m machine) (int, bool) {
	out := math.MaxInt

	for countA := 100; countA >= 0; countA-- {
		for countB := 100; countB >= 0; countB-- {
			if countA*m.bA.x+countB*m.bB.x == m.prize.x && countA*m.bA.y+countB*m.bB.y == m.prize.y {
				out = min(out, countA*3+countB)
			}
		}
	}
	if out == math.MaxInt {
		return 0, false
	}
	return out, true
}

func parseMachine(s string) machine {
	re := regexp.MustCompile(`(\d+)`)
	out := machine{}
	matches := re.FindAllStringSubmatch(s, -1)
	if len(matches) != 6 {
		log.Fatal("malformed machine input")
	}
	out.bA = vector{x: parse.Int(matches[0][1]), y: parse.Int(matches[1][1])}
	out.bB = vector{x: parse.Int(matches[2][1]), y: parse.Int(matches[3][1])}
	out.prize = coord{x: parse.Int(matches[4][1]), y: parse.Int(matches[5][1])}
	return out
}

func parseInput(s string) []machine {
	msStr := strings.Split(strings.TrimSpace(s), "\n\n")
	out := make([]machine, len(msStr))
	for i, mStr := range msStr {
		out[i] = parseMachine(mStr)
	}
	return out
}

func part1(input string) int {
	out := 0
	mp := parseInput(input)
	for _, m := range mp {
		if cost, found := solve(m); found {
			out += cost
		}
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
