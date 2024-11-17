package main

import (
	"bufio"
	_ "embed"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	println("Part 1: ", part1(input))
	println("Part 2: ", part2(input))
}

func part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	res := 0
	for scanner.Scan() {
		count := countWinning(scanner.Text())
		if count > 0 {
			res += int(math.Pow(2, float64(count-1)))
		}
	}
	return res
}

func part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	games := make(map[int]int)
	out := 0
	id := 1
	for scanner.Scan() {
		currentCardCount := games[id]
		// At least there is one card being processed
		currentCardCount++
		// Count winning
		winning := countWinning(scanner.Text())
		for i := range winning {
			games[id+1+i] += currentCardCount
		}
		out += currentCardCount
		// fmt.Printf("Card %d: count %d, winning %d\n", id, currentCardCount, winning)
		id++
	}
	// Go through the same process for the rest of the cards
	return out
}

func countWinning(s string) int {
	_, cleaned, ok := strings.Cut(s, ":")
	if !ok {
		log.Fatal("err cutting string: `:` not found")
	}
	winNumsStr, myNumsStr, ok := strings.Cut(cleaned, "|")
	if !ok {
		log.Fatal("err cutting string: `|` not found")
	}
	winNums := parseNums(winNumsStr)
	myNums := parseNums(myNumsStr)
	count := 0
	for n := range myNums {
		if _, found := winNums[n]; found {
			count++
		}
	}
	return count
}

func parseNums(s string) map[int]bool {
	numsStr := strings.Fields(s)
	out := make(map[int]bool, len(numsStr))
	for _, ns := range numsStr {
		n, err := strconv.Atoi(ns)
		if err != nil {
			log.Fatal(err)
		}
		out[n] = true
	}
	return out
}
