package musicPlayer

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"os"
	"strings"
)

const SampleRate = 44100

func PlayMP3(audioContext *audio.Context, filepath string) (*audio.Player, error) {
	// Open the file
	f, err := os.Open("asset/audioMap/" + filepath)
	if err != nil {
		return nil, err
	}

	// Decode the MP3
	var player *audio.Player
	if strings.HasSuffix(filepath, ".ogg") {
		d, err := vorbis.DecodeWithSampleRate(SampleRate, f)
		player, err = audioContext.NewPlayer(d)
		if err != nil {
			return nil, err
		}
	} else {
		d, err := mp3.DecodeWithSampleRate(SampleRate, f)
		player, err = audioContext.NewPlayer(d)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return player, nil
}

func GetWavHitsound(audioContext *audio.Context, filepath string) (*audio.Player, error) {
	// Open the file
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	// Decode the MP3
	d, err := wav.DecodeWithSampleRate(SampleRate, f)
	if err != nil {
		return nil, err
	}

	// Create a new player
	player, err := audioContext.NewPlayer(d)
	player.SetVolume(0.1)
	if err != nil {
		return nil, err
	}

	return player, nil
}
