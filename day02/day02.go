package main

import (
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	for game, line := range input.Lines {
		splitted := strings.Split(line.Data, ";")

		gamePossible := true

		needRed := 0
		needGreen := 0
		needBlue := 0

		for _, draw := range splitted {
			red := aoc.RegexpSubmatchAsInt(draw, `(\d+) red`)
			green := aoc.RegexpSubmatchAsInt(draw, `(\d+) green`)
			blue := aoc.RegexpSubmatchAsInt(draw, `(\d+) blue`)

			if red > 12 || green > 13 || blue > 14 {
				gamePossible = false
			}

			needRed = max(red, needRed)
			needGreen = max(green, needGreen)
			needBlue = max(blue, needBlue)
		}

		if gamePossible {
			sumPart1 += game + 1
		}

		sumPart2 += needRed * needGreen * needBlue
	}
	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
