package main

import (
	"github.com/julijane/advent-of-code-2023/aoc"
)

type Walker struct {
	grid          *aoc.Grid
	alreadyWalked map[aoc.Pointer]struct{}
	energized     map[aoc.Coordinate]struct{}
}

func (w *Walker) walk(p aoc.Pointer) {
	w.energized[p.C] = struct{}{}
	if _, ok := w.alreadyWalked[p]; ok {
		return
	}

	w.alreadyWalked[p] = struct{}{}

	p.Move()

	element := w.grid.Get(p.C, 'X')
	if element == 'X' {
		return
	}

	if element == '/' {
		if p.IsUpOrDown() {
			p.TurnRight()
		} else {
			p.TurnLeft()
		}
	} else if element == '\\' {
		if p.IsUpOrDown() {
			p.TurnLeft()
		} else {
			p.TurnRight()
		}
	} else if element != '.' &&
		(element != '|' || !p.IsUpOrDown()) &&
		(element != '-' || !p.IsLeftOrRight()) {
		p.TurnLeft()
		w.walk(p)
		p.TurnAround()
	}

	w.walk(p)
}

func walk(grid *aoc.Grid, p aoc.Pointer) int {
	w := Walker{
		grid:          grid,
		alreadyWalked: make(map[aoc.Pointer]struct{}),
		energized:     make(map[aoc.Coordinate]struct{}),
	}

	w.walk(p)

	return len(w.energized) - 1
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	grid := input.Grid()

	resultPart1 := walk(grid, aoc.NewPointer(-1, 0, 1))

	resultPart2 := 0
	for startX := 0; startX < grid.Width; startX++ {
		resultPart2 = max(
			resultPart2,
			max(
				walk(grid, aoc.NewPointer(startX, -1, 2)),
				walk(grid, aoc.NewPointer(startX, grid.Height, 0)),
			))
	}

	for startY := 0; startY < grid.Height; startY++ {
		resultPart2 = max(
			resultPart2,
			max(
				walk(grid, aoc.NewPointer(-1, startY, 1)),
				walk(grid, aoc.NewPointer(grid.Width, startY, 3)),
			))
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
