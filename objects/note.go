package objects

import (
	"GameMusic/asset"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

const ScaleMultiplier = 0.2

//const ScaleInverse = 1 / ScaleMultiplier

type Note struct {
	image        *ebiten.Image
	Pos          gmath.Pos
	color        ebiten.ColorScale
	Speed        float64
	Alive        bool
	Movement     Vec
	tickToDie    int
	expectedTime int64
}

func NewNote(pos *Vec, speed float64, time int64) Note {
	noteSize := asset.NoteImage.Bounds().Size()
	n := Note{
		image:        asset.NoteImage,
		Pos:          gmath.Pos{pos, Vec{-float64(noteSize.X) / 2, -float64(noteSize.Y) / 2}},
		color:        ColorScales[1],
		Speed:        speed,
		Alive:        true,
		expectedTime: time,
	}
	n.setMovement(CenterScreen)
	return n
}

func (n *Note) Set(newNote *Note) {
	if n.image == nil {
		n.image = asset.NoteImage
	}
	n.Pos = newNote.Pos
	n.Speed = newNote.Speed
	n.color = newNote.color
	n.Movement = newNote.Movement
	n.tickToDie = newNote.tickToDie
	n.Alive = true
	n.expectedTime = newNote.expectedTime
}

func (n *Note) Draw(screen *ebiten.Image) {
	if !n.Alive {
		return
	}
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(n.Pos.Offset.X, n.Pos.Offset.X)
	//op.GeoM.Scale(ScaleMultiplier, ScaleMultiplier)
	op.GeoM.Translate(n.Pos.Base.X+n.Pos.Offset.X, n.Pos.Base.Y+n.Pos.Offset.Y)
	op.ColorScale = n.color
	screen.DrawImage(n.image, op)
}

/*func (n *Note) UpdatePos(destination Vec) {
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
}*/

func (n *Note) Update() {
	if n.tickToDie <= 0 {
		n.Alive = false
		return
	}
	n.tickToDie--
	n.Pos.Base.X += n.Movement.X
	n.Pos.Base.Y += n.Movement.Y
}

func (n *Note) setMovement(destination Vec) {
	v := *n.Pos.Base
	direction := destination.Sub(v)
	dist := direction.Len()
	n.Movement = direction.Divf(dist).Mulf(n.Speed)
	n.tickToDie = int((v.DistanceTo(destination) - 75) / n.Speed)
}
