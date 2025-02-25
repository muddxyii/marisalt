package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"marisalt/internal/asset"
)

type Game struct {
	assets *asset.Manager
	player *Player
}

func NewGame() *Game {
	assets := asset.NewAssetManager()
	return &Game{
		assets: assets,
		player: NewPlayer(assets),
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

	g.player.Draw(screen)
}
