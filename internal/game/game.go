package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"marisalt/internal/asset"
	"marisalt/internal/vec"
)

type WorldCollider interface {
	IsPositionWalkable(vector2 vec.Vector2) bool
}

type Game struct {
	assets                *asset.Manager
	player                *Player
	camera                *Camera
	gridWidth, gridHeight int
	tileSize              int
	tileMap               [][]Tile
}

func (g *Game) IsPositionWalkable(vector2 vec.Vector2) bool {
	tileX := int(vector2.X) / g.tileSize
	tileY := int(vector2.Y) / g.tileSize

	if tileX < 0 || tileX >= g.gridWidth ||
		tileY < 0 || tileY >= g.gridHeight {
		return false
	}

	return !g.tileMap[tileY][tileX].Solid
}

func NewGame() *Game {
	g := &Game{
		assets:     asset.NewAssetManager(),
		camera:     NewCamera(),
		gridWidth:  20,
		gridHeight: 11,
		tileSize:   32,
		tileMap:    make([][]Tile, 11),
	}

	// Initialize tile map
	for y := 0; y < g.gridHeight; y++ {
		g.tileMap[y] = make([]Tile, 20)
		for x := 0; x < g.gridWidth; x++ {
			// Add walls around the edges
			if x == 0 || x == g.gridWidth-1 ||
				y == 0 || y == g.gridHeight-1 {
				g.tileMap[y][x] = Tile{
					Type:  TileWall,
					Solid: true,
				}
			}
		}
	}

	// Pass the game as a WorldCollider
	g.player = NewPlayer(g.assets, g)
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280 / 2, 720 / 2
}

func (g *Game) Update() error {
	dt := 1.0 / 60.0
	g.player.Update(dt)

	screenWidth, screenHeight := g.Layout(0, 0)
	g.camera.Update(g.player.pos, screenWidth, screenHeight)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!\n"+
		"Get ready to play Marisalt! >:)")

	for y := 0; y < g.gridHeight; y++ {
		for x := 0; x < g.gridWidth; x++ {
			tile := g.tileMap[y][x]
			x1 := float32(x*g.tileSize) - g.camera.Position.X
			y1 := float32(y*g.tileSize) - g.camera.Position.Y

			switch tile.Type {
			case TileWall:
				// Fill wall tiles with a solid color
				vector.DrawFilledRect(screen, x1, y1,
					float32(g.tileSize), float32(g.tileSize),
					color.RGBA{100, 100, 100, 255}, false)

				// Add border to walls
				vector.StrokeRect(screen, x1, y1,
					float32(g.tileSize), float32(g.tileSize),
					1, color.White, false)
			}
		}
	}

	g.player.Draw(screen, g.camera)
}

func (g *Game) IsSolidTileAt(vec2 vec.Vector2) bool {
	tileX := int(vec2.X) / g.tileSize
	tileY := int(vec2.Y) / g.tileSize

	if tileX < 0 || tileX >= g.gridWidth || tileY < 0 || tileY >= g.gridHeight {
		return true
	}

	return g.tileMap[tileY][tileX].Solid
}

func (g *Game) CheckCollision(x, y, width, height float64) bool {
	points := [][2]float64{
		{x, y},                  // Top-left
		{x + width, y},          // Top-right
		{x, y + height},         // Bottom-left
		{x + width, y + height}, // Bottom-right
	}

	for _, p := range points {
		if g.IsSolidTileAt(vec.Vector2{X: float32(p[0]), Y: float32(p[1])}) {
			return true
		}
	}

	return false
}
