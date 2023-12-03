package aoc

import (
	"regexp"
	"strings"
)

// Line represents a line of input
type Line struct {
	LineNo int
	Data   string
}

// FindObjects returns a list of objects that match the given regular expression on this Line
func (l *Line) FindObjects(re string) []*Object {
	var objects []*Object
	matcher := regexp.MustCompile(re)

	matches := matcher.FindAllStringIndex(l.Data, -1)
	for _, match := range matches {
		objects = append(objects, &Object{
			Line:  l,
			left:  match[0],
			right: match[1] - 1,
		})
	}

	return objects
}

// ReplaceText replaces all occurrences of the given string with the given replacement
// on this Line. Note: length of find and replace must be the same
func (l *Line) ReplaceText(find, replace string) {
	if len(find) != len(replace) {
		return
	}
	l.Data = strings.ReplaceAll(l.Data, find, replace)
}
