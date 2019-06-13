package lib

import (
	"math"
	"math/rand"
)

// Camera is a data type that represents a camera
type Camera struct {
	Position Vec
	ScreenWidth int
	ScreenHeight int
}

// CastRay casts a ray from the camera's origin to a point on the screen
func (cam *Camera) CastRay(x int, y int) *Ray {
	return &Ray{
		Position: cam.Position,
		Direction: Vec{
			cam.Position.X + (float64(x) / float64(cam.ScreenWidth) - 0.5) * 
				float64(cam.ScreenWidth) / float64(cam.ScreenHeight) +
				(1.0 / float64(cam.ScreenWidth)) * (rand.Float64() * 2.0 - 1.0) / 2.0,
			cam.Position.Y + float64(y) /
				float64(cam.ScreenHeight) - 0.5 +
				(1.0 / float64(cam.ScreenHeight)) * (rand.Float64() * 2.0 - 1.0) / 2.0,
			cam.Position.Z + 1}.Norm()}
}

// RenderRay renders a ray
func (cam *Camera) RenderRay(world Surface, ray *Ray, depth, maxDepth int) Color {
	hit := world.Hit(ray, 0.00001, math.Inf(1))
	scatterResult := hit.Shader.Scatter(ray, hit)
	if depth > maxDepth || scatterResult == nil {
		return Color{0, 0, 0}
	}
	if scatterResult.Scattered == nil {
		return scatterResult.Albedo
	}
	scatteredPixel := cam.RenderRay(
		world,
		scatterResult.Scattered,
		depth + 1,
		maxDepth)
	return scatterResult.Albedo.Multiply(scatteredPixel)
}

// RenderPixel renders a pixel
func (cam *Camera) RenderPixel(world Surface, x, y, maxDepth int) Color {
	return cam.RenderRay(world, cam.CastRay(x, y), 0, maxDepth)
}
