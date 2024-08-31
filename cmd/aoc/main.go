package main

import (
	"fmt"
	"sync"

	"github.com/jskrd/advent-of-code-2023/internal/day01"
	"github.com/jskrd/advent-of-code-2023/internal/day02"
)

type Puzzle struct {
	name     string
	solution func() (int, int)
}

type result struct {
	name  string
	part1 int
	part2 int
}

func main() {
	puzzles := []Puzzle{
		{"Day 1: Trebuchet?!", day01.Solve},
		{"Day 2: Cube Conundrum", day02.Solve},
	}

	var wg sync.WaitGroup
	results := make([]result, len(puzzles))

	for i, puzzle := range puzzles {
		wg.Add(1)
		go func(i int, p Puzzle) {
			defer wg.Done()
			part1, part2 := p.solution()
			results[i] = result{p.name, part1, part2}
		}(i, puzzle)
	}

	wg.Wait()

	for i, result := range results {
		fmt.Printf("%s\n  Part One: %d\n  Part Two: %d\n", result.name, result.part1, result.part2)
		if i < len(results)-1 {
			fmt.Println()
		}
	}
}
