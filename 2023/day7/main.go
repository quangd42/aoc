package main

import (
	_ "embed"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type handType int

const (
	HighCard handType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type hand struct {
	val      string
	handType handType
	bid      int
}

func (h *hand) calcType1() {
	if h.val == "" {
		return
	}
	index := make(map[rune]int)
	for _, r := range h.val {
		if _, ok := index[r]; ok {
			index[r]++
		} else {
			index[r] = 1
		}
	}
	switch len(index) {
	case 5:
		h.handType = HighCard
	case 4:
		h.handType = OnePair
	case 3:
		for _, v := range index {
			if v == 3 {
				h.handType = ThreeOfAKind
				return
			}
		}
		h.handType = TwoPair
	case 2:
		for _, v := range index {
			if v == 4 {
				h.handType = FourOfAKind
				return
			}
		}
		h.handType = FullHouse
	case 1:
		h.handType = FiveOfAKind
	default:
		log.Fatalf("invalid hand value: %s", h.val)
	}
}

func (h *hand) calcType2() {
	if h.val == "" {
		return
	}
	index := make(map[rune]int)
	jokerCount := 0
	for _, r := range h.val {
		if r == rune('J') {
			jokerCount++
			continue
		}
		if _, ok := index[r]; ok {
			index[r]++
		} else {
			index[r] = 1
		}
	}
	maxCount := 0
	for _, v := range index {
		maxCount = max(maxCount, v)
	}
	switch maxCount + jokerCount {
	case 1:
		h.handType = HighCard
	case 2:
		if jokerCount == 0 && len(index) == 3 {
			h.handType = TwoPair
			return
		}
		h.handType = OnePair
	case 3:
		if jokerCount == 0 {
			if len(index) == 3 {
				h.handType = ThreeOfAKind
				return
			}
			h.handType = FullHouse
			return
		}
		if jokerCount == 1 && len(index) == 2 {
			h.handType = FullHouse
			return
		}
		h.handType = ThreeOfAKind
	case 4:
		h.handType = FourOfAKind
	case 5:
		h.handType = FiveOfAKind
	default:
		log.Fatalf("invalid hand value: %s", h.val)
	}
}

type game []hand

func (g *game) rankHands1() {
	order := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}
	slices.SortFunc[[]hand, hand]([]hand(*g), func(a, b hand) int {
		if n := int(a.handType) - int(b.handType); n != 0 {
			return n
		}
		for i := range a.val {
			ac, ok := order[rune(a.val[i])]
			if !ok {
				log.Fatalf("char in hand is not valid, hand: %s", a.val)
			}
			bc, ok := order[rune(b.val[i])]
			if !ok {
				log.Fatalf("char in hand is not valid, hand: %s", a.val)
			}
			if n := ac - bc; n != 0 {
				return n
			}
		}
		return 0
	})
}

func (g *game) rankHands2() {
	order := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}
	slices.SortFunc[[]hand, hand]([]hand(*g), func(a, b hand) int {
		if n := int(a.handType) - int(b.handType); n != 0 {
			return n
		}
		for i := range a.val {
			ac, ok := order[rune(a.val[i])]
			if !ok {
				log.Fatalf("char in hand is not valid, hand: %s", a.val)
			}
			bc, ok := order[rune(b.val[i])]
			if !ok {
				log.Fatalf("char in hand is not valid, hand: %s", a.val)
			}
			if n := ac - bc; n != 0 {
				return n
			}
		}
		return 0
	})
}

func (g *game) calcWinnings() int {
	hands := []hand(*g)
	out := 0
	for i := len(hands); i > 0; i-- {
		// fmt.Printf("hand %s, rank %d, handtype %#v, bid %d\n", hands[i-1].val, i, hands[i-1].handType, hands[i-1].bid)
		out += i * hands[i-1].bid
	}
	return out
}

func parseInput(input string) []hand {
	lines := strings.Split(strings.Trim(input, "\n "), "\n")
	hands := make([]hand, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}
		handVal, bidStr, found := strings.Cut(line, " ")
		if !found {
			log.Fatalf("invalid line: %s\n", line)
		}
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			log.Fatalf("invalid bid val: %s\n", bidStr)
		}
		hands[i] = hand{
			val: handVal,
			bid: bid,
		}
	}
	return hands
}

func part1(input string) int {
	hands := parseInput(input)
	for i := range hands {
		hands[i].calcType1()
	}
	g := game(hands)
	g.rankHands1()
	return g.calcWinnings()
}

func part2(input string) int {
	hands := parseInput(input)
	for i := range hands {
		hands[i].calcType2()
	}
	g := game(hands)
	g.rankHands2()
	return g.calcWinnings()
}

func main() {
	println("Part 1: ", part1(input))
	println("Part 2: ", part2(input))
}
