package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, fmt.Sprint("TPS:", ebiten.ActualTPS()))
	g.system.Draw(screen)
}
