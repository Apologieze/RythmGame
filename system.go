package main

import (
	"GameMusic/config"
	"GameMusic/musicPlayer"
	"GameMusic/objects"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	eInput "github.com/quasilyte/ebitengine-input"
	"log"
	"math/rand/v2"
)

type System struct {
	line         objects.Line
	player       objects.Player
	notes        objects.NoteList
	audioContext *audio.Context
	audioPlayer  *audio.Player
	input        *eInput.Handler
	rectangle    objects.Rectangle
}

func NewSystem(config config.Config, input *eInput.Handler) System {
	s := System{
		audioContext: audio.NewContext(musicPlayer.SampleRate),
		line:         objects.Line{0, 200, float32(config.WindowSizeX), 200},
		player:       objects.NewPlayer(config, input),
		notes:        objects.NewNoteList(),
		input:        input,
		rectangle:    objects.Rectangle{420, 0, 1080, 1080},
	}
	s.notes.Add(&objects.Vec{X: 100, Y: 100}, 1, 1)
	s.notes.Add(&objects.Vec{X: 800, Y: 0}, 1, 2)
	s.notes.Add(&objects.Vec{X: 400, Y: 1000}, 1, 1)

	var err error
	s.audioPlayer, err = musicPlayer.PlayMP3(s.audioContext, "asset/Geoxor.mp3")
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
	if s.input.ActionIsJustPressed(config.ActionPlayPause) {
		if s.audioPlayer.IsPlaying() {
			s.audioPlayer.Pause()
		} else {
			s.audioPlayer.Play()
		}
	}
	if s.input.ActionIsJustPressed(config.ActionClickLeft) {
		s.notes.Add(randomVec(s.rectangle), 2, 2)
	}
	s.notes.Update(s.player.DstPost)
	s.player.Update()
}

func randomVec(rec objects.Rectangle) *objects.Vec {
	xOrY := rand.IntN(2) != 0
	bound := rand.IntN(2) != 0
	if xOrY {
		x := rec.X + float64(rand.IntN(int(rec.Width)+1))
		y := rec.Y
		if !bound {
			y = rec.Y + rec.Height
		}
		return &objects.Vec{X: x, Y: y}
	}

	y := rec.Y + float64(rand.IntN(int(rec.Height)+1))
	x := rec.X
	if !bound {
		x = rec.X + rec.Width
	}
	return &objects.Vec{X: x, Y: y}
}
