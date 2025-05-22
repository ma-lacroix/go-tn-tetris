package logic

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	screenWidth  = 240
	screenHeight = 240
	numSquares   = 100
	padding      = 20
)

var (
	mplusFaceSource *text.GoTextFaceSource
	mplusNormalFace *text.GoTextFace
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	mplusNormalFace = &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}
	rand.Seed(time.Now().UnixNano())
}

type Game struct {
	Squares     []Square
	Speed       float32
	Current     int // useless now
	OutOfBounds int
}

type Square struct {
	x           float32
	y           float32
	size        float32
	color       color.Color
	antialias   bool
	outOfBounds bool
}

func (s *Square) DrawSquare(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, s.x, s.y, s.size, s.size, s.color, s.antialias)
}

func (g *Game) DrawGameLimits(screen *ebiten.Image) {
	vector.StrokeLine(screen, padding, padding, screenWidth-padding, padding, 1,
		color.RGBA{20, 200, 20, 0xff}, true)
	vector.StrokeLine(screen, padding, padding, padding, screenHeight-padding, 1,
		color.RGBA{20, 200, 20, 0xff}, true)
	vector.StrokeLine(screen, padding, screenHeight-padding, screenWidth-padding, screenHeight-padding, 1,
		color.RGBA{20, 200, 20, 0xff}, true)
	vector.StrokeLine(screen, screenWidth-padding, screenHeight-padding, screenWidth-padding, padding, 1,
		color.RGBA{20, 200, 20, 0xff}, true)
}

func (g *Game) DrawWinningMessage(screen *ebiten.Image) {
	{
		const x, y = 20, 20
		op := &text.DrawOptions{}
		op.GeoM.Translate(x, y)
		op.LineSpacing = mplusNormalFace.Size * 1.5
		text.Draw(screen, "You won!", mplusNormalFace, op)
	}
}

func (g *Game) CheckCollision(sq1 Square, sq2 Square) bool {
	return math.Abs(float64(sq1.x-sq2.x)) < 15.0 && math.Abs(float64(sq1.y-sq2.y)) < 15.0
}

func (g *Game) HandleCollision() {
	for i := 1; i < len(g.Squares); i++ {
		if g.CheckCollision(g.Squares[g.Current], g.Squares[i]) {
			g.Squares[i].x += (g.Squares[i].x - g.Squares[g.Current].x) * 0.1
			g.Squares[i].y += (g.Squares[i].y - g.Squares[g.Current].y) * 0.1
			if IsOutBounds(g.Squares[i]) {
				g.OutOfBounds++
				g.Squares[i].x += (g.Squares[i].x - g.Squares[g.Current].x) * 3.0
				g.Squares[i].y += (g.Squares[i].y - g.Squares[g.Current].y) * 3.0
			}
		}
	}
}

func IsOutBounds(s Square) bool {
	return (s.x) < padding || s.x > screenWidth-s.size*1.5 || s.y < padding || s.y > screenHeight-s.size*1.5
}

func IsInBounds(s Square, newX float32, newY float32) bool {
	return (s.x+newX) > padding && (s.x+newX) < screenWidth-s.size*2 && (s.y+newY) > padding && (s.y+newY) < screenHeight-s.size*2
}

func (g *Game) Reset() {
	g.Squares = InitializeSquares(numSquares)
	g.OutOfBounds = 0
}

func InitializeSquares(amount int) []Square {
	listOfSquares := make([]Square, 0)
	// that's the player square
	listOfSquares = append(listOfSquares, Square{screenWidth / 2, screenWidth / 2, 15.0,
		color.RGBA{255, 100, 80, 255}, true, false})
	for i := 0; i < amount; i++ {
		listOfSquares = append(listOfSquares, Square{screenWidth * rand.Float32() * 0.9, screenWidth * rand.Float32() * 0.9, 15.0,
			color.RGBA{76, 76, 76, 255}, true, false})
	}
	return listOfSquares
}

func (g *Game) Update() error {
	g.HandleCollision()
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if IsInBounds(g.Squares[g.Current], -g.Speed, 0.0) {
			g.Squares[g.Current].x += -g.Speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if IsInBounds(g.Squares[g.Current], g.Speed, 0.0) {
			g.Squares[g.Current].x += g.Speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if IsInBounds(g.Squares[g.Current], 0.0, -g.Speed) {
			g.Squares[g.Current].y += -g.Speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if IsInBounds(g.Squares[g.Current], 0.0, g.Speed) {
			g.Squares[g.Current].y += g.Speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.Reset()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawGameLimits(screen)
	for _, square := range g.Squares {
		square.DrawSquare(screen)
	}
	if g.OutOfBounds == len(g.Squares)-1 {
		g.DrawWinningMessage(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
