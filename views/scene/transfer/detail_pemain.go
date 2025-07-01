package transfer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

func DrawDetailPemain(screen *ebiten.Image) {
	text.Draw(screen, "Submenu: Detail Pemain", basicfont.Face7x13, 100, 280, color.White)
}
