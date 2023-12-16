package main

import (
	"fmt"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

func printGrid(grid aoc.Grid, snake aoc.Coordinates) {
	// cleanup the trash
	grid.Map(func(pos aoc.Coordinate, value byte) byte {
		if !snake.Includes(pos) {
			return ' '
		}
		return value
	})

	replacer := strings.NewReplacer(
		"F", "┌",
		"L", "└",
		"7", "┐",
		"J", "┘",
		".", " ",
		"-", "─",
		"|", "│",
		"S", "█",
	)

	for y := 0; y < grid.Height; y++ {
		s := ""
		if grid.Data[y] != nil {
			s = replacer.Replace(string(grid.Data[y]))
		}
		fmt.Println(s)
	}
}
