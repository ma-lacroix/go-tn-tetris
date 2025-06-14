package logic

// This source file handles the main game loop and user input

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	_ "image/png"
	"time"
)

const (
	rows             = 20
	cols             = 10
	imageScaleX      = 0.5625
	imageScaleY      = 0.5333
	blockImageScaleX = 0.21
	blockImageScaleY = 0.21
)

var (
	ctx = audio.NewContext(44100)
)

//go:embed media/images/*
var imageFS embed.FS

//go:embed media/font/*
var fontFS embed.FS

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
	dropFactor        int
	dropInterval      time.Duration
	animationTime     time.Time
	animationInterval time.Duration
	backgroundImage   *ebiten.Image
	Messages          *Messages
	superDrop         bool
	superDropNoise    bool
	gameOver          *GameOver
	soundBank         *SoundBank
	musicBank         *MusicBank
}

func NewGame(width, height int) *Game {
	nextPieceIndex := RandomPieceIndex()
	bgImage := loadImage("media/images/b_background.png")
	dropFactor := 1000
	return &Game{
		ScreenWidth:       width,
		ScreenHeight:      height,
		Menu:              NewMenu(),
		NextPieceArea:     NewNextPieceArea(nextPieceIndex, width, height),
		NextPieceIndex:    nextPieceIndex,
		ScoreBoard:        NewScoreBoard(width, height),
		PlayingArea:       NewPlayingArea(width, height),
		moveCooldownMax:   10,
		dropFactor:        1000,
		dropInterval:      time.Duration(dropFactor) * time.Millisecond,
		animationInterval: 10 * time.Millisecond,
		backgroundImage:   bgImage,
		Messages:          NewMessages(),
		superDrop:         false,
		superDropNoise:    false,
		gameOver:          NewGameOver(width, height),
		soundBank:         NewSoundBank(ctx),
		musicBank:         NewMusicBankBank(ctx),
	}
}

func (g *Game) increaseDifficulty(score int32) {
	if score > 0 && score%2 == 0 && g.dropFactor >= 100 {
		newDropFactor := g.dropFactor - 40
		g.dropInterval = time.Duration(newDropFactor) * time.Millisecond
		g.dropFactor = newDropFactor
	}
}

func (g *Game) Reset(width int, height int) {
	nextPieceIndex := RandomPieceIndex()
	g.musicBank.Stop("s_playing")
	g.dropInterval = 1000 * time.Millisecond
	g.NextPieceArea = NewNextPieceArea(nextPieceIndex, width, height)
	g.PlayingArea = NewPlayingArea(g.ScreenWidth, g.ScreenHeight)
	g.gameOver.isActive = false
	g.Menu.isActive = true
	g.ScoreBoard = NewScoreBoard(width, height)
}

func (g *Game) HandleMenuInput() {
	if g.moveCooldown > 0 {
		g.moveCooldown--
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.soundBank.Play("n_switch", 1)
		selected--
		if selected < 0 {
			selected = len(menuOptions) - 1
		}
		g.moveCooldown = g.moveCooldownMax
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.soundBank.Play("n_switch", 1)
		selected++
		if selected >= len(menuOptions) {
			selected = 0
		}
		g.moveCooldown = g.moveCooldownMax
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.soundBank.Play("n_explode1", 1)
		switch selected {
		case 0:
			fmt.Println("Easy")
		case 1:
			fmt.Println("Medium")
			g.dropInterval = 500 * time.Millisecond
		case 2:
			fmt.Println("Hard")
			g.dropInterval = 200 * time.Millisecond
		}
		g.Menu.isActive = false
		g.musicBank.Stop("s_menu")
	}

}

