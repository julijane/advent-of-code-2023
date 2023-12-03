package aoc

type Object struct {
	Line        *Line
	left, right int
}

func (o *Object) String() string {
	return o.Line.Data[o.left : o.right+1]
}

func (o *Object) Int() int {
	return Atoi(o.String())
}

func (o *Object) Adjacent(other *Object) bool {
	if other.right < o.left-1 || other.left > o.right+1 {
		return false
	}
	if other.Line.LineNo < o.Line.LineNo-1 || other.Line.LineNo > o.Line.LineNo+1 {
		return false
	}

	return true
}
