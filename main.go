package main

import (
  "flag"
  "log"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/renatofmachado/game-of-life/pkg"
)

const (
  WINDOW_WIDTH = 800
  WINDOW_HEIGHT = 600
)

var (
  random = flag.Int("random", -1, "How much of a chance there is of generating random life. 0-100")
)

func main() {
  flag.Parse()

  ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
  ebiten.SetWindowTitle("Game of Life")
  ebiten.SetTPS(20)

  game := pkg.NewGame(WINDOW_WIDTH, WINDOW_HEIGHT)

  if *random > 0 && *random <= 100 {
    game.SeedRandomLife(*random)
  }

  if err := ebiten.RunGame(game); err != nil {
    log.Fatal(err)
  }
}
