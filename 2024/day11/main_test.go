package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

var example1 string = `0`

type test struct {
	name, input string
	count       int
	want        int
}

func Test_part1(t *testing.T) {
	tests := []test{
		// {
		// 	"6",
		// 	example,
		// 	6,
		// 	22,
		// },
		{
			"0",
			example1,
			75,
			22,
		},
		// {
		// 	"25",
		// 	example,
		// 	25,
		// 	55312,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.count); got != tt.want {
				t.Errorf("part1() = %d; want %d\n", got, tt.want)
			}
		})
	}
}
