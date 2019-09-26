package util

// Ray is a ray data type used for path tracing
type Ray struct {
	Position  Vec
	Direction Vec
}

// AtPosition returns the position of the ray at a given point
func (r Ray) AtPosition(p float64) Vec {
	return r.Position.Add(r.Direction.Mul(p))
}

// Clone creates a copy of a ray
func (r Ray) Clone() Ray {
	return Ray{
		Position:  r.Position,
		Direction: r.Direction}
}
