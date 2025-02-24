package asset

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"os"
	"path/filepath"
	"runtime"
)

type Manager struct {
	basePath string
}

func NewAssetManager() *Manager {
	// Get the directory of the current file
	_, filename, _, _ := runtime.Caller(0)
	// Navigate to project root/assets directory
	basePath := filepath.Join(filepath.Dir(filename), "..", "..", "assets")

	return &Manager{
		basePath: basePath,
	}
}

func (am *Manager) LoadImage(path string) (*ebiten.Image, error) {
	fullPath := filepath.Join(am.basePath, path)

	f, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image %s: %w", path, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image %s: %w", path, err)
	}

	return ebiten.NewImageFromImage(img), nil
}
