package main

import (
	_ "embed"
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
			"first line",
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n",
			1,
		},
		{
			"full example",
			example,
			8,
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

func Test_part2(t *testing.T) {
	tests := []test{
		{
			"first line",
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n",
			48,
		},
		{
			"full example",
			example,
			2286,
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
