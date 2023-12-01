package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/julijane/advent-of-code-2023/internal/aoc"
)

func calc(input []string) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	regBlue := regexp.MustCompile("([0-9]+) blue")
	regRed := regexp.MustCompile("([0-9]+) red")
	regGreen := regexp.MustCompile("([0-9]+) green")

	for game, line := range input {
		split1 := strings.Split(line, ": ")
		split2 := strings.Split(split1[1], ";")

		gamePossible := true
		needRed := 0
		needGreen := 0
		needBlue := 0

		for _, draw := range split2 {
			matchRed := regRed.FindSubmatch([]byte(draw))
			matchGreen := regGreen.FindSubmatch([]byte(draw))
			matchBlue := regBlue.FindSubmatch([]byte(draw))

			red := 0
			green := 0
			blue := 0

			if len(matchRed) == 2 {
				red, _ = strconv.Atoi(string(matchRed[1]))
			}
			if len(matchGreen) == 2 {
				green, _ = strconv.Atoi(string(matchGreen[1]))
			}
			if len(matchBlue) == 2 {
				blue, _ = strconv.Atoi(string(matchBlue[1]))
			}

			if red > 12 || green > 13 || blue > 14 {
				gamePossible = false
			}

			if red > needRed {
				needRed = red
			}
			if green > needGreen {
				needGreen = green
			}
			if blue > needBlue {
				needBlue = blue
			}
		}

		if gamePossible {
			sumPart1 += game + 1
		}

		sumPart2 += (needRed * needGreen * needBlue)
	}
	return sumPart1, sumPart2
}

func main() {
	s1, s2 := calc(aoc.ReadFileAsLines("sample1.txt"))
	fmt.Println("Demo:", s1, s2)

	s1, s2 = calc(aoc.ReadFileAsLines("input.txt"))
	fmt.Println("Main:", s1, s2)
}
