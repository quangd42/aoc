package main

import (
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
	start, end        bool
	steps             int
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
		if len(name) != 3 {
			log.Fatal("node name is too short:", name)
		}
		left, right, found = strings.Cut(dests, ",")
		if !found {
			fmt.Println("malformed line", line)
		}
		out[name] = node{
			name:  name,
			left:  strings.Trim(left, "( "),
			right: strings.Trim(right, ") "),
			start: name[2] == 'A',
			end:   name[2] == 'Z',
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
			case 'R':
				cur = nodes[cur.right]
			default:
				log.Fatal("invalid instruction")
			}
			out++
		}
	}
	return out
}

func moveAllNodes(ns []node, nodeMap map[string]node, r rune) {
	for i := range ns {
		switch r {
		case 'L':
			ns[i] = nodeMap[ns[i].left]
		case 'R':
			ns[i] = nodeMap[ns[i].right]
		default:
			log.Fatal("invalid instruction")
		}
	}
}

func part2(input string) int {
	out := 0
	ins, nodeStr, found := strings.Cut(input, "\n\n")
	if !found {
		log.Fatal("bad input")
	}
	nodeMap := parseNodes(nodeStr)

	ns := []node{}
	for _, v := range nodeMap {
		if v.start {
			ns = append(ns, v)
		}
	}

	for _, n := range ns {
		count := 0
		for !n.end {
			for _, r := range strings.TrimSpace(ins) {
				switch r {
				case 'L':
					n = nodeMap[n.left]
				case 'R':
					n = nodeMap[n.right]
				default:
					log.Fatal("invalid instruction")
				}
				count++
			}
		}
		fmt.Println(count)
	}

	// turns out the input indeed allows each route to make a loop,
	// it takes the same amount of steps for each route to get back to Z
	// (found this out from the internet, thought about lcm but could not verify the condition for it)
	//
	// so count the steps for each route and find the Least Common Multiple online

	return out
}
