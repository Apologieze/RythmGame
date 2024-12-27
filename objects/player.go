package objects

import (
	"GameMusic/asset"
	"GameMusic/config"
	"github.com/hajimehoshi/ebiten/v2"
	eInput "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
	"math"
)

const PiBy2 float64 = math.Pi / 2

type Pos = gmath.Pos
type Vec = gmath.Vec

var CenterScreen Vec

type Player struct {
	Circle, Reflector  *ebiten.Image
	circleOption       *ebiten.DrawImageOptions
	angle              float64
	DstPost, CenterPos Vec
	input              *eInput.Handler
	colorReflector     ebiten.ColorScale
	reflectorMode      uint8
}

func NewPlayer(config config.Config, input *eInput.Handler) Player {
	CenterScreen = Vec{float64(config.WindowSizeX / 2), float64(config.WindowSizeY / 2)}
	player := Player{
		Circle:       asset.GetImage("cercle.png"),
		Reflector:    asset.GetImage("reflector2.png"),
		circleOption: &ebiten.DrawImageOptions{},
		input:        input,
	}

	circleSize := player.Circle.Bounds().Size()
	player.circleOption.GeoM.Translate(-float64(circleSize.X)/2, -float64(circleSize.Y)/2)

	player.DstPost = Vec{X: float64(config.WindowSizeX / 2), Y: float64(config.WindowSizeY / 2)}

	reflectorSize := player.Reflector.Bounds().Size()
	player.CenterPos = Vec{-float64(reflectorSize.X) / 2, -float64(reflectorSize.Y) / 2}

	player.circleOption.GeoM.Translate(player.DstPost.X, player.DstPost.Y)

	//fmt.Println(player.Circle.Bounds())
	return player
}

func (p *Player) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.Circle, p.circleOption)

	reflectorOption := &ebiten.DrawImageOptions{}
	reflectorOption.GeoM.Translate(p.CenterPos.X, p.CenterPos.Y)

	reflectorOption.GeoM.Rotate(p.angle)

	reflectorOption.GeoM.Translate(
		p.DstPost.X, p.DstPost.Y,
	)
	reflectorOption.ColorScale = ColorScales[p.reflectorMode]

	screen.DrawImage(p.Reflector, reflectorOption)
}

func (p *Player) Update() {
	mPosX, mPosY := ebiten.CursorPosition()
	//fmt.Println(p.DstPost.Sub(Vec{float64(mPosX), float64(mPosY)}))
	p.angle = float64(Vec{0, 0}.AngleToPoint(Vec{float64(mPosX) - p.DstPost.X, float64(mPosY) - p.DstPost.Y})) + PiBy2

	p.InputeUpdate()
}

func (p *Player) InputeUpdate() {
	p.reflectorMode = 0
	if p.input.ActionIsPressed(config.ActionClickLeft) {
		p.reflectorMode |= 1
	}
	if p.input.ActionIsPressed(config.ActionClickRight) {
		p.reflectorMode |= 2
	}
}
