package main

import (
	"regexp"
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

	re := regexp.MustCompile(`([a-z]+)([=-]+)([0-9]+)?`)

	parts := strings.Split(inputLine, ",")
	for _, part := range parts {
		resultPart1 += hash(part)

		matches := re.FindAllStringSubmatch(part, -1)

		label := matches[0][1]
		operation := matches[0][2]
		targetBox := hash(label)

		if operation == "-" {
			boxes[targetBox] = slices.DeleteFunc(
				boxes[targetBox],
				func(s string) bool { return s == label },
			)
			continue
		}

		focalLengths[label] = aoc.Atoi(matches[0][3])

		if !slices.Contains(boxes[targetBox], label) {
			boxes[targetBox] = append(boxes[targetBox], label)
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
