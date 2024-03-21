package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type menu struct{}

func (m *menu) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

}
