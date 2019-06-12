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
func (b Background) Hit(r *Ray) *Hit {
	return &Hit{r.AtPosition(math.Inf(1)), r.Direction.Neg(), b.Color}
}