package logic

// This source file contains the instructions related to Tetronimo blocks that are at the bottom of the playing area:
// holding them in memory, detecting complete lines, removing complete lines and handling the animations

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	"math/rand"
	"time"
)

const (
	amplitudeY = 100.0
	amplitudeX = 50.0
)

type FallenBlock struct {
	x0, y0, bx, by float32
	color          color.Color
	alpha          float32
	direction      float32
}

type FallenBlocks struct {
	fallenBlocks    []FallenBlock
	blocksToAnimate []FallenBlock
	rowsRemoved     int32
}

func NewFallenBlocks() *FallenBlocks {
	return &FallenBlocks{rowsRemoved: 0}
}

func Randomizer() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}

func (f *FallenBlocks) UpdateBlocks(playerPos [4][2]int, areaCoordinates [4]float32, color color.Color) {
	var direction float32
	alpha := Randomizer()
	if alpha < 0.5 {
		direction = 1
	} else {
		direction = -1
	}
	for _, pos := range playerPos {
		f.fallenBlocks = append(f.fallenBlocks, FallenBlock{
			float32(pos[0])*areaCoordinates[2] + areaCoordinates[0],
			float32(pos[1])*areaCoordinates[3] + areaCoordinates[1],
			areaCoordinates[2],
			areaCoordinates[3],
			color,
			Randomizer(),
			direction,
		})
	}
	rowsToRemove := f.findCompleteRows()
	if len(rowsToRemove) > 0 {
		minKeyValue := findMinKeyValue(rowsToRemove)
		f.removeCompleteRows(rowsToRemove)
		f.moveBlocksDownwards(minKeyValue, len(rowsToRemove), areaCoordinates[3])
		f.rowsRemoved += int32(len(rowsToRemove))
	}
}

func findMinKeyValue(rowsToRemove map[float32]bool) float32 {
	var minKeyValue float32
	first := true
	for k := range rowsToRemove {
		if first || k < minKeyValue {
			minKeyValue = k
			first = false
		}
	}
	return minKeyValue
}

func (f *FallenBlocks) moveBlocksDownwards(minRowValue float32, numDrops int, blockSize float32) {
	for i := range f.fallenBlocks {
		if f.fallenBlocks[i].y0 < minRowValue {
			f.fallenBlocks[i].y0 += blockSize * float32(numDrops)
		}
	}
}

func (f *FallenBlocks) removeCompleteRows(rowsToDelete map[float32]bool) {
	newBlocks := make([]FallenBlock, 0, len(f.fallenBlocks))
	newBlocksToAnimate := make([]FallenBlock, 0, len(f.blocksToAnimate))
	for _, block := range f.fallenBlocks {
		if !rowsToDelete[block.y0] {
			newBlocks = append(newBlocks, block)
		} else {
			newBlocksToAnimate = append(newBlocksToAnimate, block)
		}
	}
	f.fallenBlocks = newBlocks
	f.blocksToAnimate = newBlocksToAnimate
}

func (f *FallenBlocks) findCompleteRows() map[float32]bool {
	rowsToDelete := make(map[float32]bool)
	inv := make(map[float32]int)
	for _, block := range f.fallenBlocks {
		inv[block.y0]++
		if inv[block.y0] == cols {
			rowsToDelete[block.y0] = true
		}
	}
	return rowsToDelete
}

func (f *FallenBlocks) removeOutOfBoundBlocks() {
	newBlocksToAnimate := make([]FallenBlock, 0, len(f.blocksToAnimate))
	for _, block := range f.blocksToAnimate {
		if block.y0 < 1000 {
			newBlocksToAnimate = append(newBlocksToAnimate, block)
		}
	}
	f.blocksToAnimate = newBlocksToAnimate
}

func (f *FallenBlocks) MoveExplodingBlocks() {
	f.removeOutOfBoundBlocks()
	if len(f.blocksToAnimate) != 0 {
		for i, _ := range f.blocksToAnimate {
			f.blocksToAnimate[i].alpha += 0.00005
			f.blocksToAnimate[i].x0 += float32(amplitudeX*math.Cos(float64(f.blocksToAnimate[i].alpha))) * Randomizer() * f.blocksToAnimate[i].direction
			f.blocksToAnimate[i].y0 += float32(amplitudeY*math.Sin(float64(f.blocksToAnimate[i].alpha))) * Randomizer()
		}
	}
}

func (f *FallenBlocks) DrawExplodingBlocks(screen *ebiten.Image) {
	if len(f.blocksToAnimate) != 0 {
		for _, block := range f.blocksToAnimate {
			vector.DrawFilledRect(screen, block.x0, block.y0, block.bx, block.by, block.color, true)
		}
	}
}

func (f *FallenBlocks) Draw(screen *ebiten.Image) {
	for _, block := range f.fallenBlocks {
		vector.DrawFilledRect(screen, block.x0, block.y0, block.bx, block.by, block.color, true)
	}
}
