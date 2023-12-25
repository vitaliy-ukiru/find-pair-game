package entity

type Point struct {
	X, Y int
}

func NewPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}
