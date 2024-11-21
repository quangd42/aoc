package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

//go:embed input.txt
var input string

func main() {
	println("Part 1: ", part1(input))
	println("Part 2: ", part2(input))
}

type rule struct {
	to       int
	from     int
	valRange int
}

func parseRule(line string) (rule, error) {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return rule{}, fmt.Errorf("rule line malformed: %s\n", line)
	}
	partsInt := make([]int, 0)
	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			log.Fatal("rule line contains non int")
		}
		partsInt = append(partsInt, n)
	}
	return rule{
		to:       partsInt[0],
		from:     partsInt[1],
		valRange: partsInt[2],
	}, nil
}

func parseSeeds1(line string) map[int]int {
	cleaned, found := strings.CutPrefix(line, "seeds: ")
	if !found {
		log.Fatalf("seed line malformed: %s\n", line)
	}
	seedsStr := strings.Fields(cleaned)
	seeds := map[int]int{}
	for _, s := range seedsStr {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("seed line contains non int: %s", s)
		}
		seeds[n] = n
	}
	return seeds
}

func parseRuleset(paragraph string) []rule {
	out := make([]rule, 0)
	ruleLines := strings.Split(paragraph, "\n")
	for j := 1; j < len(ruleLines); j++ {
		r, err := parseRule(strings.TrimSpace(ruleLines[j]))
		if err != nil {
			fmt.Println("ruleLines: ", len(ruleLines))
			log.Fatal(paragraph)
		}
		out = append(out, r)
	}
	return out
}

func part1(input string) int {
	// Split input into paragraphs
	paragraphs := strings.Split(strings.TrimSpace(input), "\n\n")
	// First paragraph is the seed info
	seeds := parseSeeds1(paragraphs[0])
	// Parse the other paragraphs for rules
	for i := 1; i < len(paragraphs); i++ {
		ruleset := parseRuleset(paragraphs[i])
		for k, v := range seeds {
			seeds[k] = applyRule(v, ruleset)
		}
	}

	out := math.Inf(1)
	for _, loc := range seeds {
		// fmt.Printf("loc: %d, seed: %d\n", loc, seed)
		if float64(loc) < out {
			out = float64(loc)
		}
	}

	return int(out)
}

func applyRule(val int, ruleset []rule) int {
	for _, r := range ruleset {
		if val >= r.from && val <= r.from+r.valRange {
			return val + (r.to - r.from)
		}
	}
	return val
}

func applyAllMappings(val int, mappingList [][]rule) int {
	for _, ruleset := range mappingList {
		val = applyRule(val, ruleset)
	}

	return val
}

func parseSeeds2(line string, mappingList [][]rule) int {
	cleaned, found := strings.CutPrefix(line, "seeds: ")
	if !found {
		log.Fatalf("seed line malformed: %s\n", line)
	}
	seedsStr := strings.Fields(cleaned)
	if len(seedsStr)%2 != 0 {
		log.Fatalf("seed line malformed: %s\n", line)
	}

	startTime := time.Now()
	results := make(chan int, 1000)
	var wg sync.WaitGroup

	// Size of chunks to process
	chunkSize := 1_000_000

	// Process each seed range
	for i := 0; i < len(seedsStr); i += 2 {
		startVal, err := strconv.Atoi(seedsStr[i])
		if err != nil {
			log.Fatalf("seed line contains non int: %s", seedsStr[i])
		}
		valRange, err := strconv.Atoi(seedsStr[i+1])
		if err != nil {
			log.Fatalf("seed line contains non int: %s", seedsStr[i+1])
		}

		// Create work chunks
		for chunk := 0; chunk < valRange; chunk += chunkSize {
			endChunk := chunk + chunkSize
			if endChunk > valRange {
				endChunk = valRange
			}

			wg.Add(1)
			go func(start, from, to int) {
				defer wg.Done()
				localMin := math.MaxInt

				// Process chunk
				for j := from; j < to; j++ {
					loc := applyAllMappings(start+j, mappingList)
					localMin = min(localMin, loc)
				}

				results <- localMin
			}(startVal, chunk, endChunk)
		}

		fmt.Printf("Processing range starting at %d with length %d\n", startVal, valRange)
	}

	// Close results channel when all work is done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results as they come in
	minResult := math.MaxInt
	for result := range results {
		minResult = min(minResult, result)
	}

	fmt.Printf("time to process %.2f\n", time.Since(startTime).Seconds())
	return minResult
}

func part2(input string) int {
	// Split input into paragraphs
	paragraphs := strings.Split(strings.TrimSpace(input), "\n\n")
	// Create a list of mappings
	mappingList := make([][]rule, len(paragraphs)-1)
	for i := 1; i < len(paragraphs); i++ {
		ruleLines := strings.Split(paragraphs[i], "\n")
		rules := make([]rule, len(ruleLines)-1)
		for j := 1; j < len(ruleLines); j++ {
			rule, err := parseRule(strings.TrimSpace(ruleLines[j]))
			if err != nil {
				log.Fatalf("parseRule error: %v", err)
			}
			rules[j-1] = rule
		}
		mappingList[i-1] = rules
	}

	out := parseSeeds2(paragraphs[0], mappingList)

	return out
}
