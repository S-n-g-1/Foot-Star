package dialog

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Font global untuk dialog
var Font font.Face

// Struktur data dialog
type Dialog struct {
	Character string
	Text      string
}

// Daftar dialog
var DialogList = []Dialog{
	{"Suji", "Hari pertama jadi pemain cadangan. Rasanya... aneh."},
	{"Pelatih", "Santai saja, Suji. Tunjukkan di latihan nanti."},
	{"Suji", "Baik, Pak. Saya akan berusaha."},
}

// State dialog
var (
	currentDialogIndex int
	IsFinished         bool
)

// Gaya tampilan
var (
	bgColor       = color.RGBA{10, 10, 10, 200}         // Background semi-transparan
	nameColor     = color.RGBA{255, 215, 0, 255}        // Nama karakter: emas
	textColor     = color.White                         // Teks dialog
	hintColor     = color.RGBA{180, 180, 180, 255}      // Teks petunjuk

	boxWidth      = 700
	boxHeight     = 160
	dialogBoxY    = 420
	padding       = 20

	shadowOffsetX = 2
	shadowOffsetY = 2
)

// DrawDialogScreen menggambar layar dialog
func DrawDialogScreen(screen *ebiten.Image) {
	if Font == nil {
		return
	}

	boxX := (800 - boxWidth) / 2
	boxY := dialogBoxY
	drawRect(screen, boxX, boxY, boxWidth, boxHeight, bgColor)

	if IsFinished {
		msg := "Dialog selesai. Tekan [ESC] untuk kembali."
		text.Draw(screen, msg, Font, boxX+padding, boxY+padding, textColor)
		return
	}

	d := DialogList[currentDialogIndex]

	// Gambar nama karakter
	nameY := boxY + padding + 10
	drawTextWithShadow(screen, d.Character+":", Font, boxX+padding, nameY, nameColor)

	// Gambar teks dialog dengan word wrap
	dialogY := nameY + 30
	lineSpacing := 25
	maxTextWidth := boxWidth - 2*padding
	lines := wrapText(d.Text, Font, maxTextWidth)

	for i, line := range lines {
		y := dialogY + i*lineSpacing
		if y > boxY+boxHeight-padding-lineSpacing {
			break // berhenti kalau lewat batas bawah kotak
		}
		drawTextWithShadow(screen, line, Font, boxX+padding, y, textColor)
	}

	// Gambar petunjuk
	hint := "[SPASI] untuk lanjut"
	hintX := boxX + boxWidth - textWidth(Font, hint) - padding
	hintY := boxY + boxHeight - padding
	drawTextWithShadow(screen, hint, Font, hintX, hintY, hintColor)
}

// NextDialog melanjutkan ke dialog berikutnya
func NextDialog() {
	if currentDialogIndex < len(DialogList)-1 {
		currentDialogIndex++
	} else {
		IsFinished = true
	}
}

// Reset mengulang dialog dari awal
func Reset() {
	currentDialogIndex = 0
	IsFinished = false
}

// drawRect menggambar kotak berwarna
func drawRect(screen *ebiten.Image, x, y, w, h int, clr color.Color) {
	rect := ebiten.NewImage(w, h)
	rect.Fill(clr)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(rect, op)
}

// drawTextWithShadow menggambar teks dengan bayangan
func drawTextWithShadow(screen *ebiten.Image, str string, face font.Face, x, y int, clr color.Color) {
	shadow := color.RGBA{0, 0, 0, 255}
	text.Draw(screen, str, face, x+shadowOffsetX, y+shadowOffsetY, shadow)
	text.Draw(screen, str, face, x, y, clr)
}

// textWidth menghitung lebar teks
func textWidth(face font.Face, str string) int {
	return text.BoundString(face, str).Dx()
}

// wrapText membungkus teks agar tidak melebihi lebar maksimum
func wrapText(str string, face font.Face, maxWidth int) []string {
	var lines []string
	words := strings.Fields(str)
	if len(words) == 0 {
		return lines
	}

	currentLine := words[0]
	for _, word := range words[1:] {
		testLine := currentLine + " " + word
		if textWidth(face, testLine) <= maxWidth {
			currentLine = testLine
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}
	lines = append(lines, currentLine)
	return lines
}
