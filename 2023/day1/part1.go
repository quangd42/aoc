package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Solver struct{}

func (s Solver) Part1(f *os.File) int {
	scanner := bufio.NewScanner(f)
	var res int
	for scanner.Scan() {
		res += getCalNum1(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %w", err)
	}

	return res
}

func getCalNum1(line string) int {
	if len(line) == 0 {
		return 0
	}
	isFirst := true
	var first, last rune
	for _, c := range line {
		if unicode.IsDigit(c) {
			if isFirst {
				first = c
				isFirst = false
			}
			last = c
		}
	}
	firstInt, err := strconv.Atoi(string(first))
	if err != nil {
		log.Fatal(err)
	}
	lastInt, err := strconv.Atoi(string(last))
	if err != nil {
		log.Fatal(err)
	}
	return firstInt*10 + lastInt
}
