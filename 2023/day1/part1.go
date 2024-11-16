package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var res int
	for scanner.Scan() {
		res += getCalNum1(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
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
