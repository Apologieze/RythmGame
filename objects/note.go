package objects

import (
	"GameMusic/asset"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

const ScaleMultiplier = 0.2
const ScaleInverse = 1 / ScaleMultiplier

type Note struct {
	image *ebiten.Image
	pos   gmath.Pos
	color ebiten.ColorScale
	speed float64
	Alive bool
}

func NewNote(pos *Vec, speed float64) Note {
	noteSize := asset.NoteImage.Bounds().Size()
	return Note{
		image: asset.NoteImage,
		pos:   gmath.Pos{pos, Vec{-float64(noteSize.X) / 2, -float64(noteSize.Y) / 2}},
		color: ColorScales[2],
		speed: speed,
		Alive: true,
	}
}

func (n *Note) Set(pos *Vec, speed float64, color ebiten.ColorScale) {
	if n.image == nil {
		n.image = asset.NoteImage
		n.pos.Offset = Vec{-float64(n.image.Bounds().Size().X) / 2, -float64(n.image.Bounds().Size().Y) / 2}
	}
	n.pos.Base = pos
	n.speed = speed
	n.color = color
	n.Alive = true
}

func (n *Note) Draw(screen *ebiten.Image) {
	if !n.Alive {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(n.pos.Offset.X, n.pos.Offset.X)
	op.GeoM.Scale(ScaleMultiplier, ScaleMultiplier)
	op.GeoM.Translate(n.pos.Base.X, n.pos.Base.Y)
	op.ColorScale = n.color
	screen.DrawImage(n.image, op)
}

func (n *Note) UpdatePos(destination Vec) {
	v := *n.pos.Base
	direction := destination.Sub(v)
	dist := direction.Len()
	if dist <= 75 {
		fmt.Println("Note reached destination")
		n.Alive = false
		return
	}
	//fmt.Println(dist)
	add := direction.Divf(dist).Mulf(n.speed * ScaleInverse)
	fmt.Println(add.Len())
	n.pos.Base.X += add.X
	n.pos.Base.Y += add.Y
}
