package lib

// Surface is an interface for any hittable surface
type Surface interface {
	Hit(r *Ray) *Hit
}

// SurfaceList is a list of surfaces
type SurfaceList []Surface

// Hit is an implementation of the hit function for surface lists
func (l SurfaceList) Hit(r *Ray) *Hit {
	minDistanceHit := new(Hit)
	var minDistance float64 = -1
	for _, surf := range l {
		hit := surf.Hit(r)
		if hit != nil {
			distance := hit.Position.Sub(r.Position).Mag()
			if(minDistance == -1 || distance < minDistance) {
				minDistance = distance
				minDistanceHit = hit
			}
		}
	}
	return minDistanceHit
}
