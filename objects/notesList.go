package objects

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type NoteList struct {
	List                   []Note
	nextNewIndex, EndIndex int
}

func NewNoteList() NoteList {
	return NoteList{List: make([]Note, 50, 50), EndIndex: -1}
}

func (nl *NoteList) Add(pos *Vec, speed float64, color int) {
	if !nl.List[0].Alive {
		nl.nextNewIndex = 0
	}
	nl.List[nl.nextNewIndex].Set(pos, speed, ColorScales[color])
	if nl.EndIndex < nl.nextNewIndex {
		nl.EndIndex = nl.nextNewIndex
	}
	size := len(nl.List)

	for i := nl.nextNewIndex + 1; i < size; i++ {
		if !nl.List[i].Alive {
			nl.nextNewIndex = i
			return
		}
	}
	for i := 0; i < nl.nextNewIndex; i++ {
		if !nl.List[i].Alive {
			nl.nextNewIndex = i
			return
		}
	}
}

func (nl *NoteList) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprint("\nEndIndex:", nl.EndIndex))
	for i := 0; i <= nl.EndIndex; i++ {
		nl.List[i].Draw(screen)
	}
}

func (nl *NoteList) Update(destination Vec) {
	var tempEnd int = -1
	for i := 0; i <= nl.EndIndex; i++ {
		note := &nl.List[i]
		if note.Alive {
			note.UpdatePos(destination)
			tempEnd = i
		}
	}
	if tempEnd < nl.EndIndex {
		nl.EndIndex = tempEnd
	}
}
