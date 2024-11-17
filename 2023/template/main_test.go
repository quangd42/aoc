package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

var example1 string = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

var example2 string = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

type test struct {
	name, input string
	want        int
}

func Test_part1(t *testing.T) {
	tests := []test{
		{
			"first line",
			"1abc2",
			12,
		},
		{
			"full example 1",
			example1,
			142,
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
			"two1nine",
			29,
		},
		{
			"full example",
			example2,
			281,
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
