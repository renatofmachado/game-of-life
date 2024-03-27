package pkg

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)


var CELL_SIZE int = 5

const (
  DPI = 72
)

var mplusNormalFont font.Face

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

type Game struct{
  cells [][]bool
  screenW int
  width int
  screenH int
  height int
  interactive bool
  running bool
  showHud bool
}

func getCellsLimits(width, height int) (int, int) {
  return width / CELL_SIZE, height / CELL_SIZE
}

func NewGame(width, height int) *Game {
  w, h := getCellsLimits(width, height)
  cells := make([][]bool, h)
  for i := range cells {
    cells[i] = make([]bool, w)
  }

  return &Game{
    cells: cells,
    screenW: width,
    width: w,
    screenH: height,
    height: h,
    interactive: false,
    running: true,
    showHud: true,
  }
}

func (g *Game) SeedLife(x, y int) {
  if x < 0 || x >= g.width || y < 0 || y >= g.height {
    return
  }

  g.cells[y][x] = true

  for _, coordinates := range directions {
    neighborX := x + coordinates[0]
    neighborY := y + coordinates[1]

    if neighborX < 0 || neighborX >= g.width || neighborY < 0 || neighborY >= g.height {
      continue
    }

    if rand.Intn(100) <= 50 {
      g.cells[neighborY][neighborX] = true
    }
  }
}

func (g *Game) SeedRandomLife(chance int) {
  for i := 0; i < g.height / 2; i++ {
    if rand.Intn(100) <= chance {
      randomX := rand.Intn(g.width)
      randomY := rand.Intn(g.height)
      g.SeedLife(randomX, randomY)
    }
  }
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

func (g *Game) setNewCellsSize() {
  w, h := getCellsLimits(g.screenW, g.screenH)

  newCells := make([][]bool, h)
  for i := range newCells {
    newCells[i] = make([]bool, w)
    if i < len(g.cells) {
      for j := range newCells[i] {
        if j < len(g.cells[i]) {
          newCells[i][j] = g.cells[i][j]
        }
      }
    }
  }

  g.cells = newCells
  g.width = w
  g.height = h
}

func (g *Game) Reset() {
  g.cells = make([][]bool, g.height)
  for i := range g.cells {
    g.cells[i] = make([]bool, g.width)
  }
}

func (g *Game) Update() error {
  if ebiten.IsKeyPressed(ebiten.KeyA) {
    g.SeedRandomLife(20)
  }

  if ebiten.IsKeyPressed(ebiten.KeyR) {
    g.Reset()
  }

  if inpututil.IsKeyJustReleased(ebiten.KeyI) {
    g.interactive = !g.interactive
  }

  if inpututil.IsKeyJustReleased(ebiten.KeyH) {
    g.showHud = !g.showHud
  }

  if inpututil.IsKeyJustReleased(ebiten.KeyP) {
    g.running = !g.running
  }

  if inpututil.IsKeyJustReleased(ebiten.KeyMinus) {
    if (CELL_SIZE - 1 > 3) {
      CELL_SIZE -= 1
      
      g.setNewCellsSize()
    }
  }
  
  if inpututil.IsKeyJustReleased(ebiten.KeyEqual) {
    if (CELL_SIZE + 1 < 15) {
      CELL_SIZE += 1

      g.setNewCellsSize()
    }
  }

  if g.interactive {
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
      x, y := ebiten.CursorPosition()
      g.SeedLife(x / CELL_SIZE, y / CELL_SIZE)
    }
  }

  if !g.running {
    return nil
  }

	g.nextGeneration()

  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  for y := range g.cells {
    for x := range g.cells[y] {
      if g.cells[y][x] {
        vector.DrawFilledRect(screen, float32(x * CELL_SIZE), float32(y * CELL_SIZE), float32(CELL_SIZE), float32(CELL_SIZE), color.White, true) }
      }
  }

  if g.showHud {
    var mode string

    if g.interactive { 
      mode = "Interactive"
    }

    if !g.running {
      mode = "Paused"
    }

    if !g.running && g.interactive {
      mode = "Paused - Interactive"
    }

    if mode != "" {
      text.Draw(screen, mode, mplusNormalFont, 7, 590, color.White)
    }
  }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func init() {
  tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     DPI,
		Hinting: font.HintingFull,
	})
}
