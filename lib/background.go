package lib

import (
	"image/color"
	"math"
)

// Background is a background surface
type Background struct {
	Color color.Color
}

// Hit is an implementation of the hit method for backgrounds
func (b *Background) Hit(r *Ray, tMin float64, tMax float64) *Hit {
	return &Hit{
		r.AtPosition(math.Pow(2.0, 64)),
		r.Direction.Neg(),
		&EmissiveMaterial{b.Color}}
}