package main

import (
	"math"

	"github.com/julijane/advent-of-code-2023/aoc"
)

type historyPoint struct {
	c              aoc.Coordinate
	dir            int
	walkedStraight int
}

type walker struct {
	grid        *aoc.Grid
	myGrid      map[aoc.Coordinate]int
	seen        map[historyPoint]struct{}
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
		maxHeatLoss: math.MaxInt,
		seen:        make(map[historyPoint]struct{}),
		target:      aoc.Coordinate{X: grid.Width - 1, Y: grid.Height - 1},
		minStraight: minStraight,
		maxStraight: maxStriaght,
	}
}

type QueueEntry struct {
	hp       historyPoint
	heatLoss int
}

type Queue []QueueEntry

func (q *Queue) Pop() QueueEntry {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}

func (q *Queue) Push(qe QueueEntry) {
	for n, qe2 := range *q {
		if qe2.heatLoss < qe.heatLoss {
			continue
		}
		*q = append((*q)[:n], append([]QueueEntry{qe}, (*q)[n:]...)...)
		return
	}
	*q = append(*q, qe)
}

func (w *walker) walk() int {
	queue := Queue{QueueEntry{
		hp: historyPoint{
			c:              aoc.Coordinate{X: 0, Y: 0},
			dir:            1,
			walkedStraight: 0,
		},
		heatLoss: 0,
	}}

	for len(queue) > 0 {
		cur := queue.Pop()

		if !w.grid.Inside(cur.hp.c) {
			continue
		}

		if _, ok := w.seen[cur.hp]; ok {
			continue
		}

		heatLoss := cur.heatLoss + w.myGrid[cur.hp.c]

		if cur.hp.c == w.target {
			return heatLoss
		}

		w.seen[cur.hp] = struct{}{}

		if cur.hp.walkedStraight < w.maxStraight {
			queue.Push(QueueEntry{
				hp: historyPoint{
					c:              cur.hp.c.Move(cur.hp.dir),
					dir:            cur.hp.dir,
					walkedStraight: cur.hp.walkedStraight + 1,
				},
				heatLoss: heatLoss,
			})
		}

		if cur.hp.walkedStraight >= w.minStraight {
			queue.Push(QueueEntry{
				hp: historyPoint{
					c:              cur.hp.c.Move((cur.hp.dir + 3) % 4),
					dir:            (cur.hp.dir + 3) % 4,
					walkedStraight: 0,
				},
				heatLoss: heatLoss,
			})
			queue.Push(QueueEntry{
				hp: historyPoint{
					c:              cur.hp.c.Move((cur.hp.dir + 1) % 4),
					dir:            (cur.hp.dir + 1) % 4,
					walkedStraight: 0,
				},
				heatLoss: heatLoss,
			})
		}

	}

	return 0
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	grid := input.Grid()
	grid.Set(aoc.Coordinate{X: 0, Y: 0}, byte('0'))

	resultPart1 := newWalker(grid, 0, 2).walk()
	resultPart2 := newWalker(grid, 3, 9).walk()

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
