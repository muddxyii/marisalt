package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"marisalt/internal/asset"
	"marisalt/internal/vec"
)

type Player struct {
	pos           vec.Vector2
	speed         float32
	width, height float32
	sprite        *ebiten.Image
}

func NewPlayer(assets *asset.Manager) *Player {
	sprite, err := assets.LoadImage("human-8px.png")
	if err != nil {
		panic(err)
	}

	return &Player{
		pos:    vec.New(100, 100),
		speed:  8,
		width:  32,
		height: 32,
		sprite: sprite,
	}
}

func (p *Player) HandleInput() {
	vel := vec.Vector2Zero()

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		vel.Y -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		vel.Y += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		vel.X += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		vel.X -= p.speed
	}

	vel.Normalize()
	p.pos.Add(vel)
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.pos.X), float64(p.pos.Y))

	screen.DrawImage(p.sprite, op)
}
