package main

import (
	"bufio"
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	RedCap = 12 + iota
	GreenCap
	BlueCap
)

//go:embed input.txt
var input string

func part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	res := 0
	lineCount := 0
	for scanner.Scan() {
		lineCount += 1
		if validateLine(scanner.Text()) {
			res += lineCount
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}

	return res
}

func validateColors(colors string) bool {
	// println("input", colors)
	var green, blue, red int
	var err error
	// For each part of the game, extract number of cubes drawn
	blueExp := regexp.MustCompile(`(?:.*\s+(\d+)\s*blue)?`)
	greenExp := regexp.MustCompile(`(?:.*\s+(\d+)\s*green)?`)
	redExp := regexp.MustCompile(`(?:.*\s+(\d+)\s*red)?`)

	greenMatch := greenExp.FindStringSubmatch(colors)
	blueMatch := blueExp.FindStringSubmatch(colors)
	redMatch := redExp.FindStringSubmatch(colors)
	// fmt.Printf("greenMatch: %#v ", greenMatch)
	// fmt.Printf("blueMatch: %#v ", blueMatch)
	// fmt.Printf("redMatch: %#v\n", redMatch)
	if greenMatch[0] == "" {
		green = 0
	} else {
		green, err = strconv.Atoi(greenMatch[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	if blueMatch[0] == "" {
		blue = 0
	} else {
		blue, err = strconv.Atoi(blueMatch[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	if redMatch[0] == "" {
		red = 0
	} else {
		red, err = strconv.Atoi(redMatch[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	// fmt.Printf("result: green: %d, blue: %d, red: %d\n", green, blue, red)
	if green > GreenCap || blue > BlueCap || red > RedCap {
		return false
	}
	return true
}

func validateLine(line string) bool {
	// For each line, find groups
	_, colorsStr, found := strings.Cut(line, ":")
	if !found {
		log.Fatalf("invalid format: %s", line)
	}
	groups := strings.Split(colorsStr, ";")
	for _, group := range groups {
		if !validateColors(group) {
			return false
		}
	}
	return true
}
