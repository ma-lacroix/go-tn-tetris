package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ma-lacroix/go-tn-tetris/src/logic"
	"log"
)

const (
	screenWidth  = 450
	screenHeight = 800
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("TnTetris on the go!")

	g := logic.NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
