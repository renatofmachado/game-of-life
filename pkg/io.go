package pkg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func RegisterIO(g *Game) {
    if ebiten.IsKeyPressed(ebiten.KeyA) {
    g.SeedRandomLife(20)
  }

  if ebiten.IsKeyPressed(ebiten.KeyR) {
    g.Reset()
  }

  if inpututil.IsKeyJustReleased(ebiten.KeyI) {
    g.interactive = !g.interactive
  }

  if inpututil.IsKeyJustReleased(ebiten.KeyP) {
    g.running = !g.running
  }

  if inpututil.IsKeyJustReleased(ebiten.KeyMinus) {
    if (g.cellSize - 1 > 3) {
      g.cellSize -= 1
      
      g.setNewCellsSize()
    }
  }
  
  if inpututil.IsKeyJustReleased(ebiten.KeyEqual) {
    if (g.cellSize + 1 < 15) {
      g.cellSize += 1

      g.setNewCellsSize()
    }
  }

  if g.interactive {
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
      x, y := ebiten.CursorPosition()
      cellX, cellY := getCellsLimits(x, y, g.cellSize)

      g.SeedLife(cellX, cellY)
    }
  }
} 
