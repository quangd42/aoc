package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	res := 0
	for scanner.Scan() {
		res += getPowerLine(scanner.Text())
		// println("res: ", res)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}

	return res
}

func getPowerLine(line string) int {
	// For each line, find groups
	_, colorsStr, found := strings.Cut(line, ":")
	if !found {
		log.Fatalf("invalid format: %s", line)
	}
	maxGreen := getMaxCountColor("green", colorsStr)
	maxBlue := getMaxCountColor("blue", colorsStr)
	maxRed := getMaxCountColor("red", colorsStr)
	// fmt.Printf("green: %d, blue: %d, red: %d\n", maxGreen, maxBlue, maxRed)

	fmt.Printf("power: %d\n", maxBlue*maxGreen*maxRed)
	return maxGreen * maxBlue * maxRed
}

func getMaxCountColor(color, line string) int {
	var exp *regexp.Regexp
	switch color {
	case "green":
		exp = regexp.MustCompile(`(?:(\d+)\s*green)`)
	case "blue":
		exp = regexp.MustCompile(`(?:(\d+)\s*blue)`)
	case "red":
		exp = regexp.MustCompile(`(?:(\d+)\s*red)`)
	}

	matches := exp.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return 0
	}
	// fmt.Printf("%s matches: %q\n", color, matches)

	maxCount := 0
	for _, match := range matches {
		if match[0] == "" {
			continue
		}
		count, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
