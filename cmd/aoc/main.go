package main

import (
	"fmt"

	"github.com/jskrd/advent-of-code-2023/internal/day01"
	"github.com/jskrd/advent-of-code-2023/internal/day02"
)

type Puzzle struct {
	Name     string
	Solution func() (int, int)
}

func main() {
	puzzles := []Puzzle{
		{"Day 1: Trebuchet?!", day01.Solve},
		{"Day 2: Cube Conundrum", day02.Solve},
	}

	for i, puzzle := range puzzles {
		part1, part2 := puzzle.Solution()
		fmt.Printf("%s\n  Part One: %d\n  Part Two: %d\n", puzzle.Name, part1, part2)

		if i < len(puzzles)-1 {
			fmt.Println()
		}
	}
}
