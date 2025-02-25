package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"marisalt/internal/asset"
	"marisalt/internal/vec"
)

type Player struct {
	pos           vec.Vector2
	targetPos     vec.Vector2
	speed         float32
	width, height float32
	sprite        *asset.Sprite
	isMoving      bool
	tileSize      float32
}

func NewPlayer(assets *asset.Manager) *Player {
	spriteImage, err := assets.LoadImage("human-8px.png")
	if err != nil {
		panic(err)
	}

	sprite := asset.NewSprite(spriteImage, 8, 8, 4)
	sprite.AddAnimation("idle", 2, 0.5)

	return &Player{
		pos:       vec.New(64, 64),
		targetPos: vec.New(64, 64),
		speed:     8,
		width:     32,
		height:    32,
		sprite:    sprite,
		isMoving:  false,
		tileSize:  32,
	}
}

func (p *Player) handleInput() {
	if p.isMoving {
		return
	}

	var moveDir vec.Vector2
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		moveDir = vec.New(0, -p.tileSize)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		moveDir = vec.New(0, p.tileSize)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		moveDir = vec.New(p.tileSize, 0)
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		moveDir = vec.New(-p.tileSize, 0)
	}

	if moveDir.X != 0 || moveDir.Y != 0 {
		p.targetPos = p.pos.Add(moveDir)
		p.isMoving = true
	}
}

func (p *Player) Update(dt float64) {
	p.handleInput()

	// If we're moving, interpolate towards the target position
	if p.isMoving {
		moveVector := p.targetPos.Sub(p.pos)
		distance := moveVector.Length()

		if distance < 1 { // We've basically reached the target
			p.pos = p.targetPos
			p.isMoving = false
		} else {
			// Move towards target
			moveDir := moveVector.Normalized()
			movement := moveDir.Mul(p.speed)
			p.pos = p.pos.Add(movement)
		}
	}

	p.sprite.Update(dt)
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.Draw(screen, p.pos)
}
