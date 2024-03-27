package pkg

import (
  "image/color"
  "log"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
  "github.com/hajimehoshi/ebiten/v2/text"
  "golang.org/x/image/font"
  "golang.org/x/image/font/opentype"
)

const (
  DPI = 72
)

var mplusNormalFont font.Face

func DrawHud(g *Game, screen *ebiten.Image) {
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
