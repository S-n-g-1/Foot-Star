package schedule

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

func DrawJadwal(screen *ebiten.Image) {
	text.Draw(screen, "Tab: Jadwal", basicfont.Face7x13, 100, 200, color.White)
}
