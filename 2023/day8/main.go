package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type node struct {
	name, left, right string
}

func parseNodes(s string) map[string]node {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	out := make(map[string]node)
	var name, left, right, dests string
	var found bool
	for _, line := range lines {
		name, dests, found = strings.Cut(line, "=")
		if !found {
			fmt.Println("malformed line", line)
		}
		name = strings.TrimSpace(name)
		left, right, found = strings.Cut(dests, ",")
		if !found {
			fmt.Println("malformed line", line)
		}
		out[name] = node{
			name:  name,
			left:  strings.Trim(left, "( "),
			right: strings.Trim(right, ") "),
		}
	}
	return out
}

func main() {
	input = strings.TrimSpace(input)
	st1 := time.Now()
	fmt.Printf("Part 1: %d     %v\n", part1(input), time.Since(st1))
	st2 := time.Now()
	fmt.Printf("Part 2: %d     %v\n", part2(input), time.Since(st2))
}

func part1(input string) int {
	out := 0
	ins, nodeStr, found := strings.Cut(input, "\n\n")
	if !found {
		log.Fatal("bad input")
	}
	nodes := parseNodes(nodeStr)
	cur := nodes["AAA"]
	for cur.name != "ZZZ" {
		for _, r := range ins {
			if cur.name == "ZZZ" {
				break
			}
			switch r {
			case 'L':
				cur = nodes[cur.left]
				out++
			case 'R':
				cur = nodes[cur.right]
				out++
			default:
				log.Fatal("invalid instruction")
			}
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
