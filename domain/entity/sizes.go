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

// Contains returns that point in sizes diapason.
// p.X in [0, width) and p.Y in [0, width)
func (s Sizes) Contains(p Point) bool {
	return notGreat(p.X, s.x) && notGreat(p.Y, s.y)
}

// notGreat returns that value belongs to the range [0, b)
func notGreat(value int, b int) bool {
	return value >= 0 && value < b
}
