package lib

import "math"

// Surface is an interface for any hittable surface
type Surface interface {
	Hit(r *Ray) *Hit
}

// SurfaceList is a list of surfaces
type SurfaceList []Surface

// Hit is an implementation of the hit function for surface lists
func (l SurfaceList) Hit(r *Ray) *Hit {
	minDistanceHit := new(Hit)
	minDistance := math.Inf(1)
	for _, surf := range l {
		hit := surf.Hit(r)
		if hit != nil {
			distance := hit.position.Sub(r.Position).Mag()
			if(distance < minDistance) {
				minDistance = distance
				minDistanceHit = hit
			}
		}
	}
	return minDistanceHit
}
