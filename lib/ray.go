package lib

// Ray is a ray data type used for path tracing
type Ray struct {
	position Vec
    direction Vec
}

func (r Ray) atPoint(t float64) Vec {
	return r.position.add(r.direction.norm().mul(t))
}
