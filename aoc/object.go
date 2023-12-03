package aoc

// Object represents an object found on a line
type Object struct {
	Line        *Line
	left, right int
}

// String returns the object as a string
func (o *Object) String() string {
	return o.Line.Data[o.left : o.right+1]
}

// Int returns the object as an int
func (o *Object) Int() int {
	return Atoi(o.String())
}

// Adjacent returns true if the given object is adjacent to this object
func (o *Object) Adjacent(other *Object) bool {
	if other.right < o.left-1 || other.left > o.right+1 {
		return false
	}

	if other.Line.LineNo < o.Line.LineNo-1 || other.Line.LineNo > o.Line.LineNo+1 {
		return false
	}

	return true
}
