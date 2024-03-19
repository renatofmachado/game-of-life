package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/renatofmachado/game-of-life/pkg"
)

const (
  WINDOW_WIDTH = 640
  WINDOW_HEIGHT = 480
)

func main() {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Game of Life")

  game := pkg.NewGame(WINDOW_WIDTH, WINDOW_HEIGHT)
  game.SeedRandomLife()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
