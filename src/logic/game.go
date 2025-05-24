package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	rows = 20
	cols = 10
)

type Game struct {
	ScreenWidth     int
	ScreenHeight    int
	FallenBlocks    *FallenBlocks
	NextPieceArea   *NextPieceArea
	PlayingArea     *PlayingArea
	moveCooldown    int
	moveCooldownMax int
}

func NewGame(width, height int) *Game {
	return &Game{
		ScreenWidth:     width,
		ScreenHeight:    height,
		FallenBlocks:    NewFallenBlocks(),
		NextPieceArea:   NewNextPieceArea(),
		PlayingArea:     NewPlayingArea(width, height),
		moveCooldownMax: 10,
	}
}

func (g *Game) Reset() {
	g.FallenBlocks = NewFallenBlocks()
	g.NextPieceArea = NewNextPieceArea()
	g.PlayingArea = NewPlayingArea(g.ScreenWidth, g.ScreenHeight)
}

func (g *Game) Update() error {
	if g.moveCooldown > 0 {
		g.moveCooldown--
		return nil
	}
	move := [2]int{0, 0}
	rotate := false
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.Reset()
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		move[0] = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		move[0] = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		move[1] = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		move[1] = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		rotate = true
	}
	if rotate {
		g.PlayingArea.playerPiece.Rotation()
		g.moveCooldown = g.moveCooldownMax
	}
	if move != [2]int{0, 0} {
		if g.PlayingArea.playerPiece.CollisionDetection(move, cols, rows) {
			g.PlayingArea.playerPiece.UpdatePlayerPiece(move)
		}
		g.moveCooldown = g.moveCooldownMax
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{240, 255, 255, 255})
	g.PlayingArea.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
