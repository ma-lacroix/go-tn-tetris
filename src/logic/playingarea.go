package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const (
	rows = 20
	cols = 10
)

type PlayingArea struct {
	x0, y0, x1, y1, bx, by float32
	board                  [20][10]bool
	playerPiece            *PlayerPiece
}

type PlayerPiece struct {
	position [4][2]int
	color    color.Color
}

func NewPlayingArea(ScreenWidth int, ScreenHeight int) *PlayingArea {
	offSet := float32(50)
	paddingX := float32(ScreenWidth / 5)
	paddingY := float32(ScreenWidth / 10)
	var grid [rows][cols]bool
	for i := 0; i < 20; i++ {
		for j := 0; j < 10; j++ {
			grid[i][j] = true
		}
	}
	return &PlayingArea{
		x0:    paddingX - offSet*1.2,
		y0:    paddingY + offSet*4,
		x1:    float32(ScreenWidth) - paddingX - offSet*1.5,
		y1:    float32(ScreenHeight) - paddingY + offSet*0.4,
		bx:    ((float32(ScreenWidth) - paddingX - offSet*1.5) - (paddingX - offSet*1.2)) / cols,
		by:    ((float32(ScreenHeight) - paddingY + offSet*0.4) - (paddingY + offSet*4)) / rows,
		board: grid,
		playerPiece: &PlayerPiece{
			position: [4][2]int{
				{4, 1}, {5, 1}, {5, 2}, {5, 3},
			},
			color: color.RGBA{255, 0, 0, 255},
		},
	}
}

func (pp *PlayerPiece) DrawPlayerPiece(screen *ebiten.Image, p *PlayingArea) {

	for _, pos := range pp.position {
		vector.DrawFilledRect(screen,
			float32(pos[0])*p.bx+p.x0,
			float32(pos[1])*p.by+p.y0,
			p.bx,
			p.by,
			pp.color,
			true)
	}
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

func (p *PlayingArea) Init() {

}

func (p *PlayingArea) UpdatePlayerPiece(newMove [2]int) {
	for i := range p.playerPiece.position {
		p.playerPiece.position[i][0] += newMove[0]
		p.playerPiece.position[i][1] += newMove[1]
	}
}

func (p *PlayingArea) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.x0, p.y0, p.x1-p.x0, p.y1-p.y0,
		color.RGBA{200, 200, 150, 0xFF}, false)
	p.DrawBorders(screen)
	p.playerPiece.DrawPlayerPiece(screen, p)
}
