package main

import "github.com/julijane/advent-of-code-2023/aoc"

func calc(field *aoc.Field) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	numbers := field.FindObjects(`\d+`)
	symbols := field.FindObjects(`[^\d\.]`)
	stars := field.FindObjects(`\*`)

	for _, number := range numbers {
		for _, symbol := range symbols {
			if number.Adjacent(symbol) {
				sumPart1 += field.ObjectAsInt(number)
				break
			}
		}
	}

	for _, star := range stars {
		var adjacentNumbers []int
		for _, number := range numbers {
			if star.Adjacent(number) {
				adjacentNumbers = append(adjacentNumbers, field.ObjectAsInt(number))
			}
		}
		if len(adjacentNumbers) == 2 {
			sumPart2 += adjacentNumbers[0] * adjacentNumbers[1]
		}
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("Sample", "sample1.txt", calc)
	aoc.Run("Main", "input.txt", calc)
}
