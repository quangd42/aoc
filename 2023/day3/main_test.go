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
	tt := test{
		"example",
		example,
		4361,
	}
	if got := part1(tt.input); got != tt.want {
		t.Errorf("part1() = %d; want %d\n", got, tt.want)
	}
}

func Test_part2(t *testing.T) {
	tt := test{
		"example",
		example,
		467835,
	}
	if got := part2(tt.input); got != tt.want {
		t.Errorf("part2() = %d; want %d\n", got, tt.want)
	}
}
