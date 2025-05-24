package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type PlayerPiece struct {
	position *[4][2]int
	color    color.Color
}

func NewPlayerPiece(tetronimo *[4][2]int) *PlayerPiece {
	return &PlayerPiece{
		position: tetronimo,
		color:    color.RGBA{255, 0, 0, 255},
	}
}

func (pp *PlayerPiece) Rotation() {
	origin := pp.position[0]

	for i, block := range pp.position {
		dx := block[0] - origin[0]
		dy := block[1] - origin[1]
		pp.position[i][0] = origin[0] + dy
		pp.position[i][1] = origin[1] - dx
	}
}

func (pp *PlayerPiece) UpdatePlayerPiece(newMove [2]int) {
	for i := range pp.position {
		pp.position[i][0] += newMove[0]
		pp.position[i][1] += newMove[1]
	}
}

func (pp *PlayerPiece) Draw(screen *ebiten.Image, p *PlayingArea) {
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
