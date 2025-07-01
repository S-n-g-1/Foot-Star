package views

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"

	"Foot-Star/views/scene/akademi"
	"Foot-Star/views/scene/stadion"
	"Foot-Star/views/scene/transfer"
)

//var selectedSubTab int = 0

var subTabs = []string{
	"Transfer",
	"Akademi",
	"Stadion",
}

var subTabRects []image.Rectangle

func DrawPengelolaan(screen *ebiten.Image) {
	drawSubTabs(screen)
	handleSubTabSwitch()

	switch selectedSubTab {
	case 0:
		transfer.DrawTransfer(screen)
	case 1:
		akademi.DrawAkademi(screen)
	case 2:
		stadion.DrawStadion(screen)
	}
}

func drawSubTabs(screen *ebiten.Image) {
	x := 50
	y := 60
	paddingX := 12
	paddingY := 6
	spacing := 10
	font := basicfont.Face7x13

	subTabRects = []image.Rectangle{}

	for i, name := range subTabs {
		textWidth := len(name) * 7
		tabWidth := textWidth + paddingX*2
		tabHeight := 13 + paddingY*2

		rect := image.Rect(x, y, x+tabWidth, y+tabHeight)
		subTabRects = append(subTabRects, rect)

		bgColor := color.RGBA{100, 100, 100, 255}
		textColor := color.White
		if i == selectedSubTab {
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

func handleSubTabSwitch() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for i, rect := range subTabRects {
			if x >= rect.Min.X && x <= rect.Max.X && y >= rect.Min.Y && y <= rect.Max.Y {
				selectedSubTab = i
				break
			}
		}
	}
}
