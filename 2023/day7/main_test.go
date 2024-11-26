package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

var example1 string = `32T3K 765
`

type test struct {
	name, input string
	want        int
}

func Test_calcType1(t *testing.T) {
	tests := []struct {
		input hand
		want  handType
	}{
		{hand{val: "32T3K"}, OnePair},
		{hand{val: "KK677"}, TwoPair},
		{hand{val: "T55J5"}, ThreeOfAKind},
	}
	for _, tt := range tests {
		t.Run(tt.input.val, func(t *testing.T) {
			tt.input.calcType1()
			got := tt.input.handType
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []test{
		{
			"example 1",
			example1,
			765,
		},
		{
			"full example",
			example,
			6440,
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

func Test_calcType2(t *testing.T) {
	tests := []struct {
		input hand
		want  handType
	}{
		{hand{val: "32T3K"}, OnePair},
		{hand{val: "KK677"}, TwoPair},
		{hand{val: "KJ677"}, ThreeOfAKind},
		{hand{val: "T55J5"}, FourOfAKind},
		{hand{val: "KTJJT"}, FourOfAKind},
		{hand{val: "KTJJ9"}, ThreeOfAKind},
		{hand{val: "KTJKT"}, FullHouse},
		{hand{val: "JJJ12"}, FourOfAKind},
		{hand{val: "JJJJ2"}, FiveOfAKind},
	}
	for _, tt := range tests {
		t.Run(tt.input.val, func(t *testing.T) {
			tt.input.calcType2()
			got := tt.input.handType
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []test{
		{
			"full example",
			example,
			5905,
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
