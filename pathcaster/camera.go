package pathcaster

import (
	"math/rand"

	"github.com/neverix/pathcaster/util"
)

// Camera is a data type that represents a camera
type Camera struct {
	Position      util.Vec
	ScreenWidth   int
	ScreenHeight  int
	FOVMultiplier float64
}

// CastRay casts a ray from the camera's origin to a point on the screen
func (cam *Camera) CastRay(x int, y int) *util.Ray {
	if cam.FOVMultiplier == 0 {
		cam.FOVMultiplier = 1
	}
	y = cam.ScreenHeight - y
	xNoise := (1.0 / float64(cam.ScreenWidth)) *
		(rand.Float64()*2.0 - 1.0) / 2.0 * cam.FOVMultiplier
	yNoise := (1.0 / float64(cam.ScreenHeight)) *
		(rand.Float64()*2.0 - 1.0) / 2.0 * cam.FOVMultiplier
	return &util.Ray{
		Position: cam.Position,
		Direction: util.Vec{
			X: (float64(x)/float64(cam.ScreenWidth)-0.5)*
				float64(cam.ScreenWidth)/float64(cam.ScreenHeight)*
				cam.FOVMultiplier + xNoise,
			Y: (float64(y)/float64(cam.ScreenHeight)-0.5)*
				cam.FOVMultiplier + yNoise,
			Z: 1}.Norm()}
}
