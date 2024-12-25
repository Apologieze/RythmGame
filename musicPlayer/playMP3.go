package musicPlayer

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"os"
)

const SampleRate = 48000

func PlayMP3(audioContext *audio.Context, filepath string) (*audio.Player, error) {
	// Open the file
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	// Decode the MP3
	d, err := mp3.DecodeWithSampleRate(SampleRate, f)
	if err != nil {
		return nil, err
	}

	// Create a new player
	player, err := audioContext.NewPlayer(d)
	if err != nil {
		return nil, err
	}

	return player, nil
}
