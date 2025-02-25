package game

import "marisalt/internal/vec"

type Camera struct {
	Position vec.Vector2
}

func NewCamera() *Camera {
	return &Camera{
		Position: vec.New(0, 0),
	}
}

func (c *Camera) Update(target vec.Vector2, screenWidth, screenHeight int) {
	c.Position.X = target.X - float32(screenWidth)/2
	c.Position.Y = target.Y - float32(screenHeight)/2
}
