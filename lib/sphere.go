package lib

import "math"

// Sphere is a sphere shape
type Sphere struct {
	Position Vec
	Radius   float64
	Shader   Shader
}

// Hit is an implementation of the hit method for a sphere
func (s *Sphere) Hit(r *Ray, tMin, tMax float64) *Hit {
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
	return &Hit{
		hitPosition,
		normal,
		s.Shader}
}
