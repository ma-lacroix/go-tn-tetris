package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type PlayingArea struct {
	x0 float32
	y0 float32
	x1 float32
	y1 float32
}

func NewPlayingArea(ScreenWidth int, ScreenHeight int) *PlayingArea {
	offSet := float32(50)
	paddingX := float32(ScreenWidth / 5)
	paddingY := float32(ScreenWidth / 10)
	return &PlayingArea{
		x0: paddingX - offSet*1.2,
		y0: paddingY + offSet*4,
		x1: float32(ScreenWidth) - paddingX - offSet,
		y1: float32(ScreenHeight) - paddingY + offSet*0.4,
	}
}

func (p *PlayingArea) Init() {

}

func (p *PlayingArea) Update() {

}

func (p *PlayingArea) Draw(screen *ebiten.Image) {
	strokeWidth := float32(1.0)
	borderColor := color.RGBA{200, 200, 20, 0xFF}
	// Top
	vector.StrokeLine(screen, p.x0, p.y0, p.x1, p.y0, strokeWidth, borderColor, true)
	// Left
	vector.StrokeLine(screen, p.x0, p.y0, p.x0, p.y1, strokeWidth, borderColor, true)
	// Bottom
	vector.StrokeLine(screen, p.x0, p.y1, p.x1, p.y1, strokeWidth, borderColor, true)
	// Right
	vector.StrokeLine(screen, p.x1, p.y1, p.x1, p.y0, strokeWidth, borderColor, true)
}
