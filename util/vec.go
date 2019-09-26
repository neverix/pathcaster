package util

import "math"

// Vec is a 3D Vector type
type Vec struct {
	X float64
	Y float64
	Z float64
}

// Add adds one vector to another
func (v Vec) Add(o Vec) Vec {
	return Vec{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Sub subtracts one vector from another
func (v Vec) Sub(o Vec) Vec {
	return Vec{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

// Mul multiplies a vector by a number
func (v Vec) Mul(o float64) Vec {
	return Vec{v.X * o, v.Y * o, v.Z * o}
}

// Div divides a vector by a number
func (v Vec) Div(o float64) Vec {
	return Vec{v.X / o, v.Y / o, v.Z / o}
}

// Neg computes the negative of a vector
func (v Vec) Neg() Vec {
	return Vec{-v.X, -v.Y, -v.Z}
}

// Mag returns the magnitude of a vector
func (v Vec) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// SqrMag returns the squared magnitude of a vector
func (v Vec) SqrMag() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Norm returns a normalized copy of a vector
func (v Vec) Norm() Vec {
	return v.Div(v.Mag())
}

// Dot computes the dot product of two vectors
func (v Vec) Dot(o Vec) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

// Cross computes the cross product of two 3D vectors
func (v Vec) Cross(o Vec) Vec {
	return Vec{
		X: v.Y*o.Z - v.Z*o.Y,
		Y: v.Z*o.X - v.X*o.Z,
		Z: v.X*o.Y - v.Y*o.X}
}
