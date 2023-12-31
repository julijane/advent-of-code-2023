package aoc

type Coordinate struct {
	X, Y int
}

func (c Coordinate) Same(other Coordinate) bool {
	return c.X == other.X && c.Y == other.Y
}

func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X + other.X,
		Y: c.Y + other.Y,
	}
}

func (c Coordinate) AddXY(x, y int) Coordinate {
	return Coordinate{
		X: c.X + x,
		Y: c.Y + y,
	}
}

func (c Coordinate) Copy() Coordinate {
	return Coordinate{
		X: c.X,
		Y: c.Y,
	}
}

func (c Coordinate) Above() Coordinate {
	return c.AddXY(0, -1)
}

func (c Coordinate) Below() Coordinate {
	return c.AddXY(0, 1)
}

func (c Coordinate) Left() Coordinate {
	return c.AddXY(-1, 0)
}

func (c Coordinate) Right() Coordinate {
	return c.AddXY(1, 0)
}

func (c Coordinate) Move(direction int) Coordinate {
	switch direction {
	case 0:
		return c.Above()
	case 1:
		return c.Right()
	case 2:
		return c.Below()
	case 3:
		return c.Left()
	}
	return c
}

func (c Coordinate) MoveBy(direction int, amount int) Coordinate {
	switch direction {
	case 0:
		return c.AddXY(0, -amount)
	case 1:
		return c.AddXY(amount, 0)
	case 2:
		return c.AddXY(0, amount)
	case 3:
		return c.AddXY(-amount, 0)
	}
	return c
}

type Coordinates []Coordinate

func (cs Coordinates) Includes(c Coordinate) bool {
	for _, test := range cs {
		if test.X == c.X && test.Y == c.Y {
			return true
		}
	}
	return false
}
