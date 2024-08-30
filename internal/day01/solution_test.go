package day01_test

import (
	"testing"

	"github.com/jskrd/advent-of-code-2023/internal/day01"
)

func TestSolvePart1(t *testing.T) {
	input := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
	expected := "142"
	actual := day01.SolvePart1(&input)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	expected := "281"
	actual := day01.SolvePart2(&input)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
