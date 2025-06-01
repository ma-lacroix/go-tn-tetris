package logic

// This file contains the instructions related to the creation of a new Tetronimo

type TetronimoType int

const (
	TetronimoTypeNone TetronimoType = iota
	TetronimoType1
	TetronimoType2
	TetronimoType3
	TetronimoType4
	TetronimoType5
	TetronimoType6
	TetronimoType7
)

type BlockPieces struct {
	Pieces         map[TetronimoType][4][2]int
	ImageLocations map[TetronimoType][4][2]int
}

func NewBlockPieces() *BlockPieces {
	return &BlockPieces{
		Pieces: map[TetronimoType][4][2]int{
			TetronimoType1: {
				{4, 1}, {4, 2}, {4, 3}, {5, 3}, // L 1
			},
			TetronimoType2: {
				{6, 1}, {5, 1}, {5, 2}, {5, 3}, // L 2
			},
			TetronimoType3: {
				{4, 1}, {5, 1}, {4, 2}, {5, 2}, // Cube
			},
			TetronimoType4: {
				{4, 1}, {5, 1}, {6, 1}, {5, 2}, // T
			},
			TetronimoType5: {
				{4, 2}, {5, 2}, {5, 1}, {6, 1}, // S 1
			},
			TetronimoType6: {
				{4, 1}, {5, 1}, {5, 2}, {6, 2}, // S 2
			},
			TetronimoType7: {
				{3, 1}, {4, 1}, {5, 1}, {6, 1}, // Line
			},
		},
		ImageLocations: map[TetronimoType][4][2]int{
			TetronimoType1: {
				{1, 2}, {1, 3}, {1, 4}, {2, 4}, // L 1
			},
			TetronimoType2: {
				{10, 3}, {11, 3}, {11, 2}, {11, 1}, // L 2
			},
			TetronimoType3: {
				{6, 1}, {7, 1}, {6, 2}, {7, 2}, // Cube
			},
			TetronimoType4: {
				{5, 3}, {6, 3}, {7, 3}, {6, 4}, // T
			},
			TetronimoType5: {
				{9, 1}, {9, 2}, {8, 2}, {8, 3}, // S 1
			},
			TetronimoType6: {
				{3, 2}, {3, 3}, {4, 3}, {4, 4}, // S 2
			},
			TetronimoType7: {
				{1, 1}, {2, 1}, {3, 1}, {4, 1}, // Line
			},
		},
	}
}

func (b *BlockPieces) GenerateNewPiece(index int) [4][2]int {
	return b.Pieces[TetronimoType(index)]
}

func (b *BlockPieces) GenerateNewPieceImageLocations(index int) [4][2]int {
	return b.ImageLocations[TetronimoType(index)]
}
