package lib

// Surface is an interface for any hittable surface
type Surface interface {
	hit(r Ray) Hit
}