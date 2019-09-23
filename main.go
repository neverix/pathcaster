package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/neverix/pathcaster/lib"
)

const (
	width    = 100
	height   = 50
	samples  = 100
	maxDepth = 50
)

func main() {
	var surfaces lib.SurfaceList = []lib.Surface{
		&lib.Background{Color: lib.Color{
			R: 0, G: 0, B: 0}},
		&lib.Sphere{
			Position: lib.Vec{
				X: 0, Y: -9, Z: 15},
			Radius: 4,
			Shader: &lib.EmissiveShader{Color: lib.Color{
				R: 0, G: 0, B: 100}}},
		&lib.Sphere{
			Position: lib.Vec{
				X: 0, Y: 94, Z: 15},
			Radius: 90,
			Shader: &lib.DiffuseShader{Color: lib.Color{
				R: 0, G: 1, B: 0}}},
		&lib.Sphere{
			Position: lib.Vec{
				X: 0, Y: 0, Z: 15},
			Radius: 4,
			Shader: &lib.DiffuseShader{Color: lib.Color{
				R: 1, G: 0, B: 0}}}}
	camera := lib.Camera{
		Position:     lib.Vec{X: 0, Y: 0, Z: 0.3},
		ScreenWidth:  width,
		ScreenHeight: height}

	canvas := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var rTotal, gTotal, bTotal float64
			for i := 0; i < samples; i++ {
				r, g, b := camera.RenderPixel(surfaces, x, y, maxDepth).Unwrap()
				rTotal += r
				gTotal += g
				bTotal += b
			}
			color := color.RGBA{
				uint8(math.Min(rTotal/samples*255, 255)),
				uint8(math.Min(gTotal/samples*255, 255)),
				uint8(math.Min(bTotal/samples*255, 255)),
				255}
			canvas.Set(x, y, color)
		}
	}

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
