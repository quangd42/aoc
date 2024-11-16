package puzzle

import (
	"fmt"
	"log"
	"os"
)

type Solver func(string) int

func Run(s Solver, msg string) {
	demoFile, fCloser, err := readDemo(dir)
	defer fCloser()
	if err != nil {
		log.Fatalf("Part 1 - error reading demo file: %v", err)
	}
	println("Part 1 - demo result: ", s(demoFile))
	puzzleFile, fCloser, err := readPuzzle(dir)
	if err != nil {
		log.Fatalf("Part 1 - error reading puzzle file: %v", err)
	}
	defer fCloser()
	println("Part 1 - puzzle result: ", s(puzzleFile))

	demoFile, fCloser, err = readDemo(dir)
	defer fCloser()
	if err != nil {
		log.Fatalf("Part 2 - error reading demo file: %v", err)
	}
	println("Part 2 - demo result: ", s.Part2(demoFile))

	puzzleFile, fCloser, err = readPuzzle(dir)
	if err != nil {
		log.Fatalf("Part 2 - error reading puzzle file: %v", err)
	}
	defer fCloser()
	println("Part 2 - puzzle result: ", s.Part2(puzzleFile))
}

func readPuzzle(dir string) (*os.File, func(), error) {
	f, err := os.Open(fmt.Sprintf("%s/input.txt", dir))
	if err != nil {
		return nil, nil, err
	}
	return f, func() {
		f.Close()
	}, nil
}

func readDemo(dir string) (*os.File, func(), error) {
	f, err := os.Open(fmt.Sprintf("%s/demo.txt", dir))
	if err != nil {
		return nil, nil, err
	}
	return f, func() {
		f.Close()
	}, nil
}
