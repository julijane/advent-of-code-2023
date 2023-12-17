package aoc

type Grid struct {
	Width  int
	Height int
	Data   [][]byte
}

func (i *Input) Grid() *Grid {
	grid := &Grid{
		Width:  len(i.Lines[0].Data),
		Height: len(i.Lines),
	}

	for _, line := range i.Lines {
		grid.Data = append(grid.Data, []byte(line.Data))
	}

	return grid
}

func (g *Grid) Inside(c Coordinate) bool {
	return c.X >= 0 && c.X < g.Width && c.Y >= 0 && c.Y < g.Height
}

func (g *Grid) Get(c Coordinate, outsideVal byte) byte {
	if !g.Inside(c) {
		return outsideVal
	}

	return g.Data[c.Y][c.X]
}

func (g *Grid) GetInt(c Coordinate, outsideVal int) int {
	if !g.Inside(c) {
		return outsideVal
	}

	return Atoi(string(g.Data[c.Y][c.X]))
}

func (g *Grid) Set(c Coordinate, val byte) {
	if g.Inside(c) {
		g.Data[c.Y][c.X] = val
	}
}

func (g *Grid) Find(search byte) Coordinate {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.Data[y][x] == search {
				return Coordinate{X: x, Y: y}
			}
		}
	}

	return Coordinate{X: -1, Y: -1}
}

type GridMapFunction func(pos Coordinate, value byte) byte

func (g *Grid) Map(fn GridMapFunction) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.Data[y][x] = fn(Coordinate{
				X: x,
				Y: y,
			}, g.Data[y][x])
		}
	}
}

func (g *Grid) FindConnectedFrom(startPos Coordinate, foundBefore Coordinates, search byte) *Coordinates {
	found := Coordinates{}

	if g.Get(startPos, '#') != search {
		return &found
	}

	foundBefore = append(foundBefore, startPos)

	offsets := [4][2]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}

	for _, offset := range offsets {
		pos := startPos.AddXY(offset[0], offset[1])
		if !foundBefore.Includes(pos) {
			findNeighbors := g.FindConnectedFrom(pos, foundBefore, search)
			found = append(found, (*findNeighbors)...)
		}
	}

	return &found
}
