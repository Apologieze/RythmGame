package asset

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var NoteImage *ebiten.Image

func GetImage(fileName string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile("asset/" + fileName)
	if err != nil {
		log.Fatal("Problem while loading particle image: ", err)
	}
	return image
}

func Init() {
	NoteImage = GetImage("note.png")
}
