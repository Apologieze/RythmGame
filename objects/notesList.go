package objects

import (
	"fmt"
	"github.com/Waffle-osu/osu-parser/osu_parser"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math"
	"math/rand/v2"
	"sort"
)

const TickPerMili float64 = 60. / 1000.
const MilliPerTick = 1000.0 / 60.0

var DefaultOffset int = int(math.Round(200.0 * MilliPerTick))

var timeDiff int

type NoteList struct {
	List                        []Note
	nextNewIndex, EndIndex      int
	AllNotes                    []Note
	hitSoundPlayer, musicPlayer *audio.Player
	Playing                     bool
	currentTimeMili, noteIndex  int
}

func NewNoteList(audio, music *audio.Player) NoteList {
	return NoteList{List: make([]Note, 50, 50), EndIndex: -1, AllNotes: make([]Note, 0, 2000), hitSoundPlayer: audio, musicPlayer: music}
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
	ebitenutil.DebugPrint(screen, fmt.Sprint("EndIndex:", nl.EndIndex, "\nTimeDiff:", timeDiff))
	for i := 0; i <= nl.EndIndex; i++ {
		nl.List[i].Draw(screen)
	}
}

func (nl *NoteList) Update() {
	if nl.Playing {
		nl.CheckAdd(nl.currentTimeMili)
	}

	var tempEnd int = -1
	for i := 0; i <= nl.EndIndex; i++ {
		note := &nl.List[i]
		if note.Alive {
			note.Update()
			if !note.Alive {
				timeDiff = nl.currentTimeMili - note.compare
				playSound(nl.hitSoundPlayer)
			}
			tempEnd = i
		}
	}
	if tempEnd < nl.EndIndex {
		nl.EndIndex = tempEnd
	}
	nl.currentTimeMili = int(nl.musicPlayer.Position().Milliseconds()) + DefaultOffset
}

func (nl *NoteList) InitNoteList(file *osu_parser.OsuFile, rec Rectangle, centerScreen Vec) {
	list := file.HitObjects.List
	lenghtHit := len(list)
	//var defaultOffset = float64(file.General.AudioLeadIn+200) * TickPerMili
	var speed float64 = 7
	var color = 1
	/*var rando = gmath.Rand{}
	rando.SetSeed(1)*/
	for i := 0; i < lenghtHit; i++ {
		if list[i].NewCombo {
			color = ((color + 1) % 3) + 1
			//color = color%2 + 1
		}
		/*if list[i].Type&osu_parser.HitObjectTypeCircle == 0 {
			continue
		}*/
		hitObj := list[i]
		tempVec := objectToVect(hitObj)
		var steps float64 = ((tempVec.DistanceTo(centerScreen) - 75.) / speed) * MilliPerTick
		var startMili = int(math.Round(hitObj.Time-steps)) + DefaultOffset
		//var startTick int64 = int64(math.Round((hitObj.Time * TickPerMili) - steps + defaultOffset))
		//fmt.Println(startMili)
		nl.AllNotes = append(nl.AllNotes, NewNote(tempVec, speed, startMili, int(hitObj.Time)+DefaultOffset, color))
	}
	sort.Slice(nl.AllNotes, func(i, j int) bool {
		return nl.AllNotes[i].expectedTime < nl.AllNotes[j].expectedTime
	})
}

func (nl *NoteList) CheckAdd(currentTime int) {
	lenghtAll := len(nl.AllNotes)
	for i := nl.noteIndex; i < lenghtAll; i++ {
		if nl.AllNotes[i].expectedTime < currentTime {
			//fmt.Println("AZAFZDFA", currentTime, nl.AllNotes[i])
			nl.Add(&nl.AllNotes[i])
			//fmt.Println(nl.AllNotes[i].expectedTime - nl.currentTimeMili)
			nl.noteIndex = i + 1
		} else {
			break
		}
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
