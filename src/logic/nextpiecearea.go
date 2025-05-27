package logic

// This file contains the instructions related to the displaying of the upcoming Tetromino

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type NextPieceArea struct {
	x0, y0, x1, y1, bx, by float32
	tetronimo              [4][2]int
	blockPieces            *BlockPieces
}

func NewNextPieceArea(NextPieceIndex int, ScreenWidth int, ScreenHeight int) *NextPieceArea {
	offSet := float32(50)
	paddingX := float32(ScreenWidth / 5)
	paddingY := float32(ScreenWidth / 10)
	bp := NewBlockPieces()
	te := bp.GenerateNewPiece(NextPieceIndex)
	return &NextPieceArea{
		x0:          paddingX + offSet*2.5,
		y0:          paddingY + offSet*4.5,
		x1:          float32(ScreenWidth) - paddingX - offSet*1.5,
		y1:          float32(ScreenHeight) - paddingY + offSet*0.4,
		bx:          ((float32(ScreenWidth) - paddingX - offSet*1.5) - (paddingX - offSet*1.2)) / cols,
		by:          ((float32(ScreenHeight) - paddingY + offSet*0.4) - (paddingY + offSet*4)) / rows,
		tetronimo:   te,
		blockPieces: bp,
	}
}

func (p *NextPieceArea) Update(newPieceIndex int) {
	p.tetronimo = p.blockPieces.GenerateNewPiece(newPieceIndex)
}

func (p *NextPieceArea) DrawBackground(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.x0+50, p.y0, p.x1*0.75, p.y1*0.15,
		color.RGBA{210, 230, 245, 0xFF}, false)
}

func (p *NextPieceArea) Draw(screen *ebiten.Image) {
	p.DrawBackground(screen)
	for _, pos := range p.tetronimo {
		vector.DrawFilledRect(screen,
			float32(pos[0])*p.bx+p.x0,
			float32(pos[1])*p.by+p.y0,
			p.bx,
			p.by,
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
			true)
	}

}
