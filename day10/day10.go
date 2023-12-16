package main

import (
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func getConnectedNeighbors(grid *aoc.Grid, currentPos aoc.Coordinate) *aoc.Coordinates {
	results := aoc.Coordinates{}

	otherPos := [4]aoc.Coordinate{
		currentPos.Above(),
		currentPos.Left(),
		currentPos.Below(),
		currentPos.Right(),
	}

	thisContent := grid.Get(currentPos, '#')

	// first string is what the other field is allowed to be
	// second string is what we then must be
	okList := [4][2]string{
		{"S|7F", "S|JL"}, // above
		{"S-LF", "S-J7"}, // left
		{"S|JL", "S|7F"}, // below
		{"S-J7", "S-FL"}, // right
	}

	for neighbor := 0; neighbor < 4; neighbor++ {
		neighborPos := otherPos[neighbor]
		neighborContent := grid.Get(neighborPos, '#')

		if strings.IndexAny(okList[neighbor][0], string(neighborContent)) > -1 &&
			strings.Index(okList[neighbor][1], string(thisContent)) > -1 {
			results = append(results, neighborPos)
		}
	}

	return &results
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	resultPart2 := 0

	grid := input.Grid()
	snake := aoc.Coordinates{}
	start := grid.Find('S')
	currentPos := start.Copy()
	priorPos := aoc.Coordinate{X: -1, Y: -1}
steps:
	for {
		if start.Same(currentPos) && len(snake) > 0 {
			break
		}
		snake = append(snake, currentPos)
		for _, connectedNeighbor := range *getConnectedNeighbors(grid, currentPos) {
			if !priorPos.Same(connectedNeighbor) {
				priorPos = currentPos
				currentPos = connectedNeighbor
				continue steps
			}
		}

		panic("Could not move")
	}
	resultPart1 := len(snake) / 2

	// part 2
	// Gauss's area formula / Shoe Lace formula
	// -> https://en.wikipedia.org/wiki/Shoelace_formula
	area := 0
	point0 := snake[len(snake)-1]
	for _, point1 := range snake {
		area += point0.Y*point1.X - point0.X*point1.Y
		point0 = point1
	}
	if area < 0 {
		area = -area
	}

	resultPart2 = (area-len(snake))/2 + 1

	printGrid(*grid, snake)

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, false)
	aoc.Run("sample2.txt", calc, false, true)
	aoc.Run("sample3.txt", calc, false, true)
	aoc.Run("sample4.txt", calc, false, true)
	aoc.Run("input.txt", calc, true, true)
}
