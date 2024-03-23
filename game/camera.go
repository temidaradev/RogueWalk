package game

import "github.com/hajimehoshi/ebiten/v2"

type camera struct {
	x int
	y int

	drawable *ebiten.Image // the image that the camera will draw
}

func (c *camera) setPos(x, y int) {
	c.x = x
	c.y = y
}

func (c *camera) init() {
	c.drawable = ebiten.NewImage(640, 480)
}

func (camera *camera) render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(camera.drawable, op)
}

func (c *camera) draw(image *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(float64(-c.x), float64(-c.y))

	c.drawable.DrawImage(image, op)
}

func (c *camera) clear() {
	c.drawable.Clear()
}
