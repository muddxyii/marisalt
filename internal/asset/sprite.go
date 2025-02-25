package asset

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"marisalt/internal/vec"
)

type Sprite struct {
	image                   *ebiten.Image
	frameWidth, frameHeight int
	scale                   float64
	animations              map[string]*Animation
	currentAnim             string
}

func NewSprite(image *ebiten.Image, frameWidth, frameHeight int, scale float64) *Sprite {
	return &Sprite{
		image:       image,
		frameWidth:  frameWidth,
		frameHeight: frameHeight,
		scale:       scale,
		animations:  make(map[string]*Animation),
		currentAnim: "",
	}
}

func (s *Sprite) AddAnimation(name string, frameCount int, frameTime float64) {
	s.animations[name] = NewAnimation(frameCount, frameTime)
	if s.currentAnim == "" {
		s.currentAnim = name
	}
}

func (s *Sprite) PlayAnimation(name string) {
	if _, exists := s.animations[name]; exists {
		s.currentAnim = name
	}
}

func (s *Sprite) Update(dt float64) {
	if anim, exists := s.animations[s.currentAnim]; exists {
		anim.Update(dt)
	}
}

func (s *Sprite) Draw(screen *ebiten.Image, pos vec.Vector2, color *ebiten.ColorScale) {
	anim, exists := s.animations[s.currentAnim]
	if !exists {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s.scale, s.scale)
	op.GeoM.Translate(float64(pos.X), float64(pos.Y))

	// Apply color if provided
	if color != nil {
		op.ColorScale = *color
	}

	sx := anim.CurrentFrame * s.frameWidth
	sy := 0

	r := image.Rectangle{
		Min: image.Point{X: sx, Y: sy},
		Max: image.Point{X: sx + s.frameWidth, Y: sy + s.frameHeight},
	}

	screen.DrawImage(s.image.SubImage(r).(*ebiten.Image), op)
}
