package main

import (
	"github.com/julijane/advent-of-code-2023/aoc"
)

func linevalue(line *aoc.Line) int {
	numbers := line.FindObjects(`\d`)
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0].Int()*10 + numbers[len(numbers)-1].Int()
}

func part1(field *aoc.Field) int {
	sum := 0
	for _, line := range field.Lines {
		sum += linevalue(line)
	}
	return sum
}

func part2(input *aoc.Field) int {
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

func calc(field *aoc.Field) (int, int) {
	sumPart1 := part1(field)
	sumPart2 := part2(field)

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("Sample 1", "sample1.txt", calc)
	aoc.Run("Sample 2", "sample2.txt", calc)
	aoc.Run("Main", "input.txt", calc)
}
