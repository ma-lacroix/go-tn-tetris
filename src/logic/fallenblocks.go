package logic

// This source file contains the instructions related to Tetronimo blocks that are at the bottom of the playing area:
// holding them in memory, detecting complete lines, removing complete lines and handling the animations

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
)

const (
	amplitudeY = 100.0
	amplitudeX = 20.0
)

type FallenBlock struct {
	x0, y0, bx, by float32
	imagePositions [2]int
	color          color.Color
	alpha          float32
	direction      float32
	rotation       float64
}

type FallenBlocks struct {
	fallenBlocks     []FallenBlock
	blocksToAnimate  []FallenBlock
	rowsRemoved      int32
	blockPiecesImage *ebiten.Image
}

func NewFallenBlocks() *FallenBlocks {
	blockPiecesImage := loadImage("../media/images/p_tetris_blocks_1.png")
	return &FallenBlocks{
		rowsRemoved:      0,
		blockPiecesImage: blockPiecesImage,
	}
}

func Randomizer() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}

func (f *FallenBlocks) ResetRowsToRemove() {
	f.rowsRemoved = 0
}

func (f *FallenBlocks) UpdateBlocks(playerPos [4][2]int, imagePositions [4][2]int, areaCoordinates [4]float32, color color.Color, rotation float64) {
	var direction float32
	alpha := Randomizer()
	if alpha < 0.5 {
		direction = 1
	} else {
		direction = -1
	}
	for i := 0; i < len(playerPos); i++ {
		f.fallenBlocks = append(f.fallenBlocks, FallenBlock{
			float32(playerPos[i][0])*areaCoordinates[2] + areaCoordinates[0],
			float32(playerPos[i][1])*areaCoordinates[3] + areaCoordinates[1],
			areaCoordinates[2],
			areaCoordinates[3],
			imagePositions[i],
			color,
			Randomizer(),
			direction,
			rotation,
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
		for i := range f.blocksToAnimate {
			f.blocksToAnimate[i].alpha += 0.01
			x := float32(amplitudeX * math.Sin(-float64(f.blocksToAnimate[i].alpha*Randomizer())))
			y := float32(amplitudeY * math.Sin(float64(f.blocksToAnimate[i].alpha)))
			f.blocksToAnimate[i].x0 += x * f.blocksToAnimate[i].direction
			f.blocksToAnimate[i].y0 += y
			f.blocksToAnimate[i].rotation += float64(Randomizer())
		}
	}
}

func (f *FallenBlocks) DrawExplodingBlocks(screen *ebiten.Image) {
	if len(f.blocksToAnimate) != 0 {
		for _, block := range f.blocksToAnimate {
			vector.DrawFilledRect(screen, block.x0, block.y0, block.bx, block.by, block.color, true)
			f.AddPieceTexture(screen, block)
		}
	}
}

func (f *FallenBlocks) AddPieceTexture(screen *ebiten.Image, block FallenBlock) {
	tileWidth := blockTexturesX / blockTexturesLenX
	tileHeight := blockTexturesY / blockTexturesLenY
	sx := tileWidth * block.imagePositions[0]
	sy := tileHeight * block.imagePositions[1] // ✅ fixed index
	rect := image.Rect(sx, sy, sx+tileWidth, sy+tileHeight)
	cropped := f.blockPiecesImage.SubImage(rect).(*ebiten.Image)
	op := &ebiten.DrawImageOptions{}
	w, h := cropped.Size()
	scaleX := blockImageScaleX
	scaleY := blockImageScaleY
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(-float64(w)*scaleX/2, -float64(h)*scaleY/2)
	op.GeoM.Rotate(block.rotation)
	op.GeoM.Translate(
		float64(block.x0)+float64(block.bx)/2,
		float64(block.y0)+float64(block.by)/2,
	)
	screen.DrawImage(cropped, op)
}

func (f *FallenBlocks) Draw(screen *ebiten.Image) {
	for _, block := range f.fallenBlocks {
		vector.DrawFilledRect(screen, block.x0, block.y0, block.bx, block.by, block.color, true)
		f.AddPieceTexture(screen, block)
	}
}
