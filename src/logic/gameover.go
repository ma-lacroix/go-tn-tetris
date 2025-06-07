package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type GameOver struct {
	isActive          bool
	backgroundImage   *ebiten.Image
	backgroundLayover *ebiten.Image
	ScreenWidth       int
	ScreenHeight      int
	x0, y0            float64
	rotationAngle     float64
}

func NewGameOver(ScreenWidth int, ScreenHeight int) *GameOver {
	backgroundImage := loadImage("../media/images/b_game_over_layer.png")
	backgroundLayover := loadImage("../media/images/b_game_over_text.png")
	return &GameOver{
		isActive:          false,
		backgroundImage:   backgroundImage,
		backgroundLayover: backgroundLayover,
		ScreenWidth:       ScreenWidth,
		ScreenHeight:      ScreenHeight,
		x0:                0,
		y0:                0,
		rotationAngle:     0,
	}
}

func (g *GameOver) FlickerBackground() {
	g.rotationAngle += 0.5
	cx, cy := float64(-100), float64(-100)
	radius := 60.0

	// Compute new position
	g.x0 = cx + math.Cos(g.rotationAngle)*radius
	g.y0 = cy + math.Sin(g.rotationAngle)*radius
}

func (g *GameOver) Draw(screen *ebiten.Image) {
	opb := &ebiten.DrawImageOptions{}
	opb.GeoM.Scale(imageScaleX+0.2, imageScaleY+0.2)
	opb.GeoM.Translate(g.x0, g.y0)
	screen.DrawImage(g.backgroundImage, opb)
	opc := &ebiten.DrawImageOptions{}
	opc.GeoM.Scale(imageScaleX, imageScaleY)
	screen.DrawImage(g.backgroundLayover, opc)
}
