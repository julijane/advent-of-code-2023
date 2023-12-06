package main

import (
	"math"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func calcVariants(raceTime, toBeat int) int {
	rf := float64(raceTime)
	tbf := float64(toBeat)

	sq := math.Sqrt(rf*rf - 4*tbf)
	return int(
		math.Ceil((rf+sq)/2) - math.Floor((rf-sq)/2) - 1)

	// for buttonTime := 0; buttonTime < raceTime; buttonTime++ {
	// 	if (raceTime-buttonTime)*buttonTime > toBeat {
	// 		return raceTime - 2*buttonTime + 1
	// 	}
	// }
	//
	// return 0
}

func getSingleNumber(line string) int {
	numString := ""
	for _, match := range aoc.ExtractRegexps(line, `\d`) {
		numString += match
	}
	return aoc.Atoi(numString)
}

func calc(input *aoc.Input) (int, int) {
	resultPart1 := 1

	// Part 1
	times := aoc.ExtractNumbers(input.Lines[0].Data)
	distances := aoc.ExtractNumbers(input.Lines[1].Data)

	for race := range times {
		resultPart1 *= calcVariants(times[race], distances[race])
	}

	// Part 2
	resultPart2 := calcVariants(
		getSingleNumber(input.Lines[0].Data),
		getSingleNumber(input.Lines[1].Data))

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("Sample", "sample1.txt", calc)
	aoc.Run("Main", "input.txt", calc)
}
