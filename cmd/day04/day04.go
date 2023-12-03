package main

import "github.com/julijane/advent-of-code-2023/aoc"

func calc(field *aoc.Field) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("Sample", "sample1.txt", calc)
	aoc.Run("Main", "input.txt", calc)
}
