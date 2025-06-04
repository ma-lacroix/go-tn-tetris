package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

const fontsize = 150
const startingX = 400
const startingY = 350

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
		Message{"GOOD!", [2]int{startingX, startingY}, false},
		Message{"ALL RIGHT!!", [2]int{startingX, startingY}, false},
		Message{"YEAH YEAH YEAH!!!", [2]int{startingX, startingY}, false},
		Message{"OW NOW BROWN COW!!!!", [2]int{startingX, startingY}, false},
	},
		font: LoadFont("../media/font/Excludedi.ttf", fontsize)}
}

func (m *Messages) ActivateMessage(index int32) {
	m.allMessages[index-1].active = true
}

func (m *Messages) MoveActiveMessage() {
	for i := 0; i < len(m.allMessages); i++ {
		if m.allMessages[i].active {
			// reset message position
			if m.allMessages[i].pos[0] <= -1500 {
				m.allMessages[i].active = false
				m.allMessages[i].pos[0] = startingX
			}
			m.allMessages[i].pos[0] -= 20
		}

	}
	for _, msg := range m.allMessages {
		if msg.active {
			if msg.pos[0] <= -200 {
				msg.active = false
			}
			msg.pos[0] -= 10
		}
	}
}

func (m *Messages) Draw(screen *ebiten.Image) {
	for _, message := range m.allMessages {
		if message.active {
			text.Draw(screen, message.msg, m.font, message.pos[0], message.pos[1],
				color.RGBA{20, 20, 30, 230})
		}
	}
}
