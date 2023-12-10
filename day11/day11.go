package main

import (
	"math"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	resultPart1 := 0
	resultPart2 := 0

	emptyCols := []int{}
	emptyRows := []int{}

	grid := input.Grid()
	for y := 0; y < grid.Height; y++ {
		lineString := strings.ReplaceAll(string(grid.Data[y]), ".", "")
		if len(lineString) == 0 {
			emptyRows = append(emptyRows, y)
		}
	}

outer:
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			if grid.Data[y][x] != '.' {
				continue outer
			}
		}
		emptyCols = append(emptyCols, x)
	}

	galaxies := aoc.Coordinates{}
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			if grid.Data[y][x] == '#' {
				galaxies = append(galaxies, &aoc.Coordinate{X: x, Y: y})
			}
		}
	}

	for first := 0; first < len(galaxies); first++ {
		for second := first + 1; second < len(galaxies); second++ {
			distanceX := math.Abs(float64(galaxies[first].X - galaxies[second].X))
			distanceY := math.Abs(float64(galaxies[first].Y - galaxies[second].Y))
			distance := int(distanceX + distanceY)

			x1 := min(galaxies[first].X, galaxies[second].X)
			x2 := max(galaxies[first].X, galaxies[second].X)
			y1 := min(galaxies[first].Y, galaxies[second].Y)
			y2 := max(galaxies[first].Y, galaxies[second].Y)

			resultPart1 += distance
			resultPart2 += distance

			for _, emptycol := range emptyCols {
				if emptycol > x1 && emptycol < x2 {
					resultPart1 += 1
					resultPart2 += 999999
				}
			}
			for _, emptyrow := range emptyRows {
				if emptyrow > y1 && emptyrow < y2 {
					resultPart1 += 1
					resultPart2 += 999999
				}
			}
			// fmt.Println(galaxies[first], galaxies[second], distance)

		}
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, false)
	aoc.Run("input.txt", calc, true, true)
}
