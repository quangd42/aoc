package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed example.txt
var example string

type test struct {
	name, input string
	want        int
}

func Test_part1(t *testing.T) {
	tests := []test{
		{
			"full example",
			example,
			35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %d; want %d\n", got, tt.want)
			}
		})
	}
}

func Test_mutate(t *testing.T) {
	paragraphs := strings.Split(strings.TrimSpace(example), "\n\n")
	tests := []struct {
		name  string
		input []string
		want  map[int]int
	}{
		{
			"seed-to-soil",
			paragraphs[:2],
			map[int]int{79: 81, 14: 14, 55: 57, 13: 13},
		},
		{
			"seed-to-fertilizer",
			paragraphs[:3],
			map[int]int{79: 81, 14: 53, 55: 57, 13: 52},
		},
		{
			"seed-to-water",
			paragraphs[:4],
			map[int]int{79: 81, 14: 49, 55: 53, 13: 41},
		},
		{
			"seed-to-light",
			paragraphs[:5],
			map[int]int{79: 74, 14: 42, 55: 46, 13: 34},
		},
		{
			"seed-to-temperature",
			paragraphs[:6],
			map[int]int{79: 78, 14: 42, 55: 82, 13: 34},
		},
		{
			"seed-to-humidity",
			paragraphs[:7],
			map[int]int{79: 78, 14: 43, 55: 82, 13: 35},
		},
		{
			"seed-to-location",
			paragraphs[:8],
			map[int]int{79: 82, 14: 43, 55: 86, 13: 35},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[int]int{79: 79, 14: 14, 55: 55, 13: 13}
			for i := 1; i < len(tt.input); i++ {
				ruleset := parseRuleset(tt.input[i])
				for k, v := range got {
					got[k] = applyRule(v, ruleset)
				}
			}
			for seed, valueWant := range tt.want {
				if valueGot := got[seed]; valueGot != valueWant {
					t.Errorf("wrong mapping, seed %d, got %d, want %d\n", seed, valueGot, valueWant)
				}
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []test{
		{
			"full example",
			example,
			46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %d; want %d\n", got, tt.want)
			}
		})
	}
}
