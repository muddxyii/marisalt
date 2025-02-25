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

func (v Vector2) Normalized() Vector2 {
	magnitude := v.Length()
	if magnitude != 0 {
		v.X /= magnitude
		v.Y /= magnitude
	}

	return v
}

func (v Vector2) Add(other Vector2) Vector2 {
	v.X += other.X
	v.Y += other.Y

	return v
}

func (v Vector2) Sub(other Vector2) Vector2 {
	v.X -= other.X
	v.Y -= other.Y

	return v
}

func (v Vector2) Mul(scalar float32) Vector2 {
	v.X *= scalar
	v.Y *= scalar

	return v
}

func (v Vector2) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}
