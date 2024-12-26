package objects

import (
	"GameMusic/asset"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

const ScaleMultiplier = 0.2
const ScaleInverse = 1 / ScaleMultiplier

type Note struct {
	image *ebiten.Image
	Pos   gmath.Pos
	color ebiten.ColorScale
	Speed float64
	Alive bool
}

func NewNote(pos *Vec, speed float64) Note {
	noteSize := asset.NoteImage.Bounds().Size()
	return Note{
		image: asset.NoteImage,
		Pos:   gmath.Pos{pos, Vec{-float64(noteSize.X) / 2, -float64(noteSize.Y) / 2}},
		color: ColorScales[2],
		Speed: speed,
		Alive: true,
	}
}

func (n *Note) Set(pos *Vec, speed float64, color ebiten.ColorScale) {
	if n.image == nil {
		n.image = asset.NoteImage
		n.Pos.Offset = Vec{-float64(n.image.Bounds().Size().X) / 2, -float64(n.image.Bounds().Size().Y) / 2}
	}
	n.Pos.Base = pos
	n.Speed = speed
	n.color = color
	n.Alive = true
}

func (n *Note) Draw(screen *ebiten.Image) {
	if !n.Alive {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(n.Pos.Offset.X, n.Pos.Offset.X)
	op.GeoM.Scale(ScaleMultiplier, ScaleMultiplier)
	op.GeoM.Translate(n.Pos.Base.X, n.Pos.Base.Y)
	op.ColorScale = n.color
	screen.DrawImage(n.image, op)
}

func (n *Note) UpdatePos(destination Vec) {
	v := *n.Pos.Base
	direction := destination.Sub(v)
	dist := direction.Len()
	if dist <= 75 {
		//fmt.Println("Note reached destination")
		n.Alive = false
		return
	}
	//fmt.Println(dist)
	add := direction.Divf(dist).Mulf(n.Speed)
	//fmt.Println(add.Len())
	n.Pos.Base.X += add.X
	n.Pos.Base.Y += add.Y
}
