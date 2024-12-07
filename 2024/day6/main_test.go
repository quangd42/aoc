package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

//go:embed example1.txt
var example1 string

type test struct {
	name, input string
	want        int
}

func Test_part1(t *testing.T) {
	tests := []test{
		{
			"full example 1",
			example,
			41,
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
			6,
		},
		{
			"smaller example",
			example1,
			3,
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
