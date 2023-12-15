package main

import (
	"slices"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func hash(input string) int {
	value := uint8(0)
	for _, c := range input {
		value += uint8(c)
		value *= 17
	}

	return int(value)
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	resultPart1 := 0
	resultPart2 := 0

	inputLine := input.PlainLines()[0]

	focalLengths := make(map[string]int)
	boxes := [256][]string{}

	parts := strings.Split(inputLine, ",")
	for _, part := range parts {
		resultPart1 += hash(part)

		posEqual := strings.Index(part, "=")
		posDash := strings.Index(part, "-")
		if posEqual >= 0 {
			focalLengths[part[:posEqual]] = aoc.Atoi(part[posEqual+1:])
		}

		label := part[:max(posDash, posEqual)]
		targetBox := hash(label)

		posLens := slices.Index(boxes[targetBox], label)
		if posDash > -1 {
			if posLens > -1 {
				boxes[targetBox] = append(boxes[targetBox][:posLens], boxes[targetBox][posLens+1:]...)
			}
		} else {
			if posLens < 0 {
				boxes[targetBox] = append(boxes[targetBox], label)
			}
		}
	}

	for m, box := range boxes {
		for n, lens := range box {
			resultPart2 += (m + 1) * (n + 1) * focalLengths[lens]
		}
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
