package logic

// This file contains the instructions related to the Player's active Tetronimo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	"math"
	"time"
)

var (
	angles = []float64{0, math.Pi / 2, math.Pi, 3 * math.Pi / 2}
)

const (
	blockTexturesX    = 1733
	blockTexturesLenX = 13
	blockTexturesY    = 800
	blockTexturesLenY = 6
)

type PlayerPiece struct {
	position         [4][2]int
	imagePositions   [4][2]int
	color            color.Color
	lockDelayTimer   time.Time
	blockPiecesImage *ebiten.Image
	rotationIndex    int
}

func NewPlayerPiece(tetronimo [4][2]int, imagePositions [4][2]int) *PlayerPiece {
	blockPiecesImage := loadImage("../media/images/p_tetris_blocks_1.png")
	return &PlayerPiece{
		position:         tetronimo,
		imagePositions:   imagePositions,
		color:            color.RGBA{90, 90, 90, 255},
		blockPiecesImage: blockPiecesImage,
		rotationIndex:    0,
	}
}

func GetPieceMinMaxValues(newPos [4][2]int) [4]int {
	minX, minY := 1000, 1000
	maxX, maxY := 0, 0
	for _, block := range newPos {
		if block[0] < minX {
			minX = block[0]
		}
		if block[1] < minY {
			minY = block[1]
		}
		if block[0] > maxX {
			maxX = block[0]
		}
		if block[1] > maxY {
			maxY = block[1]
		}
	}
	return [4]int{minX, minY, maxX, maxY}
}

func (pp *PlayerPiece) Rotation(grid *[rows][cols]bool) {
	origin := pp.position[0]
	var rotated [4][2]int

	for i, block := range pp.position {
		// Translate to origin
		dx := block[0] - origin[0]
		dy := block[1] - origin[1]
		newX := origin[0] + dy
		newY := origin[1] - dx

		rotated[i][0] = newX
		rotated[i][1] = newY
	}
	rotated = AdjustRotationPosition(rotated)
	if pp.DetectFallenPiecesCollision(rotated, grid) && pp.DetectPlayingAreaCollision(rotated) {
		pp.position = rotated
		if pp.rotationIndex == 3 {
			pp.rotationIndex = 0
		} else {
			pp.rotationIndex += 1
		}
	}
}

func AdjustRotationPosition(rotated [4][2]int) [4][2]int {
	minMaxValues := GetPieceMinMaxValues(rotated)
	move := [2]int{0, 0}
	if minMaxValues[0] < 0 {
		move[0] = 0 - minMaxValues[0]
	}
	if minMaxValues[1] < 0 {
		move[1] = 0 - minMaxValues[1]
	}
	if minMaxValues[2] >= cols {
		move[0] = cols - 1 - minMaxValues[2]
	}
	if minMaxValues[3] >= rows {
		move[1] = rows - 1 - minMaxValues[3]
	}
	for i := range rotated {
		rotated[i][0] += move[0]
		rotated[i][1] += move[1]
	}
	return rotated
}

func (pp *PlayerPiece) DetectPlayingAreaCollision(newPos [4][2]int) bool {
	minMaxValues := GetPieceMinMaxValues(newPos)
	return minMaxValues[0] >= 0 && minMaxValues[1] >= 0 && minMaxValues[2] < cols && minMaxValues[3] < rows
}

func (pp *PlayerPiece) DetectFallenPiecesCollision(newPos [4][2]int, grid *[rows][cols]bool) bool {
	for _, pos := range newPos {
		if !grid[pos[1]][pos[0]] {
			return false
		}
	}
	return true
}

func (pp *PlayerPiece) CollisionDetection(newMove [2]int, grid *[rows][cols]bool) bool {
	var moved [4][2]int
	for i, pos := range pp.position {
		moved[i][0] = pos[0] + newMove[0]
		moved[i][1] = pos[1] + newMove[1]
	}
	return pp.DetectPlayingAreaCollision(moved) && pp.DetectFallenPiecesCollision(moved, grid)
}

func (pp *PlayerPiece) ShouldLock(row int, lockDelay time.Duration, grid *[rows][cols]bool) bool {
	if pp.BottomCollisionDetection(row, grid) {
		if pp.lockDelayTimer.IsZero() {
			pp.lockDelayTimer = time.Now()
		} else if time.Since(pp.lockDelayTimer) > lockDelay {
			pp.lockDelayTimer = time.Time{} // reset for next piece
			return true
		}
	} else {
		pp.lockDelayTimer = time.Time{}
	}
	return false
}

func (pp *PlayerPiece) BottomCollisionFallenPieces(grid *[rows][cols]bool) bool {
	for _, pos := range pp.position {
		x, y := pos[0], pos[1]
		if y == rows-1 {
			continue
		}
		if !grid[y+1][x] {
			return true
		}
	}
	return false
}

func (pp *PlayerPiece) BottomCollisionDetection(row int, grid *[rows][cols]bool) bool {
	currentY := GetPieceMinMaxValues(pp.position)[3]
	return currentY == row-1 || pp.BottomCollisionFallenPieces(grid)
}

func (pp *PlayerPiece) UpdatePlayerPiece(newMove [2]int) {
	for i := range pp.position {
		pp.position[i][0] += newMove[0]
		pp.position[i][1] += newMove[1]
	}
}

func (pp *PlayerPiece) AddPiecesTextures(screen *ebiten.Image, p *PlayingArea) {
	tileWidth := blockTexturesX / blockTexturesLenX
	tileHeight := blockTexturesY / blockTexturesLenY

	for i := 0; i < len(pp.position); i++ {
		sx := tileWidth * pp.imagePositions[i][0]
		sy := tileHeight * pp.imagePositions[i][1]
		rect := image.Rect(sx, sy, sx+tileWidth, sy+tileHeight)
		cropped := pp.blockPiecesImage.SubImage(rect).(*ebiten.Image)

		op := &ebiten.DrawImageOptions{}
		w, h := cropped.Size()
		scaleX := blockImageScaleX
		scaleY := blockImageScaleY
		op.GeoM.Scale(scaleX, scaleY)
		op.GeoM.Translate(
			-float64(w)*scaleX/2,
			-float64(h)*scaleY/2,
		)
		op.GeoM.Rotate(angles[pp.rotationIndex])
		op.GeoM.Translate(
			float64(pp.position[i][0])*float64(p.bx)+float64(p.x0)+float64(p.bx)/2,
			float64(pp.position[i][1])*float64(p.by)+float64(p.y0)+float64(p.by)/2,
		)
		screen.DrawImage(cropped, op)
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
	pp.AddPiecesTextures(screen, p)
}

func (pp *PlayerPiece) DrawSuperDrop(screen *ebiten.Image, p *PlayingArea) {
	middlePoint := math.Abs(float64(pp.position[3][0]+pp.position[0][0])) / 2
	for i := 50 - 1; i > 0; i-- {
		vector.DrawFilledRect(screen,
			float32(middlePoint)*p.bx+p.x0*Randomizer(),
			float32(pp.position[2][1])*p.by+p.y0-float32(5*i),
			p.bx*1.5,
			p.by,
			color.RGBA{uint8(255 / i), uint8(1 * i), uint8(1 * i), uint8(255 / i)},
			true)
	}
}
