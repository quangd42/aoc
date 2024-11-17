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
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			8,
		},
		{
			"full example 1",
			example,
			13,
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
			30,
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
