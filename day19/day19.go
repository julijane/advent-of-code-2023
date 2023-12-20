package main

import (
	"regexp"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

type ratingrange struct {
	from, to int
}

func (rr *ratingrange) len() int {
	return rr.to - rr.from + 1
}

type ratingranges struct {
	categories [4]ratingrange
	workflow   string
}

type part [4]int // x, m, a, s

type rule struct {
	category     int    // x, m, a, s
	op           byte   // <, >
	compare      int    // n
	nextWorkflow string // workflow, in, R or A
}

type workflow struct {
	rules     []rule
	otherwise string
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	blocks := input.TextBlocks()

	re := regexp.MustCompile(`(.+){(.*),([^,]+)}`)
	re2 := regexp.MustCompile(`([a-z]+)([<>])(\d+):([a-zAR]+)`)

	workflows := make(map[string]workflow)

	matches := re.FindAllStringSubmatch(input.SingleString(), -1)
	for _, match := range matches {
		newWorkflow := workflow{
			otherwise: match[3],
		}

		for _, rulespec := range strings.Split(match[2], ",") {
			ruleparts := re2.FindStringSubmatch(rulespec)
			newWorkflow.rules = append(newWorkflow.rules, rule{
				category:     strings.Index("xmas", ruleparts[1]),
				op:           ruleparts[2][0],
				compare:      aoc.Atoi(ruleparts[3]),
				nextWorkflow: ruleparts[4],
			})
		}

		workflows[match[1]] = newWorkflow
	}

	// part 1

	sumPart1 := 0
	for _, line := range blocks[1] {
		nums := aoc.ExtractNumbers(line)
		part := [4]int{nums[0], nums[1], nums[2], nums[3]}
		currentWorkflow := "in"

	sequence:
		for currentWorkflow != "R" && currentWorkflow != "A" {
			for _, rule := range workflows[currentWorkflow].rules {
				if rule.op == '<' && part[rule.category] < rule.compare {
					currentWorkflow = rule.nextWorkflow
					continue sequence
				} else if rule.op == '>' && part[rule.category] > rule.compare {
					currentWorkflow = rule.nextWorkflow
					continue sequence
				}
			}

			currentWorkflow = workflows[currentWorkflow].otherwise
		}

		if currentWorkflow == "A" {
			sumPart1 += part[0] + part[1] + part[2] + part[3]
		}
	}

	// part 2
	sumPart2 := 0
	ranges := []ratingranges{
		{
			categories: [4]ratingrange{
				{1, 4000},
				{1, 4000},
				{1, 4000},
				{1, 4000},
			},
			workflow: "in",
		},
	}

	for len(ranges) > 0 {
		newRanges := []ratingranges{}
		for _, rangeIn := range ranges {
			if rangeIn.workflow == "R" || rangeIn.workflow == "A" {
				if rangeIn.workflow == "A" {
					sumPart2 += rangeIn.categories[0].len() *
						rangeIn.categories[1].len() *
						rangeIn.categories[2].len() *
						rangeIn.categories[3].len()
				}
				continue
			}

			for _, rule := range workflows[rangeIn.workflow].rules {
				rangeNew := ratingranges{
					categories: rangeIn.categories,
					workflow:   rule.nextWorkflow,
				}

				if rule.op == '<' {
					rangeNew.categories[rule.category].to = rule.compare - 1
					rangeIn.categories[rule.category].from = rule.compare
				} else {
					rangeNew.categories[rule.category].from = rule.compare + 1
					rangeIn.categories[rule.category].to = rule.compare
				}

				newRanges = append(newRanges, rangeNew)
			}

			rangeIn.workflow = workflows[rangeIn.workflow].otherwise
			newRanges = append(newRanges, rangeIn)
		}

		ranges = newRanges
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
