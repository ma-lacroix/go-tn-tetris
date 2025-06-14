package logic

// This source file holds all the instructions related to the playing area, including calling all methods related to the
// fallen blocks, player piece and blocks generation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type PlayingArea struct {
	x0, y0, x1, y1, bx, by float32
	board                  [rows][cols]bool
	blockPieces            *BlockPieces
	playerPiece            *PlayerPiece
	fallenBlocks           *FallenBlocks
	backgroundImage        *ebiten.Image
	isActive               bool
}

func NewPlayingArea(ScreenWidth int, ScreenHeight int) *PlayingArea {
	newPieceIndex := RandomPieceIndex()
	offSet := float32(50)
	paddingX := float32(ScreenWidth / 5)
	paddingY := float32(ScreenWidth / 10)
	var grid [rows][cols]bool
	bp := NewBlockPieces()
	pp := NewPlayerPiece(bp.GenerateNewPiece(newPieceIndex), bp.GenerateNewPieceImageLocations(newPieceIndex))
	fb := NewFallenBlocks()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			grid[i][j] = true
		}
	}
	backgroundImage := loadImage("media/images/b_playing_area.png")
	return &PlayingArea{
		x0:              paddingX - offSet*1.55,
		y0:              paddingY + offSet*2.2,
		x1:              float32(ScreenWidth) - paddingX - offSet*1.65,
		y1:              float32(ScreenHeight) - paddingY + offSet*0.65,
		bx:              (float32(ScreenWidth) - 2*paddingX - offSet*0.10) / cols,
		by:              (float32(ScreenHeight) - 2*paddingY - offSet*1.55) / rows,
		board:           grid,
		blockPieces:     bp,
		playerPiece:     pp,
		fallenBlocks:    fb,
		backgroundImage: backgroundImage,
		isActive:        true,
	}
}

func (p *PlayingArea) CanPlaceNewPiece() bool {
	return p.board[1][4] == true || p.board[1][5] == true
}

func (p *PlayingArea) StopGame() {
	p.isActive = false
}

func (p *PlayingArea) clearFullRowsAndShiftDown() {
	for row := rows - 1; row >= 0; row-- {
		if checkCompletedRows(p.board[row][:]) {
			p.clearRow(row)
			p.shiftRowsDown(row)
			row++
		}
	}
}

func checkCompletedRows(row []bool) bool {
	for _, v := range row {
		if v {
			return false
		}
	}
	return true
}

func (p *PlayingArea) clearRow(row int) {
	for col := 0; col < cols; col++ {
		p.board[row][col] = true
	}
}

func (p *PlayingArea) shiftRowsDown(startRow int) {
	for row := startRow; row > 0; row-- {
		p.board[row] = p.board[row-1]
	}
	for col := 0; col < cols; col++ {
		p.board[0][col] = true
	}
}

func (p *PlayingArea) UpdateBoard() {
	for _, pos := range p.playerPiece.position {
		p.board[pos[1]][pos[0]] = false
	}
	p.clearFullRowsAndShiftDown()
}

func (p *PlayingArea) ResetPlayerPiece(pieceIndex int) {
	p.UpdateBoard()
	p.fallenBlocks.UpdateBlocks(p.playerPiece.position,
		p.playerPiece.imagePositions,
		[4]float32{p.x0, p.y0, p.bx, p.by},
		p.playerPiece.color,
		angles[p.playerPiece.rotationIndex])
	p.playerPiece = NewPlayerPiece(
		p.blockPieces.GenerateNewPiece(pieceIndex),
		p.blockPieces.GenerateNewPieceImageLocations(pieceIndex),
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

func (p *PlayingArea) Draw(screen *ebiten.Image, superDrop bool, skipDrawPlayer bool, score int32) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(imageScaleX, imageScaleY)
	op.GeoM.Rotate(float64(score) / 200)
	op.GeoM.Translate(float64(p.x0), float64(p.y0)-float64(score))
	screen.DrawImage(p.backgroundImage, op)
	if !skipDrawPlayer {
		if superDrop {
			p.playerPiece.DrawSuperDrop(screen, p)
		}
		if p.isActive {
			p.playerPiece.Draw(screen, p)
		}
	}
	p.fallenBlocks.Draw(screen)
	p.fallenBlocks.DrawExplodingBlocks(screen)
}
