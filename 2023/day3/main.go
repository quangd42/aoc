package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"

	"github.com/quangd42/aoc/2023/puzzle"
)

type Solver struct{}

type number struct {
	val, xStart, xEnd, y int
}

// searchGrid searches 3x3 grid to find digit
// if found digit, search left and right to find number
// add number to map to track parsed numbers
func searchGrid(m []string, x, y int, tracker map[number]bool) []int {
	maxY := len(m) - 1
	maxX := len(m[0]) - 1
	out := []int{}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if x+i < 0 || x+i > maxX || y+j < 0 || y+j > maxY {
				continue
			}
			if b := m[y+j][x+i]; b >= '0' && b <= '9' {
				n := number{
					y: y + j,
				}
				for start := x + i; start >= 0; start-- {
					if m[y+j][start] >= '0' && m[y+j][start] <= '9' {
						// log.Printf("walking left: %s", string(m[y+j][start]))
						n.xStart = start
						continue
					}
					break
				}
				for end := x + i; end <= maxX; end++ {
					if m[y+j][end] >= '0' && m[y+j][end] <= '9' {
						// log.Printf("walking right: %s\n", string(m[y+j][end]))
						n.xEnd = end
						continue
					}
					break
				}
				val, err := strconv.Atoi(string(m[y+j][n.xStart : n.xEnd+1]))
				if err != nil {
					log.Printf("current cell: %s", string(b))
					log.Printf("current coord: %d %d", x+i, y+j)
					log.Printf("number: %#v", n)
					log.Fatal(err)
				}
				n.val = val
				if _, ok := tracker[n]; !ok {
					// log.Printf("parsed number: %#v", n)
					out = append(out, val)
					tracker[n] = true
				}
			}
		}
	}
	return out
}

func (s Solver) Part1(f *os.File) int {
	scanner := bufio.NewScanner(f)
	// make matrix from file
	// look for the symbol
	matrix := []string{}
	tracker := map[number]bool{}
	res := 0

	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	for y, s := range matrix {
		for x, r := range s {
			if string(r) == "." || unicode.IsDigit(r) {
				continue
			}
			// log.Printf("current cell: %s", string(r))
			// if found symbol,
			for _, n := range searchGrid(matrix, x, y, tracker) {
				res += n
			}
		}
	}
	return res
}

func (s Solver) Part2(f *os.File) int {
	scanner := bufio.NewScanner(f)
	// make matrix from file
	// look for the symbol
	matrix := []string{}
	tracker := map[number]bool{}
	res := 0

	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	for y, s := range matrix {
		for x, r := range s {
			if string(r) == "*" {
				// log.Printf("current cell: %s", string(r))
				numbers := searchGrid(matrix, x, y, tracker)
				if len(numbers) == 2 {
					res += numbers[0] * numbers[1]
				}
			}
		}
	}
	return res
}

func main() {
	puzzle.Run(Solver{}, "day3")
}
