package main

import (
	"fmt"
	"regexp"
	"strconv"
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
	digitwords := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sum := 0

	for _, line := range input {
		transformed := ""

	charpos:
		for x := 0; x < len(line); x++ {
			for value, word := range digitwords {
				if strings.HasPrefix(line[x:], word) {
					transformed += strconv.Itoa(value)
					continue charpos
				}
			}
			transformed += line[x : x+1]
		}

		sum += linevalue(transformed)
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
