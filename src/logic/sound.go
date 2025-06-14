package logic

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"log"
)

type SoundBank struct {
	ctx     *audio.Context
	sfxData map[string][]byte
}

//go:embed media/sound/*
var soundFS embed.FS

func NewSoundBank(ctx *audio.Context) *SoundBank {
	soundPaths := []string{
		"n_all_right", "n_explode2", "n_impact1", "n_pause", "n_switch",
		"n_enter", "n_gameOver", "n_impact2", "n_rotate", "n_yyy",
		"n_explode1", "n_good", "n_onbc", "n_start", "n_afterburner",
	}
	sfxData := make(map[string][]byte)
	for _, sound := range soundPaths {
		path := fmt.Sprintf("media/sound/%s.wav", sound)
		data, err := soundFS.ReadFile(path)
		if err != nil {
			log.Fatalf("Failed to load embedded sound %s: %v", path, err)
		}
		sfxData[sound] = data
	}
	return &SoundBank{
		ctx:     ctx,
		sfxData: sfxData,
	}
}

func (sb *SoundBank) Play(name string, volume float64) {
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
	player.SetVolume(volume)
	player.Play()
}
