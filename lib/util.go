package lib

import "math/rand"

func randInUnitSphere() Vec {
	return Vec{
		rand.Float64(),
		rand.Float64(),
		rand.Float64()}.Norm()
}
