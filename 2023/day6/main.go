package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	println("Part 1: ", part1(input))
	println("Part 2: ", part2(input))
}

type game struct {
	time     int
	distance int
}

func parseInput1(input string) []game {
	lines := strings.Split(input, "\n")
	if len(lines) < 2 {
		fmt.Printf("lines: %#v", lines)
		log.Fatal("malformed input: incorrect number of lines")
	}
	_, timesStr, found := strings.Cut(lines[0], ":")
	if !found {
		log.Fatal("malformed input: ':' not found")
	}
	_, distancesStr, found := strings.Cut(lines[1], ":")
	if !found {
		log.Fatal("malformed input: ':' not found")
	}
	times := strings.Fields(timesStr)
	distances := strings.Fields(distancesStr)
	games := make([]game, len(times))
	for i := range times {
		t, err := strconv.Atoi(times[i])
		if err != nil {
			log.Fatal("time is not int")
		}
		d, err := strconv.Atoi(distances[i])
		if err != nil {
			log.Fatal("distance is not int")
		}
		games[i] = game{
			time:     t,
			distance: d,
		}
	}

	return games
}

func calcGameOptions(g game) int {
	var start int
	var end int
	var dist int
	for i := g.time; i >= 0; i-- {
		dist = i * (g.time - i)
		if dist > g.distance {
			end = i
			break
		}
	}
	if end == 0 {
		return 0
	}
	for i := range g.time {
		dist = i * (g.time - i)
		if dist > g.distance {
			start = i
			break
		}
	}
	return end - start + 1
}

func part1(input string) int {
	games := parseInput1(input)
	out := 1
	for _, g := range games {
		out *= calcGameOptions(g)
	}
	return out
}

func parseInput2(input string) game {
	lines := strings.Split(input, "\n")
	if len(lines) < 2 {
		fmt.Printf("lines: %#v", lines)
		log.Fatal("malformed input: incorrect number of lines")
	}
	_, timeStr, found := strings.Cut(lines[0], ":")
	if !found {
		log.Fatal("malformed input: ':' not found")
	}
	_, distanceStr, found := strings.Cut(lines[1], ":")
	if !found {
		log.Fatal("malformed input: ':' not found")
	}
	time, err := strconv.Atoi(strings.Join(strings.Fields(timeStr), ""))
	if err != nil {
		log.Fatal("time is not int")
	}
	distance, err := strconv.Atoi(strings.Join(strings.Fields(distanceStr), ""))
	if err != nil {
		log.Fatal("time is not int")
	}
	return game{
		time:     time,
		distance: distance,
	}
}

func part2(input string) int {
	g := parseInput2(input)
	return calcGameOptions(g)
}
