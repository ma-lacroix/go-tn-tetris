package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ma-lacroix/go-tn-tetris/tntetris"
	"log"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
const (
	screenWidth  = 240
	screenHeight = 240
	numSquares   = 100
)

func main() {
	g := &Game{
		squares: initializeSquares(numSquares), speed: 1.5, current: 0, outOfBounds: 0,
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Super Soup")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
