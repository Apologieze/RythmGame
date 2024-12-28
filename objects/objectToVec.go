package objects

import (
	"github.com/Waffle-osu/osu-parser/osu_parser"
	"github.com/quasilyte/gmath"
)

// rec Rectangle, object osu_parser.HitObject
var PreviousPosition Vec = Vec{0, 0}
var PreviousAngle gmath.Rad
var direction int = 1

const multiplier = gmath.Rad(1. / 200.)

func objectToVect(object osu_parser.HitObject) *Vec {
	if object.NewCombo {
		direction *= -1
	}
	position := Vec{object.Position.X, object.Position.Y}
	distance := position.DistanceTo(PreviousPosition)
	//angle := gmath.Rad(rando.FloatRange(0, 2*math.Pi))
	angle := (PreviousAngle + (gmath.Rad(distance) * gmath.Rad(direction) * multiplier)).Normalized()
	det := CenterScreen.Add(gmath.RadToVec(angle).Mulf(540))
	PreviousPosition = position
	PreviousAngle = angle
	return &det
}
