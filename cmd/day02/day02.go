package main

import (
	"regexp"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func calc(field *aoc.Field) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	regBlue := regexp.MustCompile(`(\d+) blue`)
	regRed := regexp.MustCompile(`(\d+) red`)
	regGreen := regexp.MustCompile(`(\d+) green`)

	for game, line := range field.Lines() {
		split1 := strings.Split(line.Text, ": ")
		split2 := strings.Split(split1[1], ";")

		gamePossible := true
		needRed := 0
		needGreen := 0
		needBlue := 0

		for _, draw := range split2 {
			matchRed := regRed.FindStringSubmatch(draw)
			matchGreen := regGreen.FindStringSubmatch(draw)
			matchBlue := regBlue.FindStringSubmatch(draw)

			red := aoc.Atoi(aoc.SliceMemberOrEmptyString(matchRed, 1))
			green := aoc.Atoi(aoc.SliceMemberOrEmptyString(matchGreen, 1))
			blue := aoc.Atoi(aoc.SliceMemberOrEmptyString(matchBlue, 1))

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
	aoc.Run("Sample", "sample1.txt", calc)
	aoc.Run("Main", "input.txt", calc)
}
