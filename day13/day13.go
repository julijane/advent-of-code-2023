package main

import (
	"github.com/julijane/advent-of-code-2023/aoc"
)

func rotate(lines []string) []string {
	vertical := make([]string, len(lines[0]))
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			vertical[x] = string(lines[y][x]) + vertical[x]
		}
	}

	return vertical
}

func countDiff(a, b string) int {
	count := 0
	for x := 0; x < len(a); x++ {
		if a[x] != b[x] {
			count++
		}
	}
	return count
}

func findReflection(lines []string, okDiff int) int {
	for y := 0; y < len(lines)-1; y++ {
		diffCount := 0

		ltc := min(y+1, len(lines)-y-1)
		for lc := 0; lc < ltc; lc++ {
			diffCount += countDiff(lines[y-lc], lines[y+lc+1])
		}

		if diffCount == okDiff {
			return y + 1
		}
	}

	return 0
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	resultPart1 := 0
	resultPart2 := 0

	for _, normal := range input.TextBlocks() {
		rotated := rotate(normal)

		resultPart1 += 100*findReflection(normal, 0) + findReflection(rotated, 0)
		resultPart2 += 100*findReflection(normal, 1) + findReflection(rotated, 1)
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
