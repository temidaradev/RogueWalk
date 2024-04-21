package game

import (
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type char struct {
	x  int
	y  int
	vx int
	vy int
}

const (
	groundY = 395
	unit    = 10
)

func (c *char) update() {
	c.x += c.vx
	c.y += c.vy

	if c.vx > 0 {
		c.vx -= 5
	} else if c.vx < 0 {
		c.vx += 5
	}
	if c.vy > 0 {
		c.vy -= 5
	} else if c.vy < 0 {
		c.vy += 5
	}
}

type Player struct {
	player *char
}

func (p *Player) Update() error {
	if p.player == nil {
		p.player = &char{x: 50 * unit, y: groundY * unit}
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.player.vy = -4 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.player.vy = 4 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.player.vx = -4 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.player.vx = 4 * unit
	}

	if p.player.x <= -6400 {
		p.player.x = -6400
	}
	if p.player.x >= 18000 {
		p.player.x = 18000
	}

	if p.player.y <= -4850 {
		p.player.y = -4850
	}
	if p.player.y >= 19600 {
		p.player.y = 19600
	}

	p.player.update()
	return nil
}

func (p *Player) Draw(screen *ebiten.Image, cam *camera, g *Game) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.15, 0.15)
	op.GeoM.Translate(float64(p.player.x)/unit, float64(p.player.y)/unit)

	cam.draw(assets.Chars, op)
}
