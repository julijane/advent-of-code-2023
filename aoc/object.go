package aoc

type Object struct {
	field       *Field
	line        int
	left, right int
}

func (o *Object) String() string {
	return string(o.field.Data[o.line][o.left : o.right+1])
}

func (o *Object) Int() int {
	return Atoi(o.String())
}

func (o *Object) Adjacent(other *Object) bool {
	if other.right < o.left-1 || other.left > o.right+1 {
		return false
	}
	if other.line < o.line-1 || other.line > o.line+1 {
		return false
	}

	return true
}
