package main

import (
	"github.com/julijane/advent-of-code-2023/aoc"
)

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	resultPart1 := 0
	resultPart2 := 0

	for _, history := range input.PlainLines() {
		nums := aoc.ExtractNumbers(history)

		rightmosts := []int{}
		leftmosts := []int{}

		for {
			rightmosts = append(rightmosts, nums[len(nums)-1])
			leftmosts = append(leftmosts, nums[0])

			diffs := []int{}
			allZero := nums[0] == 0

			for x := 1; x < len(nums); x++ {
				diffs = append(diffs, nums[x]-nums[x-1])
				if nums[x] != 0 {
					allZero = false
				}
			}

			if allZero {
				break
			}

			nums = diffs
		}

		for _, rightMost := range rightmosts {
			resultPart1 += rightMost
		}

		leftNum := 0
		for x := len(leftmosts) - 2; x >= 0; x-- {
			leftNum = leftmosts[x] - leftNum
		}
		resultPart2 += leftNum
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
