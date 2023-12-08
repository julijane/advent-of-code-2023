package main

import (
	"github.com/julijane/advent-of-code-2023/aoc"
)

type (
	Instructions map[string]Instruction
	Instruction  struct {
		Left  string
		Right string
	}
)

func part1(turns string, instructions Instructions) int {
	currentPos := "AAA"
	numTurns := 0
	for {
		for _, turn := range turns {
			if turn == 'L' {
				currentPos = instructions[currentPos].Left
			} else {
				currentPos = instructions[currentPos].Right
			}
			numTurns++
			if currentPos == "ZZZ" {
				return numTurns
			}
		}
	}
}

// greatest common divisor
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// least common multiple
func LCM(inputs []int) int {
	result := inputs[0] * inputs[1] / GCD(inputs[0], inputs[1])

	for i := 2; i < len(inputs); i++ {
		result = LCM([]int{result, inputs[i]})
	}

	return result
}

func part2(turns string, instructions Instructions) int {
	var ghostTurns []int

	for startingPoint := range instructions {
		if startingPoint[2] == 'A' {
			currentPos := startingPoint
			numTurns := 0

		outer:
			for {
				for _, turn := range turns {
					if turn == 'L' {
						currentPos = instructions[currentPos].Left
					} else {
						currentPos = instructions[currentPos].Right
					}
					numTurns++
					if currentPos[2] == 'Z' {
						ghostTurns = append(ghostTurns, numTurns)
						break outer
					}
				}
			}

		}
	}

	return LCM(ghostTurns)
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	turns := input.Lines[0].Data

	instructions := make(map[string]Instruction)
	for _, line := range input.Lines[2:] {
		instruction := Instruction{
			Left:  line.Data[7:10],
			Right: line.Data[12:15],
		}
		instructions[line.Data[0:3]] = instruction
	}

	resultPart1 := 0
	resultPart2 := 0

	if runPart1 {
		resultPart1 = part1(turns, instructions)
	}

	if runPart2 {
		resultPart2 = part2(turns, instructions)
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, false)
	aoc.Run("sample2.txt", calc, true, false)
	aoc.Run("sample3.txt", calc, false, true)
	aoc.Run("input.txt", calc, true, true)
}
