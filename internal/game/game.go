package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"marisalt/internal/asset"
)

type Game struct {
	assets                *asset.Manager
	player                *Player
	gridWidth, gridHeight int
	tileSize              int
}

func NewGame() *Game {
	assets := asset.NewAssetManager()
	return &Game{
		assets:     assets,
		player:     NewPlayer(assets),
		gridWidth:  20,
		gridHeight: 11,
		tileSize:   32,
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280 / 2, 720 / 2
}

func (g *Game) Update() error {
	dt := 1.0 / 60.0
	g.player.Update(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!\n"+
		"Get ready to play Marisalt! >:)")

	for y := 0; y < g.gridHeight; y++ {
		for x := 0; x < g.gridWidth; x++ {
			x1 := float32(x * g.tileSize)
			y1 := float32(y * g.tileSize)
			x2 := float32((x + 1) * g.tileSize)
			y2 := float32((y + 1) * g.tileSize)

			vector.StrokeLine(screen, x1, y1, x2, y1, 1, color.White, false)
			vector.StrokeLine(screen, x1, y1, x1, y2, 1, color.White, false)
			vector.StrokeLine(screen, x2, y1, x2, y2, 1, color.White, false)
			vector.StrokeLine(screen, x1, y2, x2, y2, 1, color.White, false)
		}
	}

	g.player.Draw(screen)
}
