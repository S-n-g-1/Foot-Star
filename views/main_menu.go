package views

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"

	//"Foot-Star/views/scene/formasi"
	"Foot-Star/views/scene/schedule"
)

var selectedTab int = 1 // default ke "Lineup"
var selectedSubTab int = 0

var tabs = []string{
	"Lineup",
	"Pemain",
	"Pengelolaan",
	"Jadwal",
}

var tabRects []image.Rectangle
var backButtonRect image.Rectangle

// Fungsi callback saat tombol kembali diklik
var OnBackToMainMenu func()

func DrawMainMenu(screen *ebiten.Image) {
	screenW, _ := screen.Size()

	tabWidth := 120
	tabHeight := 30
	paddingTop := 20
	paddingLeft := 30
	tabSpacing := 10

	tabRects = []image.Rectangle{}

	// Gambar semua tab
	for i, name := range tabs {
		x := paddingLeft + i*(tabWidth+tabSpacing)
		rect := image.Rect(x, paddingTop, x+tabWidth, paddingTop+tabHeight)
		tabRects = append(tabRects, rect)

		bgColor := color.RGBA{80, 80, 80, 255}
		textColor := color.White

		if i == selectedTab {
			bgColor = color.RGBA{255, 215, 0, 255}
			textColor = color.Black
		}

		drawRect(screen, rect, bgColor)

		textX := x + (tabWidth-len(name)*7)/2
		textY := paddingTop + 20
		text.Draw(screen, name, basicfont.Face7x13, textX, textY, textColor)
	}

	// Tombol "Kembali" di pojok kanan atas
	backButtonWidth := 100
	backButtonHeight := 30
	backButtonX := screenW - backButtonWidth - paddingLeft
	backButtonY := paddingTop
	backButtonRect = image.Rect(backButtonX, backButtonY, backButtonX+backButtonWidth, backButtonY+backButtonHeight)

	drawBackButton(screen)
	handleMouseTabSwitch()
	handleBackButton()

	switch selectedTab {
	case 0:
		DrawLineup(screen)
	case 1:
		DrawPlayerData(screen)
	case 2:
		DrawPengelolaan(screen)
	case 3:
		schedule.DrawJadwal(screen)
	}
}

func drawBackButton(screen *ebiten.Image) {
	bgColor := color.RGBA{180, 0, 0, 255}
	textColor := color.White

	if mouseOver(backButtonRect) {
		bgColor = color.RGBA{220, 20, 20, 255}
	}

	drawRect(screen, backButtonRect, bgColor)

	textX := backButtonRect.Min.X + (100 - len("Kembali")*7)/2
	textY := backButtonRect.Min.Y + 20
	text.Draw(screen, "Kembali", basicfont.Face7x13, textX, textY, textColor)
}

func handleMouseTabSwitch() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for i, rect := range tabRects {
			if x >= rect.Min.X && x <= rect.Max.X && y >= rect.Min.Y && y <= rect.Max.Y {
				selectedTab = i
				selectedSubTab = 0
				break
			}
		}
	}
}

func handleBackButton() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if mouseOver(backButtonRect) && OnBackToMainMenu != nil {
			OnBackToMainMenu()
		}
	}
}

func mouseOver(rect image.Rectangle) bool {
	x, y := ebiten.CursorPosition()
	return x >= rect.Min.X && x <= rect.Max.X && y >= rect.Min.Y && y <= rect.Max.Y
}

func drawRect(screen *ebiten.Image, rect image.Rectangle, clr color.Color) {
	img := ebiten.NewImage(rect.Dx(), rect.Dy())
	img.Fill(clr)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(rect.Min.X), float64(rect.Min.Y))
	screen.DrawImage(img, op)
}

// Dummy tab content

func DrawPlayerData(screen *ebiten.Image) {
	text.Draw(screen, "Tab: Pemain", basicfont.Face7x13, 100, 140, color.White)
}
