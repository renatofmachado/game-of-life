package pkg

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

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

type Game struct {
	cells       [][]bool
	cellSize    int
	screenW     int
	width       int
	screenH     int
	height      int
	interactive bool
	running     bool
}

func getCellsLimits(width, height, cellSize int) (int, int) {
	return width / cellSize, height / cellSize
}

func NewGame(width, height int) *Game {
	initialCellSize := 5
	w, h := getCellsLimits(width, height, initialCellSize)

	cells := make([][]bool, h)
	for i := range cells {
		cells[i] = make([]bool, w)
	}

	return &Game{
		cells:       cells,
		cellSize:    initialCellSize,
		screenW:     width,
		width:       w,
		screenH:     height,
		height:      h,
		interactive: false,
		running:     true,
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
	for i := 0; i < g.height/2; i++ {
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
	w, h := getCellsLimits(g.screenW, g.screenH, g.cellSize)

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
	RegisterIO(g)

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
				posX := float32(x * g.cellSize)
				posY := float32(y * g.cellSize)
				cellSize := float32(g.cellSize)
				vector.DrawFilledRect(screen, posX, posY, cellSize, cellSize, color.White, true)
			}
		}
	}

	DrawHud(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
