package logic

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"log"
	"os"
)

type SoundBank struct {
	ctx     *audio.Context
	sfxData map[string][]byte // raw sound data
}

func NewSoundBank(ctx *audio.Context) *SoundBank {
	soundPaths := [14]string{"n_all_right", "n_explode2", "n_impact1", "n_pause", "n_switch",
		"n_enter", "n_gameOver", "n_impact2", "n_rotate", "n_yyy", "n_explode1", "n_good",
		"n_onbc", "n_start"}
	sfxData := make(map[string][]byte)
	for sound := range soundPaths {
		path := fmt.Sprintf("../media/sound/%s.wav", sound)
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Failed to load %s: %v", path, err)
		}
		key := fmt.Sprintf("%d", sound)
		sfxData[key] = data
	}
	return &SoundBank{
		ctx:     ctx,
		sfxData: sfxData,
	}
}

func (sb *SoundBank) Play(name string) {
	data, ok := sb.sfxData[name]
	if !ok {
		log.Printf("Sound %s not found", name)
		return
	}
	stream, err := wav.DecodeWithoutResampling(bytes.NewReader(data))
	if err != nil {
		log.Printf("Failed to decode %s: %v", name, err)
		return
	}
	player, err := sb.ctx.NewPlayer(stream)
	if err != nil {
		log.Printf("Failed to create player for %s: %v", name, err)
		return
	}
	player.Play()
}
