package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	collider      WorldCollider
}

func NewPlayer(assets *asset.Manager, collider WorldCollider) *Player {
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
		collider:  collider,
	}
}

func (p *Player) canMoveTo(newPos vec.Vector2) bool {
	return p.collider.IsPositionWalkable(newPos)
}

func (p *Player) handleInput() {
	if p.isMoving {
		return
	}

	var moveDir vec.Vector2
	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		moveDir = vec.New(0, -p.tileSize)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		moveDir = vec.New(0, p.tileSize)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		moveDir = vec.New(p.tileSize, 0)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		moveDir = vec.New(-p.tileSize, 0)
	}

	if moveDir.X != 0 || moveDir.Y != 0 {
		newPos := p.pos.Add(moveDir)
		if p.canMoveTo(newPos) {
			p.targetPos = newPos
			p.isMoving = true
		}
	}

}

func (p *Player) Update(dt float64) {
	p.handleInput()

	if p.isMoving {
		moveVector := p.targetPos.Sub(p.pos)
		distance := moveVector.Length()

		if distance < 1 {
			p.pos = p.targetPos
			p.isMoving = false
		} else {
			moveDir := moveVector.Normalized()
			movement := moveDir.Mul(p.speed)
			newPos := p.pos.Add(movement)

			if p.canMoveTo(newPos) {
				p.pos = newPos
			} else {
				p.pos = p.targetPos
				p.isMoving = false
			}

		}
	}

	p.sprite.Update(dt)
}

func (p *Player) GetBounds() (float32, float32, float32, float32) {
	return p.pos.X, p.pos.Y, p.width, p.height
}

func (p *Player) Draw(screen *ebiten.Image, camera *Camera) {
	spritePos := p.pos.Sub(camera.Position)
	p.sprite.Draw(screen, spritePos)
}
