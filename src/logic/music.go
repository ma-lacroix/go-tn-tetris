package logic

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"log"
	"os"
)

type MusicBank struct {
	ctx     *audio.Context
	sfxData map[string][]byte
	players map[string]*audio.Player
}

func NewMusicBankBank(ctx *audio.Context) *MusicBank {
	soundPaths := [2]string{"s_menu", "s_playing"}
	sfxData := make(map[string][]byte)
	players := make(map[string]*audio.Player)

	for _, sound := range soundPaths {
		path := fmt.Sprintf("media/music/%s.ogg", sound)
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Failed to load %s: %v", path, err)
		}
		sfxData[sound] = data
	}

	return &MusicBank{
		ctx:     ctx,
		sfxData: sfxData,
		players: players,
	}
}

func (mb *MusicBank) Play(name string, volume float64) {
	if _, ok := mb.players[name]; ok {
		return
	}

	data, ok := mb.sfxData[name]
	if !ok {
		log.Printf("Sound %s not found", name)
		return
	}

	stream, err := vorbis.DecodeWithoutResampling(bytes.NewReader(data))
	if err != nil {
		log.Printf("Failed to decode %s: %v", name, err)
		return
	}

	player, err := mb.ctx.NewPlayer(stream)
	if err != nil {
		log.Printf("Failed to create player for %s: %v", name, err)
		return
	}

	player.SetVolume(volume)
	mb.players[name] = player
	player.Play()
}

func (mb *MusicBank) Stop(name string) {
	if player, ok := mb.players[name]; ok {
		if player.IsPlaying() {
			player.Pause()
			player.Close()
		}
		delete(mb.players, name)
	}
}
