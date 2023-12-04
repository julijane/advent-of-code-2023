package main

import (
	"regexp"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func linevalue(line *aoc.Line) int {
	re := regexp.MustCompile(`[^\d]+`)
	digits := re.ReplaceAllString(line.Data, "")

	if len(digits) == 0 {
		return 0
	}

	return int(digits[0]-'0')*10 + int(digits[len(digits)-1]-'0')
}

func part1(input *aoc.Input) int {
	sum := 0
	for _, line := range input.Lines {
		sum += linevalue(line)
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

	sum := 0
	for _, line := range input.Lines {
		for _, replacement := range replacements {
			line.ReplaceText(replacement[0], replacement[1])
		}
		sum += linevalue(line)
	}

	return sum
}

func calc(input *aoc.Input) (int, int) {
	sumPart1 := part1(input)
	sumPart2 := part2(input)

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("Sample 1", "sample1.txt", calc)
	aoc.Run("Sample 2", "sample2.txt", calc)
	aoc.Run("Main", "input.txt", calc)
}
