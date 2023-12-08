package main

import (
	"math"
	"slices"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	numCopies := make([]int, len(input.Lines))

	for lineNo, line := range input.Lines {
		splitted := strings.Split(line.Data, " | ")

		winners := aoc.ExtractNumbers(splitted[0])[1:]
		picks := aoc.ExtractNumbers(splitted[1])

		numWins := 0
		for _, pick := range picks {
			if slices.Contains(winners, pick) {
				numWins++
			}
		}

		sumPart1 += int(math.Pow(2, float64(numWins-1)))

		for y := 0; y < numWins; y++ {
			numCopies[lineNo+y+1] += numCopies[lineNo] + 1
		}
	}

	for y := 0; y < len(input.Lines); y++ {
		sumPart2 += numCopies[y] + 1
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
