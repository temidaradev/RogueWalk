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
	c      *camera
}

func NewGame() *Game {
	g := &Game{
		assets.L.Layers,
		&menu{},
		&Player{},
		&camera{},
	}
	g.c.init()
	return g
}

var isStarted bool
var zoomIn, zoomOut bool

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
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		zoomIn = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		zoomIn = false
	}
	g.p.Update()
	g.c.setPos(g.p.player.x/unit-320, g.p.player.y/unit-240)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !isStarted {
		g.m.Draw(screen)
	} else {
		screen.Fill(color.RGBA{133, 198, 106, 255})
		w := assets.Tilemap.Bounds().Dx()
		tileXCount := w / tileSize

		const xCount = screenWidth / tileSize
		for _, l := range g.layers {
			for i, t := range l {
				op := &ebiten.DrawImageOptions{}
				if zoomIn {
					op.GeoM.Translate(screenWidth/2, screenHeight/2)
					op.GeoM.Scale(2, 2)
					op.GeoM.Translate(float64((i%xCount)*tileSize*2), float64((i/xCount)*tileSize*2))
				} else {
					op.GeoM.Translate(screenWidth/2, screenHeight/2)
					op.GeoM.Scale(1, 1)
					op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))
				}

				sx := (t % tileXCount) * tileSize
				sy := (t / tileXCount) * tileSize
				g.c.draw(assets.Tilemap.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
			}
		}
		g.p.Draw(screen, g.c, g)
	}
	g.c.render(screen)
	g.c.clear()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
