package aoc

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/k0kubun/pp/v3"
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

func ExtractRegexps(s, expr string) []string {
	re := regexp.MustCompile(expr)
	return re.FindAllString(s, -1)
}

// ExtractNumbers extracts all numbers from a string
func ExtractNumbers(s string) []int {
	var res []int

	for _, match := range ExtractRegexps(s, `-?\d+`) {
		res = append(res, Atoi(match))
	}
	return res
}

func ExtractDigits(s string) []int {
	var res []int

	for _, match := range ExtractRegexps(s, `\d`) {
		res = append(res, Atoi(match))
	}
	return res
}

func RegexpSubmatchAsInt(s, expr string) int {
	re := regexp.MustCompile(expr)
	match := re.FindStringSubmatch(s)
	return Atoi(SliceMemberOrEmptyString(match, 1))
}

func Print(s string, v any) {
	fmt.Print(s + ": ")
	pp.Print(v)
	fmt.Println()
}

func StringFrom(s string, from int) string {
	if len(s) < from {
		return ""
	}

	return s[from:]
}

func SliceFrom[T any](s []T, from int) []T {
	if len(s) < from {
		return []T{}
	}

	return s[from:]
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func New2DSlice[T any](rows, cols int) [][]T {
	s := make([][]T, rows)
	for i := range s {
		s[i] = make([]T, cols)
	}
	return s
}
