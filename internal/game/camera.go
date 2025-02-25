package game

import "marisalt/internal/vec"

type Camera struct {
	Position    vec.Vector2
	targetPos   vec.Vector2
	smoothSpeed float32
}

func NewCamera() *Camera {
	return &Camera{
		Position:    vec.New(0, 0),
		targetPos:   vec.New(0, 0),
		smoothSpeed: 0.1,
	}
}

func (c *Camera) Update(target vec.Vector2, screenWidth, screenHeight int, tileSize int) {
	centerX := float32(screenWidth) / 2
	centerY := float32(screenHeight) / 2

	rawX := target.X - centerX
	rawY := target.Y - centerY

	c.targetPos.X = float32((int(rawX) / tileSize) * tileSize)
	c.targetPos.Y = float32((int(rawY) / tileSize) * tileSize)

	c.Position.X += (c.targetPos.X - c.Position.X) * c.smoothSpeed
	c.Position.Y += (c.targetPos.Y - c.Position.Y) * c.smoothSpeed
}
