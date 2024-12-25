package config

import eInput "github.com/quasilyte/ebitengine-input"

const (
	ActionClickLeft eInput.Action = iota
	ActionClickRight
	ActionPlayPause
)

func InitKeymap() eInput.Keymap {
	return eInput.Keymap{
		ActionClickLeft:  {eInput.KeyGamepadLeft, eInput.KeyLeft, eInput.KeyA},
		ActionClickRight: {eInput.KeyGamepadRight, eInput.KeyRight, eInput.KeyD},
		ActionPlayPause:  {eInput.KeySpace, eInput.KeyEnter},
	}
}
