package game

import (
	"bytes"
	"image/color"
	"log"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type menu struct{}

var (
	fontFaceSource *text.GoTextFaceSource
)

const (
	normalFontSize = 24
	midFontSize    = 32
	bigFontSize    = 48
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(assets.AvenirNext_ttf))
	if err != nil {
		log.Fatal(err)
	}
	fontFaceSource = s
}

func (m *menu) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	m.title(screen)
}

func (m *menu) title(screen *ebiten.Image) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(160, 60)
	op.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})

	text.Draw(screen, "Mystery Game", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   bigFontSize,
	}, op)
}
