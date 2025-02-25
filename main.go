package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"marisalt/internal/game"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Marisalt: A Dark Fantasy Pirate Role Playing Game")

	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
