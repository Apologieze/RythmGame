package main

import (
	"GameMusic/config"
	eInput "github.com/quasilyte/ebitengine-input"
)

type Game struct {
	system      System
	inputSystem eInput.System
}

func NewGame(currentConfig config.Config) *Game {
	game := Game{}

	keymap := config.InitKeymap()
	game.inputSystem.Init(eInput.SystemConfig{DevicesEnabled: eInput.AnyDevice})
	game.system = NewSystem(currentConfig, game.inputSystem.NewHandler(0, keymap))

	return &game
}
