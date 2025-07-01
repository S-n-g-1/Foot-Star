package stadion

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

func DrawStadionUpgrade(screen *ebiten.Image) {
	text.Draw(screen, "Submenu: Upgrade Stadion", basicfont.Face7x13, 100, 300, color.White)
}
