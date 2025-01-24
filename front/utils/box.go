package utils

import "math"

type Box struct {
	X1, Y1, X2, Y2 float64
}

func NewBox(x1, y1, x2, y2 float64) Box {
	return Box{
		X1: math.Min(x1, x2),
		Y1: math.Min(y1, y2),
		X2: math.Max(x1, x2),
		Y2: math.Max(y1, y2),
	}
}

func (b *Box) IsWithin(v Vector) bool {
	return v.X >= b.X1 && v.X <= b.X2 && v.Y >= b.Y1 && v.Y <= b.Y2
}
