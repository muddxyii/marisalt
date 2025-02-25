package game

import "marisalt/internal/vec"

type TileType int

const (
	TileEmpty TileType = iota
	TileGround
	TileWall
)

type Tile struct {
	Type     TileType
	Solid    bool
	Position vec.Vector2
}
