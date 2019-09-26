package pathcaster

import "github.com/neverix/pathcaster/util"

// Surface is an interface for any hittable surface
type Surface interface {
	Hit(r *util.Ray, tMin, tMax float64) *Hit
}

