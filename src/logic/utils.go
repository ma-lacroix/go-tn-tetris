package logic

import (
	"bytes"
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"log"
	"math/rand"
	"os"
	"time"
)

var imagesFS embed.FS

func loadImage(path string) *ebiten.Image {
	data, err := imagesFS.ReadFile("../media/images/" + path)
	if err != nil {
		log.Fatalf("failed to read image %s: %v", path, err)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatalf("failed to decode image %s: %v", path, err)
	}

	return ebiten.NewImageFromImage(img)
}

func LoadFont(path string, size float64) font.Face {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
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
