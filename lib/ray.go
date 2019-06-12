package lib

// Ray is a ray data type used for path tracing
type Ray struct {
	Position Vec
    Direction Vec
}

func (r Ray) atPoint(t float64) Vec {
	return r.Position.Add(r.Direction.Norm().Mul(t))
}
