package objects

import "github.com/hajimehoshi/ebiten/v2"

var ColorScales = [4]ebiten.ColorScale{}

func ColorScaleInit() {
	ColorScales[1].Scale(0, 0, 1, 1)
	ColorScales[2].Scale(1, 0, 0, 1)
	ColorScales[3].Scale(0.5, 0, 0.5, 1)
}
