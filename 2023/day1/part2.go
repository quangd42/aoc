package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var res int
	numbers := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	trie := makeTrie(numbers)
	for scanner.Scan() {
		res += getCalNum2(trie, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}

	return res
}

func getCalNum2(trie *Node, line string) int {
	if len(line) == 0 {
		return 0
	}
	isFirst := true
	var first, last rune
	var firstInt, lastInt int
	var err error
	for i, c := range line {
		if unicode.IsNumber(c) {
			if isFirst {
				first = c
				isFirst = false
				firstInt, err = strconv.Atoi(string(first))
				if err != nil {
					log.Fatal(err)
				}
			}
			last = c
			lastInt, err = strconv.Atoi(string(last))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			num, ok := scanForNum(trie, line[i:])
			if ok {
				if isFirst {
					firstInt = num
					isFirst = false
				}
				lastInt = num
			}
		}
	}
	return firstInt*10 + lastInt
}

type Node struct {
	children map[rune]*Node
	end      int
}

func NewNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
		end:      0,
	}
}

func addToTrie(n *Node, word string, num int) {
	for _, c := range word {
		_, ok := n.children[c]
		if !ok {
			n.children[c] = NewNode()
		}
		n = n.children[c]
	}
	n.end = num
}

func makeTrie(nums map[string]int) *Node {
	head := NewNode()
	for word, num := range nums {
		addToTrie(head, word, num)
	}
	return head
}

func scanForNum(n *Node, input string) (int, bool) {
	for _, c := range input {
		child, ok := n.children[c]

		if ok {
			n = child
		} else {
			break
		}
	}
	if n.end != 0 {
		return n.end, true
	}
	return 0, false
}
