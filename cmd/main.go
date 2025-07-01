package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"

	"Foot-Star/views"
)

var (
	bigFont  font.Face
	menuFont font.Face
)

type Game struct {
	screen string
}

func init() {
	fontPath := "assets/font/PressStart2P-Regular.ttf"
	fontData, err := ioutil.ReadFile(fontPath)
	if err != nil {
		log.Fatal("Font not found in assets/font/: ", err)
	}
	tt, err := truetype.Parse(fontData)
	if err != nil {
		log.Fatal("Failed to parse font:", err)
	}
	bigFont = truetype.NewFace(tt, &truetype.Options{Size: 28})
	menuFont = truetype.NewFace(tt, &truetype.Options{Size: 14})
}

func (g *Game) Update() error {
	switch g.screen {
	case "menu":
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			g.screen = "main"
		}
	case "main":
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.screen = "menu"
		}
	}
	return nil
}

func drawRect(screen *ebiten.Image, x, y, w, h int, clr color.Color) {
	rect := ebiten.NewImage(w, h)
	rect.Fill(clr)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(rect, op)
}

func centerTextX(face font.Face, textStr string, containerX, containerW int) int {
	bounds := text.BoundString(face, textStr)
	textWidth := bounds.Dx() // Lebar dalam satuan pixel
	return containerX + (containerW-textWidth)/2
}

func drawTitleWithShadow(screen *ebiten.Image, title string, x, y int, face font.Face) {
	text.Draw(screen, title, face, x+3, y+3, color.Black)
	text.Draw(screen, title, face, x, y, color.RGBA{255, 215, 0, 255}) // Gold
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{25, 25, 25, 255}) // Latar belakang gelap

	// Ukuran dan posisi container (centered)
	containerW, containerH := 640, 440
	containerX := (800 - containerW) / 2
	containerY := (600 - containerH) / 2

	// Gambar kotak container

	// Garis tengah bantu debugging
	//drawRect(screen, containerX+containerW/2-1, containerY, 2, containerH, color.RGBA{255, 0, 0, 255})

	const padding = 20

	switch g.screen {
	case "menu":
		drawRect(screen, containerX, containerY, containerW, containerH, color.RGBA{45, 45, 45, 255})
		// Judul game
		title := "FOOT STAR"
		titleY := containerY + padding + 40
		titleX := centerTextX(bigFont, title, containerX, containerW)
		drawTitleWithShadow(screen, title, titleX, titleY, bigFont)

		// Opsi menu
		startText := "1: MULAI BARU"
		startY := titleY + 100
		startX := centerTextX(menuFont, startText, containerX, containerW)
		text.Draw(screen, startText, menuFont, startX, startY, color.White)

		// Keterangan tombol
		creditText := "Tekan [1] untuk mulai"
		creditY := startY + 50
		creditX := centerTextX(menuFont, creditText, containerX, containerW)
		text.Draw(screen, creditText, menuFont, creditX, creditY, color.RGBA{180, 180, 180, 255})

		// Debug posisi (opsional)
		posStr := fmt.Sprintf("TitleX: %d | StartX: %d | CreditX: %d", titleX, startX, creditX)
		text.Draw(screen, posStr, menuFont, 10, 590, color.RGBA{150, 150, 150, 255})

	case "main":
		views.DrawMainMenu(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowTitle("Foot-Star")
	ebiten.SetWindowSize(800, 600)

	game := &Game{screen: "menu"}
	views.OnBackToMainMenu = func() {
		game.screen = "menu"
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
