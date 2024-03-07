package assets

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var TilemapBackground = getSingleImage("Sprites/tilemap-backgrounds.png")
var TilemapChars = getSingleImage("Sprites/tilemap-characters.png")
var Tilemap = getSingleImage("Sprites/tilemap.png")

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
