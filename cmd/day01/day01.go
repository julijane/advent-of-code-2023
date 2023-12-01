package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/julijane/advent-of-code-2023/internal/aoc"
)

func linevalue(line string) int {
	r := regexp.MustCompile("[^0-9]")
	res := r.ReplaceAllString(line, "")
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

func main() {
	fmt.Println("Demo 1: ",
		part1(aoc.ReadFileAsLines("sample1.txt")))

	fmt.Println("Demo 2: ",
		part2(aoc.ReadFileAsLines("sample2.txt")))

	fmt.Println("Part 1: ",
		part1(aoc.ReadFileAsLines("input.txt")))

	fmt.Println("Part 2: ",
		part2(aoc.ReadFileAsLines("input.txt")))
}
