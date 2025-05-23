package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ScreenWidth   int
	ScreenHeight  int
	FallenBlocks  *FallenBlocks
	NextPieceArea *NextPieceArea
	PlayingArea   *PlayingArea
}

func NewGame(width, height int) *Game {
	return &Game{
		ScreenWidth:   width,
		ScreenHeight:  height,
		FallenBlocks:  NewFallenBlocks(),
		NextPieceArea: NewNextPieceArea(),
		PlayingArea:   NewPlayingArea(width, height),
	}
}

func (g *Game) Reset() {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.PlayingArea.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
