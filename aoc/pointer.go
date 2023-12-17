package aoc

import "fmt"

type Pointer struct {
	C   Coordinate
	Dir int
}

func (p *Pointer) Move() {
	p.C = p.C.Move(p.Dir)
}

func (p *Pointer) TurnRight() {
	p.Dir = (p.Dir + 1) % 4
}

func (p *Pointer) TurnLeft() {
	p.Dir = (p.Dir + 3) % 4
}

func (p *Pointer) TurnAround() {
	p.Dir = (p.Dir + 2) % 4
}

func (p *Pointer) IsUpOrDown() bool {
	return p.Dir == 0 || p.Dir == 2
}

func (p *Pointer) IsLeftOrRight() bool {
	return p.Dir == 1 || p.Dir == 3
}

func (p *Pointer) String() string {
	dirvis := [4]string{"🠉", "🠊", "🠋", "🠈"}
	return fmt.Sprintf("%d:%d%s", p.C.X, p.C.Y, dirvis[p.Dir])
}

func NewPointer(x, y, dir int) Pointer {
	return Pointer{
		C:   Coordinate{X: x, Y: y},
		Dir: dir,
	}
}