func (g *Game) HandleMainGameInput() {
	g.Messages.MoveActiveMessage()
	if g.Messages.active {
		return
	}
	if !g.PlayingArea.CanPlaceNewPiece() {
		g.PlayingArea.fallenBlocks.markAllAsAnimated()
		g.soundBank.Play("n_explode2", 1)
		g.PlayingArea.StopGame()
	}
	g.superDrop = false
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
		g.superDropNoise = false
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
		g.soundBank.Play("n_rotate", 1)
		g.PlayingArea.playerPiece.Rotation(&g.PlayingArea.board)
		g.moveCooldown = g.moveCooldownMax
		return
	}
	if move != [2]int{0, 0} {
		if g.PlayingArea.playerPiece.CollisionDetection(move, &g.PlayingArea.board) {
			g.PlayingArea.playerPiece.UpdatePlayerPiece(move)
		}
		if down {
			if !g.superDropNoise {
				g.soundBank.Play("n_afterburner", 0.2)
				g.superDropNoise = true
			}
			g.moveCooldown = g.moveCooldownMax / 10
			g.superDrop = true
			down = false
		} else {
			g.moveCooldown = g.moveCooldownMax
		}
	}
	if g.PlayingArea.playerPiece.ShouldLock(rows, 300*time.Millisecond, &g.PlayingArea.board) {
		g.soundBank.Play("n_impact2", 0.8)
		g.PlayingArea.ResetPlayerPiece(g.NextPieceIndex)
		g.NextPieceIndex = RandomPieceIndex()
		g.NextPieceArea.Update(g.NextPieceIndex)
		g.ScoreBoard.Update(g.PlayingArea.fallenBlocks.rowsRemoved)
		g.increaseDifficulty(g.ScoreBoard.score)
		if g.PlayingArea.fallenBlocks.rowsRemoved > 0 {
			g.soundBank.Play("n_explode2", 1)
			g.Messages.ActivateMessage(g.PlayingArea.fallenBlocks.rowsRemoved)
			switch g.PlayingArea.fallenBlocks.rowsRemoved {
			case 1:
				g.soundBank.Play("n_good", 0.8)
			case 2:
				g.soundBank.Play("n_all_right", 0.8)
			case 3:
				g.soundBank.Play("n_yyy", 0.8)
			case 4:
				g.soundBank.Play("n_onbc", 0.8)
			}
			g.PlayingArea.fallenBlocks.ResetRowsToRemove()
		}
	}
	move = [2]int{0, 0}
}

func (g *Game) HandleLastExplosion() {
	// TODO: turn this into a proper handler and slow down the animations
	g.PlayingArea.fallenBlocks.MoveExplodingBlocks()
	if len(g.PlayingArea.fallenBlocks.blocksToAnimate) == 0 {
		g.gameOver.isActive = true
		g.soundBank.Play("n_gameOver", 1)
	}
}

func (g *Game) HandleGameOverScreen() {
	g.gameOver.FlickerBackground()
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.Reset(g.ScreenWidth, g.ScreenHeight)
		g.moveCooldown = g.moveCooldownMax
		g.gameOver.isActive = false
		g.Menu.isActive = true
	}
}

func (g *Game) Update() error {
	if g.Menu.isActive {
		g.musicBank.Play("s_menu", 0.5)
		g.HandleMenuInput()
	} else if g.PlayingArea.isActive {
		g.musicBank.Play("s_playing", 0.5)
		g.HandleMainGameInput()
	} else if g.gameOver.isActive {
		g.HandleGameOverScreen()
	} else {
		g.HandleLastExplosion()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Menu.isActive {
		g.Menu.Draw(screen)
	} else if g.PlayingArea.isActive {
		if time.Since(g.animationTime) > g.animationInterval {
			g.PlayingArea.fallenBlocks.MoveExplodingBlocks()
			g.animationTime = time.Now()
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(imageScaleX, imageScaleY)
		screen.DrawImage(g.backgroundImage, op)
		g.PlayingArea.Draw(screen, g.superDrop, g.Messages.active, g.ScoreBoard.score)
		g.NextPieceArea.Draw(screen, g.ScoreBoard.score)
		g.ScoreBoard.Draw(screen, g.ScoreBoard.score)
		g.Messages.Draw(screen)
	} else if g.gameOver.isActive {
		g.gameOver.Draw(screen)
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(imageScaleX, imageScaleY)
		screen.DrawImage(g.backgroundImage, op)
		g.PlayingArea.Draw(screen, g.superDrop, false, g.ScoreBoard.score)
		g.NextPieceArea.Draw(screen, g.ScoreBoard.score)
		g.ScoreBoard.Draw(screen, g.ScoreBoard.score)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
