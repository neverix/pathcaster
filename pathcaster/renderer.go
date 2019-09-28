package pathcaster

import (
	"image"
	"image/color"
	"math"
	"runtime"
	"sync"

	"github.com/neverix/pathcaster/util"
)

// RenderConfig is a configuration for raytracing
type RenderConfig struct {
	MaxDepth int
	Samples  int
}

// RenderSurface renders an entire image
func (cam *Camera) RenderSurface(world Surface, renderConfig RenderConfig) *image.RGBA {
	canvas := image.NewRGBA(image.Rect(0, 0, cam.ScreenWidth, cam.ScreenHeight))
	var workerWg sync.WaitGroup
	CPUs := runtime.NumCPU()
	widthPerCPU := cam.ScreenWidth / CPUs
	for cam.ScreenWidth%CPUs != 0 {
		CPUs++
	}
	workerWg.Add(CPUs)
	for CPU := 0; CPU < CPUs; CPU++ {
		go func(id int) {
			for x := widthPerCPU * id; x < cam.ScreenWidth && x < widthPerCPU*(id+1); x++ {
				for y := 0; y < cam.ScreenHeight; y++ {
					var rTotal, gTotal, bTotal float64
					for i := 0; i < renderConfig.Samples; i++ {
						r, g, b := cam.RenderPixel(world, x, y, renderConfig.MaxDepth).Unwrap()
						rTotal += r
						gTotal += g
						bTotal += b
					}
					color := color.RGBA{
						uint8(math.Sqrt(math.Min(rTotal/float64(renderConfig.Samples), 1)) * 255),
						uint8(math.Sqrt(math.Min(gTotal/float64(renderConfig.Samples), 1)) * 255),
						uint8(math.Sqrt(math.Min(bTotal/float64(renderConfig.Samples), 1)) * 255),
						255}
					canvas.Set(x, y, color)
				}
			}
			workerWg.Done()
		}(CPU)
	}
	workerWg.Wait()

	return canvas
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
	if hit.Texture == nil {
		hit.Texture = &util.Color{R: 0.9, G: 0.9, B: 0.9}
	}
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
