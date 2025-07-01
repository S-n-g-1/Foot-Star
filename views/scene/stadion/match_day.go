package stadion

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

func DrawStadion(screen *ebiten.Image) {
	text.Draw(screen, "Tab: Stadion", basicfont.Face7x13, 100, 220, color.White)
}
