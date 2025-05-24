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
)

type BlockPieces struct {
	Pieces map[TetronimoType]*[4][2]int
}

func NewBlockPieces() *BlockPieces {
	return &BlockPieces{
		Pieces: map[TetronimoType]*[4][2]int{
			TetronimoType1: &([4][2]int{
				{4, 1}, {5, 1}, {5, 2}, {5, 3}, // rotation 0
			}),
		},
	}
}

func (b *BlockPieces) GenerateNewPiece(index int) *[4][2]int {
	t := TetronimoType(index)
	return b.Pieces[t]
}
