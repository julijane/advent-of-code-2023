package main

import (
	"math"

	"github.com/julijane/advent-of-code-2023/aoc"
)

type historyPoint struct {
	walkedStraight int
	dir            int
	c              aoc.Coordinate
}

type walker struct {
	grid        *aoc.Grid
	myGrid      map[aoc.Coordinate]int
	cache       map[historyPoint]int
	maxHeatLoss int
	target      aoc.Coordinate
	minStraight int
	maxStraight int
}

func newWalker(grid *aoc.Grid, minStraight, maxStriaght int) *walker {
	myGrid := make(map[aoc.Coordinate]int)
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			c := aoc.Coordinate{X: x, Y: y}
			myGrid[c] = grid.GetInt(c, 0)
		}
	}

	return &walker{
		grid:        grid,
		myGrid:      myGrid,
		cache:       make(map[historyPoint]int),
		maxHeatLoss: math.MaxInt,
		target:      aoc.Coordinate{X: grid.Width - 1, Y: grid.Height - 1},
		minStraight: minStraight,
		maxStraight: maxStriaght,
	}
}

func (w *walker) walk(path aoc.Coordinates, startPos aoc.Coordinate, dir int, heatLoss int, walkedStraight int) {
	if !w.grid.Inside(startPos) || path.Includes(startPos) {
		return
	}

	curHistoryPoint := historyPoint{
		c:              startPos,
		dir:            dir,
		walkedStraight: walkedStraight,
	}

	if priorLoss, ok := w.cache[curHistoryPoint]; ok {
		if priorLoss <= heatLoss {
			return
		}
	}

	w.cache[curHistoryPoint] = heatLoss
	path = append(path, startPos)

	heatLoss += w.myGrid[startPos]

	if heatLoss >= w.maxHeatLoss {
		return
	}

	if startPos == w.target {
		if heatLoss < w.maxHeatLoss {
			w.maxHeatLoss = heatLoss
		}
		return
	}

	if walkedStraight < w.maxStraight {
		w.walk(path, startPos.Move(dir), dir, heatLoss, walkedStraight+1)
	}

	if walkedStraight >= w.minStraight {
		w.walk(path, startPos.Move((dir+3)%4), (dir+3)%4, heatLoss, 0)
		w.walk(path, startPos.Move((dir+1)%4), (dir+1)%4, heatLoss, 0)
	}
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	resultPart1 := 0
	resultPart2 := 0

	grid := input.Grid()
	grid.Set(aoc.Coordinate{X: 0, Y: 0}, byte('0'))

	w := newWalker(grid, 0, 2)
	w.walk(aoc.Coordinates{}, aoc.Coordinate{X: 0, Y: 0}, 1, 0, 0)
	resultPart1 = w.maxHeatLoss

	w = newWalker(grid, 3, 9)
	w.walk(aoc.Coordinates{}, aoc.Coordinate{X: 0, Y: 0}, 1, 0, 0)
	resultPart2 = w.maxHeatLoss

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
