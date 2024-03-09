package assets

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Assets struct {
}

//go:embed *
var assets embed.FS

var Tilemap = getSingleImage("Sprites/tilemap.png")

//go:embed Fonts/AvenirNext-DemiBoldItalic.ttf
var AvenirNext_ttf []byte

func getSingleImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
