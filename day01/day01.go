package main

import (
	"github.com/julijane/advent-of-code-2023/aoc"
)

func part1(input *aoc.Input) int {
	sum := 0
	for _, line := range input.Lines {
		digits := aoc.ExtractDigits(line.Data)

		if len(digits) > 0 {
			sum += digits[0]*10 + digits[len(digits)-1]
		}
	}
	return sum
}

func part2(input *aoc.Input) int {
	replacements := [][2]string{
		{"one", "o1e"},
		{"two", "t2o"},
		{"three", "t3ree"},
		{"four", "f4ur"},
		{"five", "f5ve"},
		{"six", "s6x"},
		{"seven", "s7ven"},
		{"eight", "e8ght"},
		{"nine", "n9ne"},
	}

	for _, line := range input.Lines {
		for _, replacement := range replacements {
			line.ReplaceText(replacement[0], replacement[1])
		}
	}

	return part1(input)
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	if runPart1 {
		sumPart1 = part1(input)
	}

	if runPart2 {
		sumPart2 = part2(input)
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, false)
	aoc.Run("sample2.txt", calc, false, true)
	aoc.Run("input.txt", calc, true, true)
}
