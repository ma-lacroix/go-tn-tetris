package logic

// This file contains the instructions related to the Player's active Tetronimo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"time"
)

type PlayerPiece struct {
	position       [4][2]int
	color          color.Color
	lockDelayTimer time.Time
}

func NewPlayerPiece(tetronimo [4][2]int, colorValues [3]int) *PlayerPiece {
	return &PlayerPiece{
		position: tetronimo,
		color:    color.RGBA{uint8(colorValues[0]), uint8(colorValues[1]), uint8(colorValues[2]), 255},
	}
}

func GetPieceMinMaxValues(newPos [4][2]int) [4]int {
	minX := 1000
	maxX := 0
	minY := 1000
	maxY := 0
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

func (pp *PlayerPiece) Rotation(col int, row int) {
	origin := pp.position[0]

	for i, block := range pp.position {
		dx := block[0] - origin[0]
		dy := block[1] - origin[1]
		pp.position[i][0] = origin[0] + dy
		pp.position[i][1] = origin[1] - dx
	}
	pp.AdjustRotationPosition(col, row)
}

func (pp *PlayerPiece) AdjustRotationPosition(col int, row int) {
	pos := pp.position
	minMaxValues := GetPieceMinMaxValues(pos)
	move := [2]int{0, 0}
	if minMaxValues[0] < 0 {
		move[0] = 0 - minMaxValues[0]
	}
	if minMaxValues[1] < 0 {
		move[1] = 0 - minMaxValues[1]
	}
	if minMaxValues[2] > col {
		move[0] = col - minMaxValues[2] - 2
	}
	if minMaxValues[3] > row {
		move[1] = row - minMaxValues[3] - 2
	}
	pp.UpdatePlayerPiece(move)
}

func (pp *PlayerPiece) DetectPlayingAreaCollision(newPos [4][2]int, col int, row int) bool {
	minMaxValues := GetPieceMinMaxValues(newPos)
	return minMaxValues[0] >= 0 && minMaxValues[1] >= 0 && minMaxValues[2] < col && minMaxValues[3] < row
}

func (pp *PlayerPiece) DetectFallenPiecesCollision(newPos [4][2]int, grid *[rows][cols]bool) bool {
	for _, pos := range newPos {
		if !grid[pos[1]][pos[0]] {
			return false
		}
	}
	return true
}

func (pp *PlayerPiece) CollisionDetection(newMove [2]int, col int, row int, grid *[rows][cols]bool) bool {
	var moved [4][2]int
	for i, pos := range pp.position {
		moved[i][0] = pos[0] + newMove[0]
		moved[i][1] = pos[1] + newMove[1]
	}
	return pp.DetectPlayingAreaCollision(moved, col, row) && pp.DetectFallenPiecesCollision(moved, grid)
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
