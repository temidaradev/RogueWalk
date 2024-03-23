package game

import (
	"image"
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	layers [][]int
	m      *menu
	p      *Player
}

func NewGame() *Game {
	g := &Game{
		assets.L.Layers,
		&menu{},
		&Player{},
	}
	return g
}

var isStarted bool

const (
	screenWidth  = 640
	screenHeight = 480
)

const (
	tileSize = 16
)

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		isStarted = true
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !isStarted {
		g.p.Draw(screen)
		g.m.Draw(screen)
	} else {
		screen.Fill(color.RGBA{40, 40, 40, 255})
		w := assets.Tilemap.Bounds().Dx()
		tileXCount := w / tileSize

		const xCount = screenWidth / tileSize
		for _, l := range g.layers {
			for i, t := range l {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

				sx := (t % tileXCount) * tileSize
				sy := (t / tileXCount) * tileSize
				screen.DrawImage(assets.Tilemap.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
