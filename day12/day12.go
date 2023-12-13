package main

import (
	"fmt"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

var calculatedPrior map[string]int

func countPermutations(pattern string, lengths []int) int {
	if len(pattern) == 0 {
		// we don't have any more input
		if len(lengths) == 0 {
			// but we also don't expect any more damaged springs
			// so this is one valid solution
			return 1
		}
		// we still expect damaged springs, so this is not a valid solution
		return 0
	}

	if len(lengths) == 0 {
		// we don't expect any more damaged springs
		if strings.Index(pattern, "#") > -1 {
			// but we still have some damaged springs
			// so this is not a valid solution
			return 0
		}
		// we don't have any more damaged springs
		// so this is one valid solution
		return 1
	}

	// if we have calculated this before, return the cached value
	cacheKey := fmt.Sprintf("%v%v", pattern, lengths)
	if value, ok := calculatedPrior[cacheKey]; ok {
		return value
	}

	count := 0

	if pattern[0] == '.' || pattern[0] == '?' {
		// this is or can be a working spring, so lets first try that
		count += countPermutations(pattern[1:], lengths)
	}

	if pattern[0] == '#' || pattern[0] == '?' {
		// this is or can be damaged spring
		needLength := lengths[0]

		// we need to have needLength consecutive damaged springs
		if len(pattern) >= needLength && strings.Index(pattern[:needLength], ".") == -1 {
			// we need to match exactly or must be followed by a working spring
			if len(pattern) == needLength || pattern[needLength] != '#' {
				// count recursively, but we need to keep one field separation between blocks of damaged springs
				count += countPermutations(
					aoc.StringFrom(pattern, needLength+1),
					aoc.SliceFrom(lengths, 1))
			}
		}
	}

	calculatedPrior[cacheKey] = count
	return count
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	resultPart1 := 0
	resultPart2 := 0

	calculatedPrior = make(map[string]int)

	for _, line := range input.PlainLines() {
		parts := strings.Split(line, " ")

		pattern := parts[0]

		lengths := []int{}
		for _, length := range strings.Split(parts[1], ",") {
			lengths = append(lengths, aoc.Atoi(length))
		}

		part2lenghts := []int{}
		for x := 0; x < 5; x++ {
			part2lenghts = append(part2lenghts, lengths...)
		}

		resultPart1 += countPermutations(pattern, lengths)

		resultPart2 += countPermutations(
			fmt.Sprintf("%s?%s?%s?%s?%s", pattern, pattern, pattern, pattern, pattern),
			part2lenghts)

	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
