package logic

// This source file holds all the instructions related to the playing area, including calling all methods related to the
// fallen blocks, player piece and blocks generation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
	"time"
)

func RandomPieceIndex() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7) + 1
}

func RandomPieceColorIndex() [3]int {
	rand.Seed(time.Now().UnixNano())
	colorValues := [3]int{rand.Intn(255) + 1, rand.Intn(255) + 1, rand.Intn(255) + 1}
	return colorValues
}

type PlayingArea struct {
	x0, y0, x1, y1, bx, by float32
	board                  [20][10]bool
	blockPieces            *BlockPieces
	playerPiece            *PlayerPiece
	fallenBlocks           *FallenBlocks
}

func NewPlayingArea(ScreenWidth int, ScreenHeight int) *PlayingArea {
	offSet := float32(50)
	paddingX := float32(ScreenWidth / 5)
	paddingY := float32(ScreenWidth / 10)
	var grid [rows][cols]bool
	bp := NewBlockPieces()
	pp := NewPlayerPiece(bp.GenerateNewPiece(RandomPieceIndex()), RandomPieceColorIndex())
	fb := NewFallenBlocks()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			grid[i][j] = true
		}
	}
	return &PlayingArea{
		x0:           paddingX - offSet*1.2,
		y0:           paddingY + offSet*4,
		x1:           float32(ScreenWidth) - paddingX - offSet*1.5,
		y1:           float32(ScreenHeight) - paddingY + offSet*0.4,
		bx:           ((float32(ScreenWidth) - paddingX - offSet*1.5) - (paddingX - offSet*1.2)) / cols,
		by:           ((float32(ScreenHeight) - paddingY + offSet*0.4) - (paddingY + offSet*4)) / rows,
		board:        grid,
		blockPieces:  bp,
		playerPiece:  pp,
		fallenBlocks: fb,
	}
}

func (p *PlayingArea) UpdateBoard() {
	for _, pos := range p.playerPiece.position {
		p.board[pos[1]][pos[0]] = false
	}
}

func (p *PlayingArea) ResetPlayerPiece() {
	p.UpdateBoard()
	p.fallenBlocks.UpdateBlocks(p.playerPiece.position, [4]float32{p.x0, p.y0, p.bx, p.by}, p.playerPiece.color)
	pieceIndex := RandomPieceIndex()
	colorIndex := RandomPieceColorIndex()
	p.playerPiece = NewPlayerPiece(
		p.blockPieces.GenerateNewPiece(pieceIndex),
		colorIndex,
	)
}

func (p *PlayingArea) DrawBorders(screen *ebiten.Image) {
	strokeWidth := float32(1.0)
	borderColor := color.RGBA{0, 0, 0, 0xFF}
	// Top
	vector.StrokeLine(screen, p.x0, p.y0, p.x1, p.y0, strokeWidth, borderColor, true)
	// Left
	vector.StrokeLine(screen, p.x0, p.y0, p.x0, p.y1, strokeWidth, borderColor, true)
	// Bottom
	vector.StrokeLine(screen, p.x0, p.y1, p.x1, p.y1, strokeWidth, borderColor, true)
	// Right
	vector.StrokeLine(screen, p.x1, p.y1, p.x1, p.y0, strokeWidth, borderColor, true)
}

func (p *PlayingArea) DrawBackground(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.x0, p.y0, p.x1-p.x0, p.y1-p.y0,
		color.RGBA{240, 240, 245, 0xFF}, false)
}

func (p *PlayingArea) Draw(screen *ebiten.Image) {
	p.DrawBackground(screen)
	p.DrawBorders(screen)
	p.playerPiece.Draw(screen, p)
	p.fallenBlocks.Draw(screen)
}
