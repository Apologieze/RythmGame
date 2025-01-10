package config

import (
	eInput "github.com/quasilyte/ebitengine-input"
)

const (
	ActionClickLeft eInput.Action = iota
	ActionClickRight
	ActionPlayPause
	ActionChangeColorBlue
	ActionChangeColorRed
)

func InitKeymap() eInput.Keymap {
	return eInput.Keymap{
		ActionClickLeft:       {eInput.KeyLeft, eInput.KeyA, eInput.KeyMouseLeft},
		ActionClickRight:      {eInput.KeyRight, eInput.KeyD, eInput.KeyMouseRight},
		ActionPlayPause:       {eInput.KeyEnter},
		ActionChangeColorBlue: {eInput.KeyMouseLeft},
		ActionChangeColorRed:  {eInput.KeyMouseRight},
	}
}
