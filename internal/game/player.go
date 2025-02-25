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
	sprite        *asset.Sprite
}

func NewPlayer(assets *asset.Manager) *Player {
	spriteImage, err := assets.LoadImage("human-8px.png")
	if err != nil {
		panic(err)
	}

	sprite := asset.NewSprite(spriteImage, 8, 8, 4)
	sprite.AddAnimation("idle", 2, 0.5)

	return &Player{
		pos:    vec.New(100, 100),
		speed:  8,
		width:  32,
		height: 32,
		sprite: sprite,
	}
}

func (p *Player) Update(dt float64) {
	p.handleInput()
	p.sprite.Update(dt)
}

func (p *Player) handleInput() {
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
	p.sprite.Draw(screen, p.pos)
}
