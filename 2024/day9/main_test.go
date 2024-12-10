package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

var (
	example1 string = `2303313312141413142`
	example2 string = `6967128031985339734`
)

type test struct {
	name, input string
	want        int
}

func Test_part1(t *testing.T) {
	tests := []test{
		{
			"example 0",
			example,
			1928,
		},
		// {
		// 	"example 1",
		// 	example1,
		// 	294,
		// },
		// {
		// 	"example 2",
		// 	example2,
		// 	294,
		// },
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
			2858,
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
