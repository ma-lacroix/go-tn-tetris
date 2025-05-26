package logic

// This source file handles the main game loop and user input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
)

const (
	rows = 20
	cols = 10
)

type Game struct {
	ScreenWidth     int
	ScreenHeight    int
	NextPieceArea   *NextPieceArea
	PlayingArea     *PlayingArea
	moveCooldown    int
	moveCooldownMax int
	lastDropTime    time.Time
	dropInterval    time.Duration
}

func NewGame(width, height int) *Game {
	return &Game{
		ScreenWidth:     width,
		ScreenHeight:    height,
		NextPieceArea:   NewNextPieceArea(),
		PlayingArea:     NewPlayingArea(width, height),
		moveCooldownMax: 10,
		dropInterval:    1000 * time.Millisecond,
	}
}

func (g *Game) Reset() {
	g.NextPieceArea = NewNextPieceArea()
	g.PlayingArea = NewPlayingArea(g.ScreenWidth, g.ScreenHeight)
}

func (g *Game) Update() error {
	down := false
	if g.moveCooldown > 0 {
		g.moveCooldown--
		return nil
	}
	if time.Since(g.lastDropTime) > g.dropInterval {
		down := [2]int{0, 1}
		if g.PlayingArea.playerPiece.CollisionDetection(down, cols, rows, &g.PlayingArea.board) {
			g.PlayingArea.playerPiece.UpdatePlayerPiece(down)
		}
		g.lastDropTime = time.Now()
	}
	move := [2]int{0, 0}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.Reset()
		g.moveCooldown = g.moveCooldownMax
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
		down = true
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.PlayingArea.playerPiece.Rotation(cols, rows)
		g.moveCooldown = g.moveCooldownMax
	}
	if move != [2]int{0, 0} {
		if g.PlayingArea.playerPiece.CollisionDetection(move, cols, rows, &g.PlayingArea.board) {
			g.PlayingArea.playerPiece.UpdatePlayerPiece(move)
		}
		if down {
			g.moveCooldown = g.moveCooldownMax / 10
			down = false
		} else {
			g.moveCooldown = g.moveCooldownMax
		}
	}
	if g.PlayingArea.playerPiece.ShouldLock(rows, 500*time.Millisecond, &g.PlayingArea.board) {
		g.PlayingArea.ResetPlayerPiece()
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
