package aoc

import "strings"

// Input represents the input data
type Input struct {
	Lines []*Line
}

// NewInput creates a new Input object from the given input lines
func NewInput(inputLines []string) *Input {
	i := &Input{}

	for lineNo, line := range inputLines {
		i.Lines = append(i.Lines, &Line{
			LineNo: lineNo,
			Data:   line,
		})
	}

	return i
}

// FindObjects returns a list of objects that match the given regular expression
func (i *Input) FindObjects(re string) []*Object {
	var objects []*Object

	for _, line := range i.Lines {
		lineObjects := line.FindObjects(re)
		objects = append(objects, lineObjects...)
	}

	return objects
}

// TextBlocks returns the input data as a list of blocks (separated by empty lines in input)
func (i *Input) TextBlocks() [][]string {
	var blocks [][]string
	var block []string

	for _, line := range i.Lines {
		if line.Data == "" {
			blocks = append(blocks, block)
			block = []string{}
			continue
		}
		block = append(block, line.Data)
	}
	blocks = append(blocks, block)

	return blocks
}

func (i *Input) PlainLines() []string {
	res := []string{}
	for _, line := range i.Lines {
		res = append(res, line.Data)
	}

	return res
}

func (i *Input) SingleString() string {
	return strings.Join(i.PlainLines(), "\n")
}

func (i *Input) Map(fn func(string) string) {
	for _, line := range i.Lines {
		line.Data = fn(line.Data)
	}
}
