package logic

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"log"
	"math/rand"
	"time"
)

func loadImage(path string) *ebiten.Image {
	data, err := imageFS.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read embedded image %s: %v", path, err)
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatalf("Failed to decode image %s: %v", path, err)
	}
	return ebiten.NewImageFromImage(img)
}

func LoadFont(path string, size float64) font.Face {
	fontBytes, err := fontFS.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read embedded font %s: %v", path, err)
	}
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}
	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create font face: %v", err)
	}
	return face
}

func RandomPieceIndex() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7) + 1
}

func Randomizer() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}
