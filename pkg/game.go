package pkg

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
  CELL_SIZE = 10
)

type Game struct{
  cells [][]bool
  width int
  height int
}

func NewGame(width, height int) *Game {
  cells := make([][]bool, height)
  for i := range cells {
    cells[i] = make([]bool, width)
  }

  return &Game{
    cells: cells,
    width: width,
    height: height,
  }
}

func (g *Game) SeedRandomLife() {
  for y := range g.cells {
    for x := range g.cells[y] {
      if rand.Intn(20) == 1 {
        g.cells[y][x] = true
      }
    }
  }
}

// 8 neighbors coordinates expressed as [x, y]
var directions = [][]int{
  {-1, -1},
  {0, -1},
  {1, -1},
  {-1, 0},
  {1, 0},
  {-1, 1},
  {0, 1},
  {1, 1},
}

func countNeighbors(cells [][]bool, x, y int) int {
  count := 0
  
  for _, coordinates := range directions {
    neighborX := x + coordinates[0]
    neighborY := y + coordinates[1]

    if neighborX >= 0 && neighborX < len(cells[0]) && neighborY >= 0 && neighborY < len(cells) {
      if cells[neighborY][neighborX] {
        count++
      }
    }
  }

  return count
}

func (g *Game) nextGeneration() {
  nextGen := make([][]bool, g.height)
  for i := range nextGen {
    nextGen[i] = make([]bool, g.width)
  }

  for y := range g.cells {
    for x := range g.cells[y] {
      // Rule 2
      nextGen[y][x] = g.cells[y][x]

      neighborsCount := countNeighbors(g.cells, x, y)

      // Rule 1 and 3
      if neighborsCount < 2 || neighborsCount > 3 {
        nextGen[y][x] = false
      }

      // Rule 2 and 4
      if neighborsCount == 3 {
        nextGen[y][x] = true
      }
    }
  }

  g.cells = nextGen
}

func (g *Game) Update() error {
	g.nextGeneration()
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  for y := range g.cells {
    for x := range g.cells[y] {
      if g.cells[y][x] {
        vector.DrawFilledRect(screen, float32(x * CELL_SIZE), float32(y * CELL_SIZE), CELL_SIZE, CELL_SIZE, color.White, true)
      }
    }
  }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
