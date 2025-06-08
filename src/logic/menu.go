package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

var (
	menuOptions = []string{"Easy Peasy", "Dude, Seriously?", "Hell no. Just no."}
	selected    = 0
)

type Menu struct {
	isActive        bool
	backgroundImage *ebiten.Image
	font            font.Face
}

func NewMenu() *Menu {
	backgroundImage := loadImage("../media/images/b_intro.png")
	return &Menu{
		isActive:        true,
		backgroundImage: backgroundImage,
		font:            LoadFont("../media/font/Excludedi.ttf", fontsize/5),
	}
}

func (m *Menu) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 30, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(imageScaleX, imageScaleY)
	screen.DrawImage(m.backgroundImage, op)
	for i, option := range menuOptions {
		col := color.RGBA{20, 20, 20, 255}
		if i == selected {
			col = color.RGBA{255, 100, 0, 255} // Highlight
		}
		text.Draw(screen, option, m.font, 20, 200+i*40, col)
	}
}
