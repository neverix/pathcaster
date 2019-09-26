package surfaces

import (
	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/util"
)

// Offset offsets a surface
type Offset struct {
	Child  pathcaster.Surface
	Offset util.Vec
}

// Hit is an implementation of the hit method for an offseted object
func (o *Offset) Hit(r *util.Ray, tMin, tMax float64) *pathcaster.Hit {
	ray := r.Clone()
	ray.Position = ray.Position.Sub(o.Offset)
	return o.Child.Hit(&ray, tMin, tMax)
}
