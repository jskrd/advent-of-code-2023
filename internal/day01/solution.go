package day01

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() (int, int) {
	dat, err := os.ReadFile("internal/day01/input.txt")
	if err != nil {
		panic(err)
	}

	input := string(dat)

	return SolvePart1(&input), SolvePart2(&input)
}

func SolvePart1(input *string) int {
	sum := 0
	for _, line := range strings.Split(*input, "\n") {
		var first rune
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = rune(line[i])
				break
			}
		}

		var last rune
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = rune(line[i])
				break
			}
		}

		number, _ := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
		sum += number
	}

	return sum
}

func SolvePart2(input *string) int {
	digits := map[string]rune{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		digits[s] = rune(s[0])
	}

	sum := 0
	for _, line := range strings.Split(*input, "\n") {
		var first, last rune
		for i := 0; i < len(line); i++ {
			for word, digit := range digits {
				high := i + len(word)

				if high > len(line) {
					continue
				}

				if line[i:high] == word {
					if first == 0 {
						first = digit
					}

					last = digit
				}
			}
		}

		number, _ := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
		sum += number
	}

	return sum
}
