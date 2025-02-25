package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"marisalt/internal/asset"
	"marisalt/internal/vec"
)

type Crosshair struct {
	WorldPos vec.Vector2
	sprite   *asset.Sprite
}

func NewCrosshair(assets *asset.Manager) *Crosshair {
	spriteImage, err := assets.LoadImage("ui/crosshair.png")
	if err != nil {
		panic(err)
	}

	sprite := asset.NewSprite(spriteImage, 16, 16, 2)
	sprite.AddAnimation("idle", 2, 0.5)

	return &Crosshair{
		sprite: sprite,
	}
}

func (c *Crosshair) Update(tileSize int, camera *Camera, dt float64) {
	c.sprite.Update(dt)
	mouseX, mouseY := ebiten.CursorPosition()

	worldX := float32(mouseX) + camera.Position.X
	worldY := float32(mouseY) + camera.Position.Y

	tileX := float32((int(worldX) / tileSize) * tileSize)
	tileY := float32((int(worldY) / tileSize) * tileSize)

	c.WorldPos = vec.Vector2{X: tileX, Y: tileY}
}

func (c *Crosshair) Draw(screen *ebiten.Image, camera *Camera) {
	mouseX, mouseY := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("(%d, %d)", mouseX, mouseY))

	screenPos := vec.Vector2{
		X: c.WorldPos.X - camera.Position.X,
		Y: c.WorldPos.Y - camera.Position.Y,
	}

	colorGreen := &ebiten.ColorScale{}
	colorGreen.Scale(0, 255, 0, 255)
	c.sprite.Draw(screen, screenPos, colorGreen)
}
