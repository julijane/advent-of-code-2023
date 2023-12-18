package main

import (
	"strconv"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func getArea(points aoc.Coordinates) int {
	// First get the area enclosed by the polygon
	// this is taken from day10
	area := 0
	point0 := points[0]
	for _, point1 := range points[1:] {
		area += point0.Y*point1.X - point0.X*point1.Y
		point0 = point1
	}
	if area < 0 {
		area = -area
	}

	// But we also need to add the fields on the perimeter
	// so lets add them up
	perimeterFields := 0
	for x := 1; x < len(points); x++ {
		perimeterFields += aoc.AbsInt(points[x].X-points[x-1].X) + aoc.AbsInt(points[x].Y-points[x-1].Y)
	}

	return (area+perimeterFields)/2 + 1
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	currentPos := [2]aoc.Coordinate{{X: 0, Y: 0}, {X: 0, Y: 0}}
	polygon := [2]aoc.Coordinates{{currentPos[0]}, {currentPos[1]}}

	for _, line := range input.Lines {
		parts := strings.Split(line.Data, " ")

		steps0 := aoc.Atoi(parts[1])
		dir0 := strings.Index("RULD", parts[0])

		steps1, _ := strconv.ParseInt(parts[2][2:7], 16, 64)
		dir1 := int(parts[2][7]-'0'+3) % 4

		currentPos[0] = currentPos[0].MoveBy(dir0, steps0)
		currentPos[1] = currentPos[1].MoveBy(dir1, int(steps1))

		polygon[0] = append(polygon[0], currentPos[0])
		polygon[1] = append(polygon[1], currentPos[1])
	}

	return getArea(polygon[0]), getArea(polygon[1])
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
