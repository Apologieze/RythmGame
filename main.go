package main

import (
	"GameMusic/asset"
	"GameMusic/config"
	"GameMusic/objects"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

var (
	BACKGROUND_COLOR = color.RGBA{R: 15, G: 15, B: 15, A: 255}
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.CurrentConfig.WindowSizeX, config.CurrentConfig.WindowSizeY
}

func main() {
	objects.ColorScaleInit()
	asset.Init()
	config.Get("config.json")

	ebiten.SetWindowSize(config.CurrentConfig.WindowSizeX, config.CurrentConfig.WindowSizeY)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle(config.CurrentConfig.WindowTitle)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	game := NewGame(config.CurrentConfig)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
