package main

import (
	"regexp"
	"strings"

	"github.com/julijane/advent-of-code-2023/aoc"
)

type numrange struct {
	from, to int
}

type rangeMapping struct {
	ranges [4]numrange
	target string
}

type part [4]int // x, m, a, s

type rule struct {
	which   int    // x, m, a, s
	op      byte   // <, >
	compare int    // n
	target  string // workflow, in, R or A
}

type ruleset struct {
	rules     []rule
	otherwise string
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	blocks := input.TextBlocks()

	mappings := make(map[string]ruleset)
	re := regexp.MustCompile(`([a-z]+){(.*)}`)
	re2 := regexp.MustCompile(`([a-z]+)([<>])(\d+):([a-zAR]+)`)

	for _, line := range blocks[0] {
		matches := re.FindStringSubmatch(line)
		ruleparts := strings.Split(matches[2], ",")

		rs := ruleset{
			otherwise: ruleparts[len(ruleparts)-1],
		}

		for _, rulepart := range ruleparts[:len(ruleparts)-1] {
			rp := re2.FindStringSubmatch(rulepart)
			rule := rule{
				which:   strings.Index("xmas", rp[1]),
				op:      rp[2][0],
				compare: aoc.Atoi(rp[3]),
				target:  rp[4],
			}

			rs.rules = append(rs.rules, rule)
		}

		mappings[matches[1]] = rs
	}

	// part 1

	sumPart1 := 0
	for _, line := range blocks[1] {
		nums := aoc.ExtractNumbers(line)
		part := [4]int{nums[0], nums[1], nums[2], nums[3]}
		partBin := "in"

	sequence:
		for partBin != "R" && partBin != "A" {
			for _, rule := range mappings[partBin].rules {
				if rule.op == '<' && part[rule.which] < rule.compare {
					partBin = rule.target
					continue sequence
				} else if rule.op == '>' && part[rule.which] > rule.compare {
					partBin = rule.target
					continue sequence
				}
			}

			partBin = mappings[partBin].otherwise
		}

		if partBin == "A" {
			sumPart1 += part[0] + part[1] + part[2] + part[3]
		}
	}

	// part 2
	sumPart2 := 0
	rangeMappings := []rangeMapping{
		{
			ranges: [4]numrange{
				{1, 4000},
				{1, 4000},
				{1, 4000},
				{1, 4000},
			},
			target: "in",
		},
	}

	for len(rangeMappings) > 0 {
		newRangeMappings := []rangeMapping{}
		for _, rm := range rangeMappings {
			if rm.target == "R" || rm.target == "A" {
				if rm.target == "A" {
					sumPart2 += (rm.ranges[0].to - rm.ranges[0].from + 1) *
						(rm.ranges[1].to - rm.ranges[1].from + 1) *
						(rm.ranges[2].to - rm.ranges[2].from + 1) *
						(rm.ranges[3].to - rm.ranges[3].from + 1)
				}
				continue
			}

			rules := mappings[rm.target].rules
			for _, rule := range rules {
				newrm := rangeMapping{
					ranges: rm.ranges,
					target: rule.target,
				}

				if rule.op == '<' {
					newrm.ranges[rule.which].to = rule.compare - 1
					rm.ranges[rule.which].from = rule.compare
				} else {
					newrm.ranges[rule.which].from = rule.compare + 1
					rm.ranges[rule.which].to = rule.compare
				}

				newRangeMappings = append(newRangeMappings, newrm)
			}

			rm.target = mappings[rm.target].otherwise
			newRangeMappings = append(newRangeMappings, rm)
		}

		rangeMappings = newRangeMappings
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
