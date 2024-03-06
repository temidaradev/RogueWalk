package main

import (
	"log"
	"main/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
