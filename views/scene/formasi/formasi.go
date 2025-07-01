package formasi

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

func DrawFormasi(screen *ebiten.Image) {
	text.Draw(screen, "INI HALAMAN FORMASI", basicfont.Face7x13, 100, 150, color.White)
}

