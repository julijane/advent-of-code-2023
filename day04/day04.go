package main

import (
	"math"
	"regexp"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func calc(input *aoc.Input) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	r := regexp.MustCompile(`\s+`)

	numCopies := make([]int, len(input.Lines))
	for y := 0; y < len(input.Lines); y++ {
		numCopies[y] = 1
	}

	for lineNo, line := range input.Lines {
		split1 := strings.Split(line.Data, ": ")
		split2 := strings.Split(split1[1], " | ")

		winners := r.Split(split2[0], -1)
		picks := r.Split(split2[1], -1)

		numWins := 0
		for _, pick := range picks {
			if pick != "" {
				for _, winner := range winners {
					if pick == winner {
						numWins++
						break
					}
				}
			}
		}

		if numWins >= 0 {
			sumPart1 += int(math.Pow(2, float64(numWins-1)))

			for y := 0; y < numWins; y++ {
				numCopies[lineNo+y+1] += numCopies[lineNo]
			}
		}
	}

	for y := 0; y < len(input.Lines); y++ {
		sumPart2 += numCopies[y]
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("Sample", "sample1.txt", calc)
	aoc.Run("Main", "input.txt", calc)
}
