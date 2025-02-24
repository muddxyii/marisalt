package vec

import "math"

type Vector2 struct {
	X, Y float32
}

func New(x, y float32) Vector2 {
	return Vector2{X: x, Y: y}
}

func Vector2Zero() Vector2 {
	return Vector2{X: 0, Y: 0}
}

func (v *Vector2) Normalize() {
	magnitude := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
	if magnitude != 0 {
		v.X /= magnitude
		v.Y /= magnitude
	}
}

func (v *Vector2) Add(other Vector2) {
	v.X += other.X
	v.Y += other.Y
}
