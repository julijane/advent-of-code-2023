package main

import (
	"slices"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func rotateCW(lines []string) []string {
	lineLength := len(lines[0])

	vertical := make([]string, lineLength)
	for y := 0; y < lineLength; y++ {
		for x := 0; x < len(lines); x++ {
			vertical[y] = string(lines[x][y]) + vertical[y]
		}
	}

	return vertical
}

func tilt(platform []string) []string {
	for y := 0; y < len(platform); y++ {
		tiltedLine := ""

		for n, group := range strings.Split(platform[y], "#") {
			if n > 0 {
				tiltedLine += "#"
			}

			bytes := []byte(group)
			slices.Sort(bytes)
			tiltedLine += string(bytes)
		}
		platform[y] = tiltedLine
	}

	return platform
}

func sumPlatform(platform []string) int {
	sum := 0
	for _, line := range platform {
		for x := 0; x < len(line); x++ {
			if line[x] == 'O' {
				sum += x + 1
			}
		}
	}
	return sum
}

func calcPart1(lines []string) int {
	platform := tilt(rotateCW(lines))
	return sumPlatform(platform)
}

func calcPart2(lines []string) int {
	platform := lines

	sequence := []string{}
	repeatedFrom := -1

	for {
		platformString := strings.Join(platform, "+")
		repeatedFrom = slices.Index(sequence, platformString)
		if repeatedFrom != -1 {
			break
		}

		sequence = append(sequence, platformString)

		for r := 0; r < 4; r++ {
			platform = tilt(rotateCW(platform))
		}
	}

	loopLength := len(sequence) - repeatedFrom
	finalPos := repeatedFrom + (1000000000-repeatedFrom)%loopLength

	platform = strings.Split(sequence[finalPos], "+")

	return sumPlatform(rotateCW(platform))
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	resultPart1 := calcPart1(input.PlainLines())
	resultPart2 := calcPart2(input.PlainLines())

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
