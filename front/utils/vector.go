package utils

type Vector struct {
	X, Y float64
}

func (v *Vector) Set(x, y float64) {
	v.X = x
	v.Y = y
}
