package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

var firstExample string = `Time:      7
Distance:  9
`

type test struct {
	name, input string
	want        int
}

func Test_part1(t *testing.T) {
	tests := []test{
		{
			"first example",
			firstExample,
			4,
		},
		{
			"full example",
			example,
			288,
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
			firstExample,
			4,
		},
		{
			"full example",
			example,
			71503,
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
