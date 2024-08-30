package main

import (
	"fmt"

	"github.com/jskrd/advent-of-code-2023/internal/day01"
)

type Puzzle struct {
	Name     string
	Solution func() (string, string)
}

func main() {
	puzzles := []Puzzle{
		{"Day 1: Trebuchet?!", day01.Solve},
	}

	for i, puzzle := range puzzles {
		part1, part2 := puzzle.Solution()
		fmt.Printf("%s\n  Part One: %s\n  Part Two: %s\n", puzzle.Name, part1, part2)

		if i < len(puzzles)-1 {
			fmt.Println()
		}
	}
}
