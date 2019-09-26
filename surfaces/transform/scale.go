package transform

import (
	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/util"
)

// Scale scales a surface
type Scale struct {
	Child pathcaster.Surface
	Scale util.Vec
}

// Hit is an implementation of the hit method for a scaled object
func (s *Scale) Hit(r *util.Ray, tMin, tMax float64) *pathcaster.Hit {
	ray := r.Clone()
	ray.Position = ray.Position.Mul3D(s.Scale)
	ray.Direction = ray.Direction.Mul3D(s.Scale)
	return s.Child.Hit(&ray, tMin, tMax)
}
