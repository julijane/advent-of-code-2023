package main

import (
	"slices"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

const (
	FIVE_OF_A_KIND  = 6
	FOUR_OF_A_KIND  = 5
	FULL_HOUSE      = 4
	THREE_OF_A_KIND = 3
	TWO_PAIR        = 2
	ONE_PAIR        = 1
	HIGH_CARD       = 0
)

type Hand struct {
	Cards string
	Bid   int
}

func (h *Hand) Type(part1 bool) int {
	counts := map[string]int{}
	for _, c := range h.Cards {
		counts[string(c)]++
	}

	var numJoker int
	if !part1 {
		numJoker = counts["J"]
	}

	// Five of a Kind
	if len(counts) == 1 {
		return FIVE_OF_A_KIND
	}

	if len(counts) == 2 {
		// could be 4+1 or 3+2

		if numJoker > 0 {
			// jokers turn into missing cards
			return FIVE_OF_A_KIND
		}

		// check for 4+1
		for _, v := range counts {
			if v == 4 {
				// Four of a Kind
				return FOUR_OF_A_KIND
			}
		}

		// it is 3+2
		return FULL_HOUSE
	}

	if len(counts) == 3 {
		// could be 2+2+1 or 3+1+1

		if numJoker > 1 {
			// could be either 2 or 3 jokers
			// in both cases we can combine with other card(s) to get 4 of a kind
			return FOUR_OF_A_KIND
		}

		if numJoker == 1 {
			// check for 3+1+1
			for _, v := range counts {
				if v == 3 {
					// Joker turns it into four of a kind
					return FOUR_OF_A_KIND
				}
			}

			// Joker turns it into Full House
			return FULL_HOUSE
		}

		// check for 3+1+1
		for _, v := range counts {
			if v == 3 {
				// Three of a Kind
				return THREE_OF_A_KIND
			}
		}

		// it is 2+2+1
		return TWO_PAIR
	}

	if len(counts) == 4 {
		// must be 2+1+1+1

		if numJoker > 0 {
			// Either 2 jokers + 1 other card or
			// 1 joker + 2 x same other card
			return THREE_OF_A_KIND
		}

		// One Pair
		return ONE_PAIR
	}

	if numJoker == 1 {
		// combine joker with one other card
		return ONE_PAIR
	}

	// High Card
	return HIGH_CARD
}

func (h *Hand) Compare(isPart1 bool, other *Hand) int {
	myType := h.Type(isPart1)
	otherType := other.Type(isPart1)

	if myType > otherType {
		return -1
	}
	if myType < otherType {
		return 1
	}

	var strengths string
	if isPart1 {
		strengths = "23456789TJQKA"
	} else {
		strengths = "J23456789TQKA"
	}

	for k := 0; k < len(h.Cards); k++ {
		iIndex := strings.Index(strengths, string(h.Cards[k]))
		jIndex := strings.Index(strengths, string(other.Cards[k]))

		if iIndex > jIndex {
			return -1
		}
		if iIndex < jIndex {
			return 1
		}
	}

	return 0
}

func scoreHands(hands []*Hand) int {
	result := 0
	for i, hand := range hands {
		result += hand.Bid * (len(hands) - i)
	}

	return result
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	hands := []*Hand{}
	for _, line := range input.Lines {
		splitted := strings.Split(line.Data, " ")
		hands = append(hands, &Hand{
			Cards: splitted[0],
			Bid:   aoc.Atoi(splitted[1]),
		})
	}

	// part 1
	slices.SortFunc(hands, func(i, j *Hand) int {
		return i.Compare(true, j)
	})
	resultPart1 := scoreHands(hands)

	// part 2
	slices.SortFunc(hands, func(i, j *Hand) int {
		return i.Compare(false, j)
	})
	resultPart2 := scoreHands(hands)

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
