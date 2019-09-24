package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"runtime"
	"sync"

	"github.com/neverix/pathcaster/lib"
)

const (
	width    = 400
	height   = 200
	samples  = 50
	maxDepth = 3
)

func main() {
	var surfaces lib.SurfaceList = []lib.Surface{
		&lib.Background{Color: lib.Color{
			R: 0.5, G: 0.7, B: 1}},
		&lib.Sphere{
			Position: lib.Vec{
				X: 0, Y: 204, Z: 15},
			Radius: 200,
			Shader: &lib.DiffuseShader{Color: lib.Color{
				R: 0, G: 1, B: 0}}},
		&lib.Sphere{
			Position: lib.Vec{
				X: 8, Y: 0, Z: 11},
			Radius: 4,
			Shader: &lib.DiffuseShader{Color: lib.Color{
				R: 1, G: 0, B: 0}}},
		&lib.Sphere{
			Position: lib.Vec{
				X: -8, Y: 0, Z: 11},
			Radius: 4,
			Shader: &lib.DiffuseShader{Color: lib.Color{
				R: 0, G: 0, B: 1}}},
		&lib.Sphere{
			Position: lib.Vec{
				X: 0, Y: 0, Z: 20},
			Radius: 4,
			Shader: &lib.ReflectiveShader{Color: lib.Color{
				R: 0.5, G: 0.3, B: 1}}}}
	camera := lib.Camera{
		Position:      lib.Vec{X: 0, Y: 0, Z: 0.3},
		ScreenWidth:   width,
		ScreenHeight:  height,
		FOVMultiplier: 1.5}

	canvas := image.NewRGBA(image.Rect(0, 0, width, height))
	var workerWg sync.WaitGroup
	CPUs := runtime.NumCPU()
	widthPerCPU := width / CPUs
	if width%CPUs != 0 {
		CPUs++
	}
	workerWg.Add(CPUs)
	for CPU := 0; CPU < CPUs; CPU++ {
		go func(id int) {
			for x := widthPerCPU * id; x < width && x < widthPerCPU*(id+1); x++ {
				for y := 0; y < height; y++ {
					var rTotal, gTotal, bTotal float64
					for i := 0; i < samples; i++ {
						r, g, b := camera.RenderPixel(surfaces, x, y, maxDepth).Unwrap()
						rTotal += r
						gTotal += g
						bTotal += b
					}
					color := color.RGBA{
						uint8(math.Sqrt(math.Min(rTotal/samples, 1)) * 255),
						uint8(math.Sqrt(math.Min(gTotal/samples, 1)) * 255),
						uint8(math.Sqrt(math.Min(bTotal/samples, 1)) * 255),
						255}
					canvas.Set(x, y, color)
				}
			}
			workerWg.Done()
		}(CPU)
	}
	workerWg.Wait()

	outputFile, err := os.Create("render.png")
	if err != nil {
		log.Fatal("Opening png file failed!")
	}

	err = png.Encode(outputFile, canvas)
	if err != nil {
		log.Fatal("Saving png file failed!")
	}
	fmt.Println("Done!")
}
