package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Line struct {
	SrcX, SrcY, DstX, DstY float32
}

func (l Line) Draw(screen *ebiten.Image) {
	vector.StrokeLine(screen, l.SrcX, l.SrcY, l.DstX, l.DstY, 2, color.White, true)
}
