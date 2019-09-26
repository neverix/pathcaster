package surfaces

import (
	"github.com/neverix/pathcaster/util"
	"math"

	"github.com/neverix/pathcaster/pathcaster"
)

// Sphere is a sphere shape
type Sphere struct {
	Position util.Vec
	Radius   float64
	Shader   pathcaster.Shader
}

// Hit is an implementation of the hit method for a sphere
func (s *Sphere) Hit(r *util.Ray, tMin, tMax float64) *pathcaster.Hit {
	offset := r.Position.Sub(s.Position)
	a := r.Direction.Dot(r.Direction)
	b := offset.Dot(r.Direction)
	c := offset.Dot(offset) - s.Radius*s.Radius
	d := b*b - a*c
	if d < 0 {
		return nil
	}
	t := (-b - math.Sqrt(d)) / a
	if t < tMin || t > tMax {
		t = (-b + math.Sqrt(d)) / a
	}
	if t < tMin || t > tMax {
		return nil
	}
	hitPosition := r.AtPosition(t)
	normal := hitPosition.Sub(s.Position).Norm()
	return &pathcaster.Hit{
		Position: hitPosition,
		Normal:   normal,
		Shader:   s.Shader}
}
