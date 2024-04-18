package game

import (
	"fmt"
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	layers [][]int
	m      *menu
	p      *Player
	c      *camera
	a      *assets.Assets
}

func NewGame() *Game {
	g := &Game{
		assets.L.Layers,
		&menu{},
		&Player{},
		&camera{},
		&assets.Assets{},
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
		screen.Fill(color.RGBA{132, 198, 105, 255})
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 1)
		op.GeoM.Translate(-screenWidth/2, -screenHeight/2)
		g.c.draw(assets.Tile, op)
		g.p.Draw(screen, g.c, g)
	}
	g.c.render(screen)
	g.c.clear()
	msg := fmt.Sprintf("Pos X: %d Pos Y: %d", g.p.player.x/unit, g.p.player.y/unit)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
