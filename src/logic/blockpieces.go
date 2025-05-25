package logic

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
	Pieces map[TetronimoType][4][2]int
}

func NewBlockPieces() *BlockPieces {
	return &BlockPieces{
		Pieces: map[TetronimoType][4][2]int{
			TetronimoType1: {
				{4, 1}, {5, 1}, {5, 2}, {5, 3}, // L 1
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
	}
}

func (b *BlockPieces) GenerateNewPiece(index int) [4][2]int {
	return b.Pieces[TetronimoType(index)]
}
