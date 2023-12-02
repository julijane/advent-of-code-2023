package main

import (
	"regexp"
	"strings"

	"github.com/julijane/advent-of-code-2023/internal/input"
)

func linevalue(line string) int {
	r := regexp.MustCompile("[^0-9]")
	res := r.ReplaceAllString(line, "")
	if len(res) == 0 {
		return 0
	}
	return int(res[0]-'0')*10 + int(res[len(res)-1]-'0')
}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		sum += linevalue(line)
	}
	return sum
}

func part2(input []string) int {
	replacements := [][2]string{
		{"one", "o1e"},
		{"two", "t2o"},
		{"three", "t3e"},
		{"four", "4"},
		{"five", "5e"},
		{"six", "6"},
		{"seven", "7n"},
		{"eight", "e8t"},
		{"nine", "n9e"},
	}

	sum := 0
	for _, line := range input {
		for _, replacement := range replacements {
			line = strings.ReplaceAll(line, replacement[0], replacement[1])
		}
		sum += linevalue(line)
	}

	return sum
}

func calc(field *input.Field) (int, int) {
	lines := field.Lines()

	sumPart1 := part1(lines)
	sumPart2 := part2(lines)

	return sumPart1, sumPart2
}

func main() {
	input.Run("Sample 1", "sample1.txt", calc)
	input.Run("Sample 2", "sample2.txt", calc)
	input.Run("Main", "input.txt", calc)
}
