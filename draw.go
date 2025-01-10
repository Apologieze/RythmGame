package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, fmt.Sprint("TPS:", ebiten.ActualTPS()))
	screen.Fill(BACKGROUND_COLOR)
	g.system.Draw(screen)
}
