package aoc

import (
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
