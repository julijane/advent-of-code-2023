package main

import (
	"math"

	"github.com/julijane/advent-of-code-2023/aoc"
)

type (
	Instruction       [3]int
	InstructionBlocks [7][]Instruction
)

func runforSeed(instructionBlocks InstructionBlocks, seed int) int {
	for _, instructionBlock := range instructionBlocks {
		for _, inst := range instructionBlock {
			if seed >= inst[1] && seed < inst[1]+inst[2] {
				seed = seed + (inst[0] - inst[1])
				break
			}
		}
	}

	return seed
}

func calc(input *aoc.Input) (int, int) {
	resultPart1 := math.MaxInt
	resultPart2 := math.MaxInt

	textBlocks := input.TextBlocks()

	seeds := aoc.ExtractNumbers(textBlocks[0][0])

	instructionBlocks := InstructionBlocks{}

	for block := 1; block < 8; block++ {
		for _, line := range textBlocks[block][1:] {
			instructionData := aoc.ExtractNumbers(line)
			instructionBlocks[block-1] = append(instructionBlocks[block-1], Instruction{
				instructionData[0],
				instructionData[1],
				instructionData[2],
			})
		}
	}

	for _, seed := range seeds {
		result := runforSeed(instructionBlocks, seed)

		if result < resultPart1 {
			resultPart1 = result
		}
	}

	for num := 0; num < len(seeds); num += 2 {
		seed := seeds[num]
		length := seeds[num+1]
		for realSeed := seed; realSeed < seed+length; realSeed++ {
			result := runforSeed(instructionBlocks, realSeed)
			if result < resultPart2 {
				resultPart2 = result
			}
		}
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("Sample", "sample1.txt", calc)
	aoc.Run("Main", "input.txt", calc)
}
