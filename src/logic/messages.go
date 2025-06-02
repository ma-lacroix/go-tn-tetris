package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

const fontsize = 100

type Message struct {
	msg    string
	pos    [2]int
	active bool
}

type Messages struct {
	allMessages [4]Message
	font        font.Face
}

func NewMessages() *Messages {
	return &Messages{allMessages: [4]Message{
		Message{"GOOD!", [2]int{480, 350}, false},
		Message{"ALL RIGHT!!", [2]int{480, 350}, false},
		Message{"YEAH YEAH YEAH!!!", [2]int{480, 350}, false},
		Message{"OW NOW BROWN COW!!!!", [2]int{480, 350}, false},
	},
		font: LoadFont("../media/font/Excludedi.ttf", 30)}
}

func (m *Messages) MoveActiveMessage() {
	for _, msg := range m.allMessages {
		if msg.active {
			if msg.pos[0] <= -200 {
				msg.active = false
			}
			msg.pos[0] -= 5
		}
	}
}

func (m *Messages) Draw(screen *ebiten.Image) {
	for _, message := range m.allMessages {
		if message.active {
			text.Draw(screen, message.msg, m.font, message.pos[0], message.pos[1],
				color.RGBA{20, 20, 30, 255})
		}
	}
}
