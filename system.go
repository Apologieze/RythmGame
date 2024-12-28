package main

import (
	"GameMusic/config"
	"GameMusic/musicPlayer"
	"GameMusic/objects"
	"GameMusic/parser"
	"github.com/Waffle-osu/osu-parser/osu_parser"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	eInput "github.com/quasilyte/ebitengine-input"
	"log"
)

type System struct {
	line              objects.Line
	player            objects.Player
	notes             objects.NoteList
	audioContext      *audio.Context
	audioPlayer       *audio.Player
	input             *eInput.Handler
	rectangle         objects.Rectangle
	osuMap            *osu_parser.OsuFile
	initialTick, tick int
}

// var startTime time.Time

func NewSystem(config config.Config, input *eInput.Handler) System {
	s := System{
		audioContext: audio.NewContext(musicPlayer.SampleRate),
		line:         objects.Line{0, 30, float32(config.WindowSizeX), 30},
		player:       objects.NewPlayer(config, input),
		input:        input,
		rectangle:    objects.Rectangle{420, 0, 1080, 1080},
		osuMap:       parser.Parse("asset/map/Porter Robinson - Goodbye To A World (Monstrata) [Terminus].osu"),
	}
	//s.notes.Add(&objects.Vec{X: 100, Y: 100}, 1, 1)
	/*s.notes.Add(&objects.Vec{X: 800, Y: 0}, 1, 2)
	s.notes.Add(&objects.Vec{X: 400, Y: 1000}, 1, 1)*/

	var err error
	s.audioPlayer, err = musicPlayer.PlayMP3(s.audioContext, "asset/audioMap/Goodbye.mp3")
	hitSoundPlayer, err := musicPlayer.GetWavHitsound(s.audioContext, "asset/normal-hitclap.wav")
	s.notes = objects.NewNoteList(hitSoundPlayer, s.audioPlayer)

	s.notes.InitNoteList(s.osuMap, s.rectangle, s.player.DstPost)
	s.initialTick = 200
	//startTime = time.Now()

	if err != nil {
		log.Fatal(err)
	}
	s.audioPlayer.SetVolume(0.05)
	return s
}

func (s *System) Draw(screen *ebiten.Image) {
	s.line.Draw(screen)
	s.player.Draw(screen)
	s.notes.Draw(screen)
}

func (s *System) Update() {
	//fmt.Printf("%d\n", s.audioPlayer.Position()/time.Millisecond)
	/*if s.input.ActionIsJustPressed(config.ActionPlayPause) {
		if s.audioPlayer.IsPlaying() {
			s.audioPlayer.Pause()
		} else {
			s.audioPlayer.Play()
		}
	}*/
	s.notes.Update()
	s.player.Update()

	if s.tick == s.initialTick {
		s.audioPlayer.Play()
		s.notes.Playing = true
		s.tick += 100
	}

	if s.tick < s.initialTick {
		s.tick++
		s.notes.CheckAdd(int(float64(s.tick) * objects.MilliPerTick))
	}
	/*if s.notes.AllNotes[s.tick] != nil {
		//fmt.Println(time.Now().Sub(startTime).Milliseconds(), float64(s.tick-s.initialTick)*objects.TickPerMili)
		for i := 0; i < len(s.notes.AllNotes[s.tick]); i++ {
			s.notes.Add(&s.notes.AllNotes[s.tick][i])
		}
		//s.tick = int64(float64(time.Now().Sub(startTime).Milliseconds()) / objects.TickPerMili)
	}*/
}
