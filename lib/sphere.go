package lib

import (
	"math"
	"image/color"
)

// Sphere is a sphere
type Sphere struct {
	Position Vec
	Radius float64
	Color color.Color
}

// Hit is an implementation of the hit method for a sphere
func (s *Sphere) Hit(r *Ray) *Hit {
	offset := r.Position.Sub(s.Position)
	a := r.Direction.Dot(r.Direction)
	b := offset.Dot(r.Direction) * 2
	c := offset.Dot(offset) - s.Radius * s.Radius
	d := b * b - 4 * a * c
	if d < 0 {
		return nil
	}
	t := (-b - math.Sqrt(d)) / (a * 2)
	return &Hit{r.AtPosition(t), Vec{}, s.Color}
}
