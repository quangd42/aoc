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
			"full example 1",
			example,
			2,
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
			"full example",
			example,
			4,
		},
		{
			"73 74 77 77 80 82 83 87",
			"73 74 77 77 80 82 83 87",
			0,
		},
		{
			"4 1 4 5 6 9 10",
			"4 1 4 5 6 9 10",
			1,
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
