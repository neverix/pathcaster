package lib

// Surface is an interface for any hittable surface
type Surface interface {
	Hit(r *Ray, tMin, tMax float64) *Hit
}

// SurfaceList is a list of surfaces
type SurfaceList []Surface

// Hit is an implementation of the hit function for surface lists
func (l SurfaceList) Hit(r *Ray, tMin, tMax float64) *Hit {
	minDistanceHit := new(Hit)
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
