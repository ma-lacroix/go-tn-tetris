package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/inconsolata"
	"image/color"
	"strconv"
)

type ScoreBoard struct {
	x0, y0, x1, y1, bx, by float32
	score                  int32
}

func NewScoreBoard(ScreenWidth int, ScreenHeight int) *ScoreBoard {
	offSet := float32(50)
	paddingX := float32(ScreenWidth / 5)
	paddingY := float32(ScreenWidth / 10)
	return &ScoreBoard{
		x0:    paddingX + offSet*2.5,
		y0:    paddingY + offSet*10.0,
		x1:    float32(ScreenWidth) - paddingX - offSet*1.5,
		y1:    float32(ScreenHeight) - paddingY + offSet*0.4,
		bx:    ((float32(ScreenWidth) - paddingX - offSet*1.5) - (paddingX - offSet*1.2)) / cols,
		by:    ((float32(ScreenHeight) - paddingY + offSet*0.4) - (paddingY + offSet*4)) / rows,
		score: 0,
	}
}

func (s *ScoreBoard) Update(newLines int32) {
	s.score = newLines
}

func (s *ScoreBoard) DrawBackground(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, s.x0+80, s.y0, s.x1*0.50, s.y1*0.15,
		color.RGBA{210, 230, 245, 0xFF}, false)
}

func (s *ScoreBoard) Draw(screen *ebiten.Image) {
	s.DrawBackground(screen)
	text.Draw(screen, strconv.Itoa(int(s.score)), inconsolata.Bold8x16, int(s.x0)+100,
		int(s.y0)+50, color.RGBA{20, 20, 30, 255})
}
