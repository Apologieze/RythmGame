package objects

import (
	"fmt"
	"github.com/Waffle-osu/osu-parser/osu_parser"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math/rand/v2"
)

const TickTime float64 = 60. / 1000.

type NoteList struct {
	List                   []Note
	nextNewIndex, EndIndex int
	ElapsedTime            int64
	AllNotes               map[int64][]Note
	audio                  *audio.Player
}

func NewNoteList(audio *audio.Player) NoteList {
	return NoteList{List: make([]Note, 50, 50), EndIndex: -1, AllNotes: make(map[int64][]Note), audio: audio}
}

func (nl *NoteList) Add(note *Note) {
	if !nl.List[0].Alive {
		nl.nextNewIndex = 0
	}
	nl.List[nl.nextNewIndex].Set(note)
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

func (nl *NoteList) Update() {
	var tempEnd int = -1
	for i := 0; i <= nl.EndIndex; i++ {
		note := &nl.List[i]
		if note.Alive {
			note.Update()
			if !note.Alive {
				playSound(nl.audio)
			}
			tempEnd = i
		}
	}
	if tempEnd < nl.EndIndex {
		nl.EndIndex = tempEnd
	}
}

func (nl *NoteList) InitNoteList(file *osu_parser.OsuFile, rec Rectangle, centerScreen Vec) {
	//var deltaUpdate float64 = 1. / 60.
	list := file.HitObjects.List
	lenghtHit := len(list)
	var defaultOffset int64 = int64(float64(file.General.AudioLeadIn)*TickTime) + 200
	var speed float64 = 7
	for i := 0; i < lenghtHit; i++ {
		/*if list[i].Type&osu_parser.HitObjectTypeCircle == 0 {
			continue
		}*/
		hitObj := list[i]
		tempVec := RandomVec(rec)
		var steps float64 = ((tempVec.DistanceTo(centerScreen) - 75) / speed)
		var startTick int64 = int64((hitObj.Time*TickTime)-steps) + defaultOffset
		fmt.Println(hitObj.Time * TickTime)
		nl.AllNotes[startTick] = append(nl.AllNotes[startTick], NewNote(tempVec, speed))
	}
}

func RandomVec(rec Rectangle) *Vec {
	xOrY := rand.IntN(2) != 0
	bound := rand.IntN(2) != 0
	if xOrY {
		x := rec.X + float64(rand.IntN(int(rec.Width)+1))
		y := rec.Y
		if !bound {
			y = rec.Y + rec.Height
		}
		return &Vec{X: x, Y: y}
	}

	y := rec.Y + float64(rand.IntN(int(rec.Height)+1))
	x := rec.X
	if !bound {
		x = rec.X + rec.Width
	}
	return &Vec{X: x, Y: y}
}

func playSound(player *audio.Player) {
	_ = player.Rewind()
	player.Play()
}
