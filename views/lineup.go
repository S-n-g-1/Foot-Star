package views

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"

	"Foot-Star/views/scene/formasi"
)

var lineupSubTabs = []string{
	"Formasi",
	//"Strategi",
}

var lineupSubTabRects []image.Rectangle
var selectedLineupSubTab int = 0

func DrawLineup(screen *ebiten.Image) {
	drawLineupSubTabs(screen)
	handleLineupSubTabSwitch()

	switch selectedLineupSubTab {
	case 0:
		formasi.DrawFormasi(screen)
	case 1:
		drawStrategi(screen)
	}
}

func drawLineupSubTabs(screen *ebiten.Image) {
	x := 50
	y := 60
	paddingX := 12
	paddingY := 6
	spacing := 10
	font := basicfont.Face7x13

	lineupSubTabRects = []image.Rectangle{}

	for i, name := range lineupSubTabs {
		textWidth := len(name) * 7
		tabWidth := textWidth + paddingX*2
		tabHeight := 13 + paddingY*2

		rect := image.Rect(x, y, x+tabWidth, y+tabHeight)
		lineupSubTabRects = append(lineupSubTabRects, rect)

		bgColor := color.RGBA{100, 100, 100, 255}
		textColor := color.White
		if i == selectedLineupSubTab {
			bgColor = color.RGBA{200, 200, 200, 255}
			textColor = color.Black
		}

		drawRect(screen, rect, bgColor)
		textX := x + paddingX
		textY := y + paddingY + 13 - 2
		text.Draw(screen, name, font, textX, textY, textColor)

		x += tabWidth + spacing
	}
}

func handleLineupSubTabSwitch() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for i, rect := range lineupSubTabRects {
			if x >= rect.Min.X && x <= rect.Max.X && y >= rect.Min.Y && y <= rect.Max.Y {
				selectedLineupSubTab = i
				break
			}
		}
	}
}

// Dummy tampilan strategi
func drawStrategi(screen *ebiten.Image) {
	text.Draw(screen, "Strategi belum tersedia", basicfont.Face7x13, 100, 140, color.White)
}
