package akademi

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

func DrawAkademi(screen *ebiten.Image) {
	text.Draw(screen, "Tab: Akademi", basicfont.Face7x13, 100, 180, color.White)
}
