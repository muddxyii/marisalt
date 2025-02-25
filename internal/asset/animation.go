package asset

type Animation struct {
	frameCount   int
	CurrentFrame int
	frameTime    float64
	elapsed      float64
}

func NewAnimation(frameCount int, frameTime float64) *Animation {
	return &Animation{
		frameCount:   frameCount,
		frameTime:    frameTime,
		CurrentFrame: 0,
		elapsed:      0,
	}
}

func (a *Animation) Update(dt float64) {
	a.elapsed += dt
	if a.elapsed >= a.frameTime {
		a.CurrentFrame = (a.CurrentFrame + 1) % a.frameCount
		a.elapsed = 0
	}
}
