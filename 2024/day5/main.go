package main

import (
	_ "embed"
	"log"
	"slices"
	"strings"

	"github.com/quangd42/aoc/parse"
)

//go:embed input.txt
var input string

func main() {
	println("Part 1: ", part1(input))
	println("Part 2: ", part2(input))
}

type rules map[string]map[string]bool

func parseRules(s string) map[string]map[string]bool {
	out := make(map[string]map[string]bool)
	for _, line := range strings.Split(s, "\n") {
		l, r, found := strings.Cut(line, "|")
		if !found {
			log.Fatalf("bad input %s\n", line)
		}
		if _, ok := out[l]; !ok {
			out[l] = map[string]bool{r: true}
			continue
		}
		out[l][r] = true
	}
	return out
}

func parseUpdates(s string) [][]string {
	out := [][]string{}
	for _, line := range strings.Split(s, "\n") {
		out = append(out, strings.Split(line, ","))
	}
	return out
}

func isSortedGetMid(u []string, r rules) (string, bool) {
	sorted := slices.IsSortedFunc(u, func(a, b string) int {
		set, ok := r[a]
		found := set[b]
		if ok && found {
			return -1
		}
		return 1
	})
	if !sorted {
		return "", false
	}
	uLength := len(u)
	if uLength%2 != 1 {
		log.Fatalf("length of updates is even: %#v", u)
	}
	return u[uLength/2], true
}

// This is the basis of the sorting condition
// cmp func was based on this brute force approach
func isOK(u []string, r rules) (string, bool) {
	uLength := len(u)
	for i := 0; i < uLength-1; i++ {
		set, ok := r[u[i]]
		if !ok {
			return "", false
		}
		for j := i + 1; j < uLength; j++ {
			if _, ok := set[u[j]]; !ok {
				return "", false
			}
		}
	}
	if uLength%2 != 1 {
		log.Fatalf("length of updates is even: %#v", u)
	}
	return u[uLength/2], true
}

func part1(input string) int {
	out := 0
	rulesRaw, updatesRaw, found := strings.Cut(input, "\n\n")
	if !found {
		log.Fatal("blank line not found in input")
	}
	rules := parseRules(strings.TrimSpace(rulesRaw))
	updates := parseUpdates(strings.TrimSpace(updatesRaw))

	for _, update := range updates {
		// mid, ok := isOK(update, rules)
		mid, ok := isSortedGetMid(update, rules)
		if !ok {
			continue
		}
		out += parse.Int(mid)
	}

	return out
}

// Instead of implementing the sorting algo myself
func sortGetMid(u []string, r rules) string {
	slices.SortFunc(u, func(a, b string) int {
		set, ok := r[a]
		found := set[b]
		if ok && found {
			return -1
		}
		return 1
	})
	uLength := len(u)
	if uLength%2 != 1 {
		log.Fatalf("length of updates is even: %#v", u)
	}
	return u[uLength/2]
}

func part2(input string) int {
	out := 0
	rulesRaw, updatesRaw, found := strings.Cut(input, "\n\n")
	if !found {
		log.Fatal("blank line not found in input")
	}
	rules := parseRules(strings.TrimSpace(rulesRaw))
	updates := parseUpdates(strings.TrimSpace(updatesRaw))

	for _, update := range updates {
		_, ok := isOK(update, rules)
		if ok {
			continue
		}
		out += parse.Int(sortGetMid(update, rules))
	}

	return out
}
