package logic

// This source file handles the main game loop and user input

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"math/rand"
	"time"
)

const (
	rows        = 20
	cols        = 10
	imageScaleX = 0.5625
	imageScaleY = 0.5333
)

type Game struct {
	ScreenWidth       int
	ScreenHeight      int
	Menu              *Menu
	NextPieceArea     *NextPieceArea
	NextPieceIndex    int
	ScoreBoard        *ScoreBoard
	PlayingArea       *PlayingArea
	moveCooldown      int
	moveCooldownMax   int
	lastDropTime      time.Time
	dropInterval      time.Duration
	animationTime     time.Time
	animationInterval time.Duration
	backgroundImage   *ebiten.Image
}

func NewGame(width, height int) *Game {
	nextPieceIndex := RandomPieceIndex()
	bgImage := loadImage("../media/images/b_background.png")
	return &Game{
		ScreenWidth:       width,
		ScreenHeight:      height,
		Menu:              NewMenu(),
		NextPieceArea:     NewNextPieceArea(nextPieceIndex, width, height),
		NextPieceIndex:    nextPieceIndex,
		ScoreBoard:        NewScoreBoard(width, height),
		PlayingArea:       NewPlayingArea(width, height),
		moveCooldownMax:   10,
		dropInterval:      1000 * time.Millisecond,
		animationInterval: 10 * time.Millisecond,
		backgroundImage:   bgImage,
	}
}

func (g *Game) Reset(width int, height int) {
	nextPieceIndex := RandomPieceIndex()
	g.NextPieceArea = NewNextPieceArea(nextPieceIndex, width, height)
	g.PlayingArea = NewPlayingArea(g.ScreenWidth, g.ScreenHeight)
}

func RandomPieceIndex() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7) + 1
}

func (g *Game) HandleMenuInput() {
	if g.moveCooldown > 0 {
		g.moveCooldown--
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		selected--
		if selected < 0 {
			selected = len(menuOptions) - 1
		}
		g.moveCooldown = g.moveCooldownMax
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		selected++
		if selected >= len(menuOptions) {
			selected = 0
		}
		g.moveCooldown = g.moveCooldownMax
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		switch selected {
		case 0:
			fmt.Println("Starting game...")
			g.Menu.isActive = false
		case 1:
			fmt.Println("Options selected")
		case 2:
			fmt.Println("Quitting...")
		}
	}

}

func (g *Game) HandleMainGameInput() {
	down := false
	if g.moveCooldown > 0 {
		g.moveCooldown--
		return
	}
	if time.Since(g.lastDropTime) > g.dropInterval {
		down := [2]int{0, 1}
		if g.PlayingArea.playerPiece.CollisionDetection(down, &g.PlayingArea.board) {
			g.PlayingArea.playerPiece.UpdatePlayerPiece(down)
		}
		g.lastDropTime = time.Now()
	}
	move := [2]int{0, 0}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.Reset(g.ScreenWidth, g.ScreenHeight)
		g.moveCooldown = g.moveCooldownMax
		g.Menu.isActive = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		move[0] = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		move[0] = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		move[1] = 1
		down = true
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.PlayingArea.playerPiece.Rotation(&g.PlayingArea.board)
		g.moveCooldown = g.moveCooldownMax
		return
	}
	if move != [2]int{0, 0} {
		if g.PlayingArea.playerPiece.CollisionDetection(move, &g.PlayingArea.board) {
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
		g.PlayingArea.ResetPlayerPiece(g.NextPieceIndex)
		g.NextPieceIndex = RandomPieceIndex()
		g.NextPieceArea.Update(g.NextPieceIndex)
		g.ScoreBoard.Update(g.PlayingArea.fallenBlocks.rowsRemoved)

	}
}

func (g *Game) Update() error {
	if g.Menu.isActive {
		g.HandleMenuInput()
	} else {
		g.HandleMainGameInput()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Menu.isActive {
		g.Menu.Draw(screen)
	} else {
		if time.Since(g.animationTime) > g.animationInterval {
			g.PlayingArea.fallenBlocks.MoveExplodingBlocks()
			g.animationTime = time.Now()
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(imageScaleX, imageScaleY)
		screen.DrawImage(g.backgroundImage, op)
		g.PlayingArea.Draw(screen)
		g.NextPieceArea.Draw(screen)
		g.ScoreBoard.Draw(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
