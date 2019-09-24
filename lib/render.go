package lib

import (
	"image"
	"image/color"
	"math"
	"runtime"
	"sync"
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
	if cam.ScreenWidth%CPUs != 0 {
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
