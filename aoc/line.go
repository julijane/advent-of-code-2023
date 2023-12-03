package aoc

import (
	"regexp"
	"strings"
)

type Line struct {
	field  *Field
	lineNo int

	Text string
}

// FindObjects returns a list of objects that match the given regular expression
func (l *Line) FindObjects(re string) []*Object {
	var objects []*Object
	matcher := regexp.MustCompile(re)

	matches := matcher.FindAllStringIndex(l.Text, -1)
	for _, match := range matches {
		objects = append(objects, &Object{
			field: l.field,
			line:  l.lineNo,
			left:  match[0],
			right: match[1] - 1,
		})
	}

	return objects
}

func (l *Line) ReplaceText(find, replace string) {
	newText := strings.ReplaceAll(l.Text, find, replace)
	l.field.Data[l.lineNo] = []byte(newText)
	l.Text = newText
}
