package main

func (g *Game) Update() error {
	g.system.Update()
	g.inputSystem.Update()
	return nil
}
