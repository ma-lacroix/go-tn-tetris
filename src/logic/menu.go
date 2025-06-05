package logic

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

var (
	menuOptions = []string{"Start Game", "Options", "Quit"}
	selected    = 0
)

type Menu struct {
	isActive        bool
	backgroundImage *ebiten.Image
}

func NewMenu() *Menu {
	backgroundImage := loadImage("../media/images/b_intro.png")
	return &Menu{
		isActive:        true,
		backgroundImage: backgroundImage,
	}
}

func (m *Menu) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		selected--
		if selected < 0 {
			selected = len(menuOptions) - 1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		selected++
		if selected >= len(menuOptions) {
			selected = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		switch selected {
		case 0:
			fmt.Println("Starting game...")
			m.isActive = false
			// StartGame() or set a state variable
		case 1:
			fmt.Println("Options selected")
		case 2:
			fmt.Println("Quitting...")
		}
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
		text.Draw(screen, option, basicfont.Face7x13, 100, 200+i*30, col)
	}
}
