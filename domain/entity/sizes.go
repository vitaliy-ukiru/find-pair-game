package entity

type Sizes struct {
	x, y int
}

func NewSizes(x int, y int) Sizes {
	return Sizes{x: x, y: y}
}

func (s Sizes) CellsCount() int {
	return s.x * s.y
}

func (s Sizes) Width() int  { return s.x }
func (s Sizes) Height() int { return s.y }
