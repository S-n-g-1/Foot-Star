package transfer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

func DrawTransfer(screen *ebiten.Image) {
	text.Draw(screen, "Tab: Transfer", basicfont.Face7x13, 100, 240, color.White)
}
