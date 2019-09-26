package pathcaster

import (
	"math"
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

// RenderPixel renders a pixel
func (cam *Camera) RenderPixel(world Surface, x, y, maxDepth int) util.Color {
	return cam.RenderRay(world, cam.CastRay(x, y), 0, maxDepth)
}

// RenderRay renders a ray
func (cam *Camera) RenderRay(world Surface, ray *util.Ray, depth, maxDepth int) util.Color {
	if depth > maxDepth {
		return util.Color{R: 0, G: 0, B: 0}
	}
	hit := world.Hit(ray, 0.00001, math.Inf(1))
	scatterResult := hit.Shader.Scatter(ray, hit)
	if scatterResult == nil {
		return util.Color{R: 0, G: 0, B: 0}
	}
	if scatterResult.Scattered == nil {
		return scatterResult.Albedo
	}
	scatteredPixel := cam.RenderRay(
		world,
		scatterResult.Scattered,
		depth+1,
		maxDepth)
	return scatterResult.Albedo.Multiply(scatteredPixel)
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
