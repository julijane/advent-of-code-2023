package aoc

// Input represents the input data
type Input struct {
	Lines []*Line
}

// InputCalcFunction is the function signature for the calculation function
type InputCalcFunction func(i *Input) (int, int)

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
