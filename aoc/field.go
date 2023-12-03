package aoc

type Field struct {
	Data [][]byte
}

type FieldCalcFunction func(f *Field) (int, int)

func (f *Field) Lines() []*Line {
	var lines []*Line
	for lineNo, line := range f.Data {
		lines = append(lines, &Line{
			Text: string(line),

			field:  f,
			lineNo: lineNo,
		})
	}
	return lines
}

// FindObjects returns a list of objects that match the given regular expression
func (f *Field) FindObjects(re string) []*Object {
	var objects []*Object

	for _, line := range f.Lines() {
		lineObjects := line.FindObjects(re)
		objects = append(objects, lineObjects...)
	}

	return objects
}
