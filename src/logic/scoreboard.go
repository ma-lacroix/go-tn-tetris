package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"image/color"
	"strconv"
)

type ScoreBoard struct {
	x0, y0, x1, y1, bx, by float32
	score                  int32
	backgroundImage        *ebiten.Image
	font                   font.Face
}

func NewScoreBoard(ScreenWidth int, ScreenHeight int) *ScoreBoard {
	offSet := float32(50)
	paddingX := float32(ScreenWidth / 5)
	paddingY := float32(ScreenWidth / 10)
	backgroundImage := loadImage("../media/images/b_score.png")
	font := LoadFont("../media/font/Excludedi.ttf", 30)
	return &ScoreBoard{
		x0:              paddingX + offSet*2.7,
		y0:              paddingY + offSet*9.85,
		x1:              float32(ScreenWidth) - paddingX - offSet*2.0,
		y1:              float32(ScreenHeight) - paddingY + offSet*0.4,
		bx:              ((float32(ScreenWidth) - paddingX - offSet*1.5) - (paddingX - offSet*1.2)) / cols,
		by:              ((float32(ScreenHeight) - paddingY + offSet*0.4) - (paddingY + offSet*4)) / rows,
		score:           0,
		backgroundImage: backgroundImage,
		font:            font,
	}
}

func (s *ScoreBoard) Update(newLines int32) {
	s.score = newLines
}

func (s *ScoreBoard) DrawBackground(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, s.x0+80, s.y0, s.x1*0.45, s.y1*0.12,
		color.RGBA{210, 230, 245, 0xFF}, false)
}

func (s *ScoreBoard) Draw(screen *ebiten.Image) {
	s.DrawBackground(screen)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(imageScaleX, imageScaleY)
	op.GeoM.Translate(float64(s.x0+80), float64(s.y0-3))
	screen.DrawImage(s.backgroundImage, op)
	text.Draw(screen, strconv.Itoa(int(s.score)), s.font, int(s.x0)+120,
		int(s.y0)+70, color.RGBA{20, 20, 30, 255})
}
