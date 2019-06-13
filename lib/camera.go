package lib

import (
	"image/color"
	"math"
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
				float64(cam.ScreenWidth) / float64(cam.ScreenHeight),
			cam.Position.Y + float64(y) /
				float64(cam.ScreenHeight) - 0.5,
			cam.Position.Z + 1}.Norm()}
}

// RenderRay renders a ray
func (cam *Camera) RenderRay(world Surface, ray *Ray, depth, maxDepth int) color.Color {
	hit := world.Hit(ray, 0.00001, math.Inf(1))
	scatterResult := hit.Shader.Scatter(ray, hit)
	if depth > maxDepth || scatterResult == nil {
		return color.Black
	}
	if scatterResult.Scattered == nil {
		return scatterResult.Albedo
	}
	scatteredPixel := cam.RenderRay(
		world,
		scatterResult.Scattered,
		depth + 1,
		maxDepth)
	r, g, b, _ := scatterResult.Albedo.RGBA()
	r2, g2, b2, _ := scatteredPixel.RGBA()
	return color.RGBA{
		uint8(float64(r) / 65535.0 / 65535.0 * float64(r2) * 255.0),
		uint8(float64(g) / 65535.0 / 65535.0  * float64(g2) * 255.0),
		uint8(float64(b) / 65535.0 / 65535.0  * float64(b2) * 255.0),
		255}
}

// RenderPixel renders a pixel
func (cam *Camera) RenderPixel(world Surface, x, y, maxDepth int) color.Color {
	return cam.RenderRay(world, cam.CastRay(x, y), 0, maxDepth)
}
