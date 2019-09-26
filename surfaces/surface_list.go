package surfaces

import (
	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/util"
)

// SurfaceList is a list of surfaces
type SurfaceList []pathcaster.Surface

// Hit is an implementation of the hit function for surface lists
func (l SurfaceList) Hit(r *util.Ray, tMin, tMax float64) *pathcaster.Hit {
	minDistanceHit := new(pathcaster.Hit)
	minDistance := tMax
	for _, surf := range l {
		hit := surf.Hit(r, tMin, minDistance)
		if hit != nil {
			distance := hit.Position.Sub(r.Position).Mag()
			if distance > tMin && distance < minDistance {
				minDistanceHit = hit
				minDistance = distance
			}
		}
	}
	if minDistance == tMax {
		return nil
	}
	return minDistanceHit
}
