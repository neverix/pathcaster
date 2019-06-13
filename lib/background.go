package lib

import "math"

// Background is a background surface
type Background struct {
	Color Color
}

// Hit is an implementation of the hit method for backgrounds
func (b *Background) Hit(r *Ray, tMin, tMax float64) *Hit {
	if math.IsInf(tMax, 1) {
		return &Hit{
			r.AtPosition(math.Pow(2.0, 64)),
			r.Direction.Neg(),
			&EmissiveShader{b.Color}}
	}
	return nil
}