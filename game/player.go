package game

import (
	"image"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	g Game
}

func (p *Player) Update() error {
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	w := assets.Chars.Bounds().Dx()
	tileXCount := w / tileSize

	const xCount = screenWidth / tileSize
	for _, l := range p.g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize
			screen.DrawImage(assets.Tilemap.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}
}
