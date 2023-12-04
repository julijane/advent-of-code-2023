package main

import (
	"math"

	"github.com/julijane/advent-of-code-2023/aoc"
)

type Blocks [7][][3]int64

func runforSeed(blocks Blocks, seed int64) int64 {
	for _, block := range blocks {
		for _, inst := range block {
			if seed >= inst[1] && seed < inst[1]+inst[2] {
				seed = seed + (inst[0] - inst[1])
				break
			}
		}
	}

	return seed
}

func calc(input *aoc.Input) (int, int) {
	resultPart1 := int64(math.MaxInt64)
	resultPart2 := int64(math.MaxInt64)

	seeds := aoc.ExtractNumbers(input.Lines[0].Data)

	blocks := [7][][3]int64{}
	line := 3
	block := 0

	for {
		if line >= len(input.Lines) {
			break
		}

		lineText := input.Lines[line].Data

		if lineText == "" {
			block++
			line += 2
			continue
		}

		lineData := aoc.ExtractNumbers(lineText)
		blocks[block] = append(blocks[block], [3]int64{
			int64(lineData[0]),
			int64(lineData[1]),
			int64(lineData[2]),
		})
		line++
	}

	for _, seed := range seeds {
		result := runforSeed(blocks, int64(seed))

		if result < resultPart1 {
			resultPart1 = result
		}
	}

	for num := 0; num < len(seeds); num += 2 {
		seed := seeds[num]
		length := seeds[num+1]
		for realSeed := seed; realSeed < seed+length; realSeed++ {
			result := runforSeed(blocks, int64(realSeed))
			if result < resultPart2 {
				resultPart2 = result
			}
		}
	}

	return int(resultPart1), int(resultPart2)
}

func main() {
	aoc.Run("Sample", "sample1.txt", calc)
	aoc.Run("Main", "input.txt", calc)
}
