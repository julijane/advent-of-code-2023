package aoc

type Field struct {
	Lines []*Line
}

type FieldCalcFunction func(f *Field) (int, int)

func NewField(inputLines []string) *Field {
	f := &Field{}

	for lineNo, line := range inputLines {
		f.Lines = append(f.Lines, &Line{
			LineNo: lineNo,
			Data:   line,
		})
	}

	return f
}

// FindObjects returns a list of objects that match the given regular expression
func (f *Field) FindObjects(re string) []*Object {
	var objects []*Object

	for _, line := range f.Lines {
		lineObjects := line.FindObjects(re)
		objects = append(objects, lineObjects...)
	}

	return objects
}
