package stadion

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

func DrawStadionMenu(screen *ebiten.Image) {
	text.Draw(screen, "Submenu: Stadion Upgrade", basicfont.Face7x13, 100, 260, color.White)
}
