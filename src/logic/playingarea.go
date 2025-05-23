package logic

type PlayingArea struct {
}

func NewPlayingArea() *PlayingArea {
	return &PlayingArea{}
}

func (p *PlayingArea) Init() {

}

func (p *PlayingArea) Update() {

}

func (p *PlayingArea) Draw() {

}

//func (p *PlayingArea) DrawGameLimits(screen *ebiten.Image) {
//	vector.StrokeLine(screen, padding, padding, screenWidth-padding, padding, 1,
//		color.RGBA{20, 200, 20, 0xff}, true)
//	vector.StrokeLine(screen, padding, padding, padding, screenHeight-padding, 1,
//		color.RGBA{20, 200, 20, 0xff}, true)
//	vector.StrokeLine(screen, padding, screenHeight-padding, screenWidth-padding, screenHeight-padding, 1,
//		color.RGBA{20, 200, 20, 0xff}, true)
//	vector.StrokeLine(screen, screenWidth-padding, screenHeight-padding, screenWidth-padding, padding, 1,
//		color.RGBA{20, 200, 20, 0xff}, true)
//}
