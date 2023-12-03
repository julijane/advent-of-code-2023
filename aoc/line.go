package aoc

import (
	"regexp"
	"strings"
)

type Line struct {
	Field  *Field
	LineNo int
}

func (l *Line) String() string {
	return string(l.Field.Data[l.LineNo])
}

// FindObjects returns a list of objects that match the given regular expression
func (l *Line) FindObjects(re string) []*Object {
	var objects []*Object
	matcher := regexp.MustCompile(re)

	matches := matcher.FindAllStringIndex(l.String(), -1)
	for _, match := range matches {
		objects = append(objects, &Object{
			Line:  l,
			left:  match[0],
			right: match[1] - 1,
		})
	}

	return objects
}

func (l *Line) ReplaceText(find, replace string) {
	newText := strings.ReplaceAll(l.String(), find, replace)
	l.Field.Data[l.LineNo] = []byte(newText)
}
