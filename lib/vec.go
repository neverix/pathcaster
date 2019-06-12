package lib

import "math"

// Vec is a 3D Vector type
type Vec struct {
	x float64
	y float64
	z float64
}

func (v Vec) add(o Vec) Vec {
	return Vec{v.x + o.x, v.y + o.y, v.z + o.z}
}

func (v Vec) sub(o Vec) Vec {
	return Vec{v.x - o.x, v.y - o.y, v.z - o.z}
}

func (v Vec) mul(o float64) Vec {
	return Vec{v.x * o, v.y * o, v.z * o}
}

func (v Vec) div(o float64) Vec {
	return Vec{v.x / o, v.y / o, v.z / o}
}

func (v Vec) mag() float64 {
    return math.Sqrt(v.x * v.x + v.y * v.y + v.z * v.z)
}

func (v Vec) norm() Vec {
	return v.div(v.mag())
}
