package logic

// This source file contains the instructions related to Tetronimo blocks that are at the bottom of the playing area:
// holding them in memory, detecting complete lines, removing complete lines and handling the animations

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type FallenBlock struct {
	x0, y0, bx, by float32
	color          color.Color
}

type FallenBlocks struct {
	fallenBlocks []FallenBlock
}

func NewFallenBlocks() *FallenBlocks {
	return &FallenBlocks{}
}

func (f *FallenBlocks) UpdateBlocks(playerPos [4][2]int, areaCoordinates [4]float32, color color.Color) {
	for _, pos := range playerPos {
		f.fallenBlocks = append(f.fallenBlocks, FallenBlock{
			float32(pos[0])*areaCoordinates[2] + areaCoordinates[0],
			float32(pos[1])*areaCoordinates[3] + areaCoordinates[1],
			areaCoordinates[2],
			areaCoordinates[3],
			color,
		})
	}
}

func (f *FallenBlocks) Draw(screen *ebiten.Image) {
	for _, block := range f.fallenBlocks {
		vector.DrawFilledRect(screen, block.x0, block.y0, block.bx, block.by, block.color, true)
	}
}
