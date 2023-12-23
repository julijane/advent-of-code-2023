package main

import (
	"regexp"
	"slices"

	"github.com/julijane/advent-of-code-2023/aoc"
)

type Brick struct {
	id     int
	x1, y1 int
	z      int
	x2, y2 int
	height int
}

func (b *Brick) Overlaps(other *Brick) bool {
	if max(b.x1, other.x1) <= min(b.x2, other.x2) &&
		max(b.y1, other.y1) <= min(b.y2, other.y2) {
		return true
	}

	return false
}

type Bricks []Brick

func NewBricks(input aoc.Input) (bricks Bricks, maxX int, maxY int) {
	bricks = Bricks{}

	re := regexp.MustCompile(`(\d+),(\d+),(\d+)~(\d+),(\d+),(\d+)`)
	matches := re.FindAllStringSubmatch(input.SingleString(), -1)

	for idx, match := range matches {
		x1 := aoc.Atoi(match[1])
		y1 := aoc.Atoi(match[2])
		z1 := aoc.Atoi(match[3])

		x2 := aoc.Atoi(match[4])
		y2 := aoc.Atoi(match[5])
		z2 := aoc.Atoi(match[6])

		maxX = max(maxX, x2)
		maxY = max(maxY, y2)

		bricks = append(bricks, Brick{
			id: idx,
			x1: x1,
			y1: y1,

			x2: x2,
			y2: y2,

			z:      z1,
			height: z2 - z1 + 1,
		})
	}

	return bricks, maxX, maxY
}

func (b Bricks) Sort() {
	slices.SortFunc(b, func(a, b Brick) int {
		if a.z < b.z {
			return -1
		}
		if a.z > b.z {
			return 1
		}
		return 0
	})
}

func (b Bricks) Fall(maxX, maxY int) int {
	heights := aoc.New2DSlice[int](maxY+1, maxX+1)
	maxZ := 0

	for idx, brick := range b {
		fallsTo := 1

		for y := brick.y1; y <= brick.y2; y++ {
			for x := brick.x1; x <= brick.x2; x++ {
				if heights[y][x] >= fallsTo {
					fallsTo = heights[y][x] + 1
				}
			}
		}

		b[idx].z = fallsTo

		for y := brick.y1; y <= brick.y2; y++ {
			for x := brick.x1; x <= brick.x2; x++ {
				heights[y][x] = fallsTo + brick.height - 1
			}
		}

		if brick.z+brick.height > maxZ {
			maxZ = brick.z + brick.height
		}
	}

	return maxZ
}

func (b Bricks) CalcDependencies(maxZ int) (brickSupports map[int][]int, brickSupportedBy map[int][]int) {
	brickSupports = make(map[int][]int)
	brickSupportedBy = make(map[int][]int)

	for z := 1; z <= maxZ; z++ {
		// find bricks that have their top at this level
		tops := []Brick{}
		for _, brick := range b {
			if brick.z+brick.height-1 == z {
				tops = append(tops, brick)
			}
		}

		// go through the bricks that are lying on top of this level
		for _, above := range b {
			if above.z != z+1 {
				continue
			}

			supportedBy := []int{}
			for _, below := range tops {
				if above.Overlaps(&below) {
					supportedBy = append(supportedBy, below.id)
					brickSupports[below.id] = append(brickSupports[below.id], above.id)
				}
			}

			brickSupportedBy[above.id] = supportedBy

		}
	}

	return brickSupports, brickSupportedBy
}

func (b Bricks) Part1(brickSupports map[int][]int, brickSupportedBy map[int][]int) int {
	blocksCanDisintegrate := 0

outerloop:
	for _, brick := range b {
		// brick can be disintegrated if bricks that lie on top of it have other support
		for _, brickAbove := range brickSupports[brick.id] {
			if len(brickSupportedBy[brickAbove]) == 1 {
				// can't be disintegrated
				continue outerloop
			}
		}

		blocksCanDisintegrate++
	}

	return blocksCanDisintegrate
}

func (b Bricks) Part2(brickSupports map[int][]int, brickSupportedBy map[int][]int) int {
	sumPart2 := 0

	for _, brick := range b {
		fallen := []int{brick.id}
		x := 0

		for x < len(fallen) {
		checkcascade:
			for _, brickAbove := range brickSupports[fallen[x]] {
				for _, supporter := range brickSupportedBy[brickAbove] {
					if !slices.Contains(fallen, supporter) {
						// brick is still held
						continue checkcascade
					}
				}

				if !slices.Contains(fallen, brickAbove) {
					fallen = append(fallen, brickAbove)
				}
			}
			x++
		}

		sumPart2 += len(fallen) - 1
	}

	return sumPart2
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	bricks, maxX, maxY := NewBricks(*input)
	bricks.Sort()

	maxZ := bricks.Fall(maxX, maxY)

	brickSupports, brickSupportedBy := bricks.CalcDependencies(maxZ)

	resultPart1 := bricks.Part1(brickSupports, brickSupportedBy)
	resultPart2 := bricks.Part2(brickSupports, brickSupportedBy)

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
