package main

import (
	"strconv"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

var calculatedPrior map[string]map[string]int

func intSliceString(ints []int) string {
	result := ""
	for x, i := range ints {
		if x > 0 {
			result += ","
		}
		result += strconv.Itoa(i)
	}
	return result
}

func countPermutations(pattern string, lengths []int) int {
	lenString := intSliceString(lengths)
	if _, ok := calculatedPrior[pattern]; ok {
		if _, ok := calculatedPrior[pattern][lenString]; ok {
			return calculatedPrior[pattern][lenString]
		}
	}

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

	count := 0

	if pattern[0] == '.' || pattern[0] == '?' {
		// this can be a working spring, so lets first try that
		count += countPermutations(pattern[1:], lengths)
	}

	if pattern[0] == '#' || pattern[0] == '?' {
		// this is or can be damaged spring
		needLength := lengths[0]

		// we need to have needLength consecutive damaged springs
		if len(pattern) >= needLength && strings.Index(pattern[:needLength], ".") == -1 {
			// we need to match exactly or must be followed by a working spring
			if len(pattern) == needLength || pattern[needLength] != '#' {
				// count recursively, but we need to keep separation between blocks of damaged springs
				subpattern := ""
				if len(pattern) > needLength {
					subpattern = pattern[needLength+1:]
				}
				sublengths := []int{}
				if len(lengths) > 1 {
					sublengths = lengths[1:]
				}

				count += countPermutations(subpattern, sublengths)
			}
		}
	}

	if _, ok := calculatedPrior[pattern]; !ok {
		calculatedPrior[pattern] = make(map[string]int)
	}

	calculatedPrior[pattern][lenString] = count
	return count
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	resultPart1 := 0
	resultPart2 := 0

	calculatedPrior = make(map[string]map[string]int)

	for _, line := range input.PlainLines() {
		parts := strings.Split(line, " ")

		pattern := parts[0]

		lengths := []int{}
		for _, length := range strings.Split(parts[1], ",") {
			lengths = append(lengths, aoc.Atoi(length))
		}

		res := countPermutations(pattern, lengths)
		resultPart1 += res

		pattern = pattern + "?" + pattern + "?" + pattern + "?" + pattern + "?" + pattern
		newlengths := append(lengths, lengths...)
		newlengths = append(newlengths, lengths...)
		newlengths = append(newlengths, lengths...)
		newlengths = append(newlengths, lengths...)

		res = countPermutations(pattern, newlengths)
		resultPart2 += res

	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
