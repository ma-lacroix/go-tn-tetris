package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"malacroix/tntetris/logic"
)

const (
	screenWidth  = 800
	screenHeight = 1500
)

func main() {
	g := &logic.Game{screenWidth, screenHeight}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("TnTetris on the go!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
