package assets

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

type Assets struct {
}

//go:embed *
var assets embed.FS

// var Tilemap = getSingleImage("Sprites/tilemap.png")
var Chars = getSingleImage("Sprites/chars.png")
var Tile = getTiled("./assets/Sprites/Map/tile.tmx")

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

func getTiled(name string) *ebiten.Image {
	gameMap, err := tiled.LoadFile(name)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}

	fmt.Println(gameMap)

	renderer, err := render.NewRenderer(gameMap)
	if err != nil {
		fmt.Printf("map unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	// Render just layer 0 to the Renderer.
	err = renderer.RenderLayer(0)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	err = renderer.RenderLayer(1)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	err = renderer.RenderLayer(2)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	err = renderer.RenderLayer(3)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	err = renderer.RenderLayer(4)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	img := renderer.Result

	renderer.Clear()

	return ebiten.NewImageFromImage(img)
}
