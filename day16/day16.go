package main

import (
	"github.com/julijane/advent-of-code-2023/aoc"
)

type Pointer struct {
	c         aoc.Coordinate
	direction int
}

var (
	alreadyWalked = make(map[Pointer]struct{})
	energized     = make(map[aoc.Coordinate]struct{})
	grid          *aoc.Grid
)

func reset() {
	alreadyWalked = make(map[Pointer]struct{})
	energized = make(map[aoc.Coordinate]struct{})
}

func walk(startPoint aoc.Coordinate, direction int) {
	energized[startPoint] = struct{}{}
	if _, ok := alreadyWalked[Pointer{startPoint, direction}]; ok {
		return
	}
	alreadyWalked[Pointer{startPoint, direction}] = struct{}{}

	newPos := startPoint.Copy().Move(direction)
	element := grid.Get(newPos, 'X')

	// dirvis := [4]string{"↑", "→", "↓", "←"}
	// fmt.Printf("%v %v %v %c\n", startPoint, dirvis[direction], newPos, element)

	if element == 'X' {
		return
	}

	if element == '/' {
		if direction == 0 || direction == 2 {
			direction = direction + 1
		} else {
			direction = direction - 1
		}
		walk(newPos, direction)
		return
	}

	if element == '\\' {
		if direction == 0 || direction == 2 {
			direction = (direction + 3) % 4
		} else {
			direction = (direction + 1) % 4
		}
		walk(newPos, direction)
		return
	}

	if element == '.' ||
		(element == '|' && (direction == 0 || direction == 2)) ||
		(element == '-' && (direction == 1 || direction == 3)) {
		walk(newPos, direction)
		return
	}

	walk(newPos, (direction+3)%4)
	walk(newPos, (direction+1)%4)
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	resultPart1 := 0
	resultPart2 := 0

	grid = input.Grid()

	reset()
	walk(aoc.Coordinate{X: -1, Y: 0}, 1)
	resultPart1 = len(energized) - 1

	for startX := 0; startX < grid.Width; startX++ {
		reset()
		walk(aoc.Coordinate{X: startX, Y: -1}, 2)

		resultPart2 = max(resultPart2, len(energized))

		reset()
		walk(aoc.Coordinate{X: startX, Y: grid.Height}, 0)

		resultPart2 = max(resultPart2, len(energized))

	}

	for startY := 0; startY < grid.Height; startY++ {
		reset()
		walk(aoc.Coordinate{X: -1, Y: startY}, 1)

		resultPart2 = max(resultPart2, len(energized))

		reset()
		walk(aoc.Coordinate{X: grid.Width, Y: startY}, 3)

		resultPart2 = max(resultPart2, len(energized))
	}

	resultPart2--

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
