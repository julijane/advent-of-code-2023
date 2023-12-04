package aoc

import (
	"regexp"
	"strconv"
)

// SliceMemberOrEmptyString returns the member of a slice at the given index,
// or an empty string if the index is out of bounds
func SliceMemberOrEmptyString(slice []string, index int) string {
	if index < len(slice) {
		return slice[index]
	}
	return ""
}

// Atoi converts a string to an int, ignoring errors (return zero instead)
func Atoi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

// ExtractNumbers extracts all numbers from a string
func ExtractNumbers(s string) []int {
	var res []int

	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)
	for _, match := range matches {
		res = append(res, Atoi(match))
	}

	return res
}

func RegexpSubmatchAsInt(s, expr string) int {
	re := regexp.MustCompile(expr)
	match := re.FindStringSubmatch(s)
	return Atoi(SliceMemberOrEmptyString(match, 1))
}
