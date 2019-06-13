package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"math"
	"github.com/neverix/pathcaster/lib"
)

const (
	width = 100
	height = 50
	samples = 100
	maxDepth = 50
)

func main() {
	var surfaces lib.SurfaceList = []lib.Surface{
		&lib.Background{Color: color.White},
		&lib.Sphere{
			Position: lib.Vec{
				X: 0, Y: 94, Z: 15},
			Radius: 90,
			Color: color.RGBA{0, 255, 0, 255}},
		&lib.Sphere{
			Position: lib.Vec{
				X: 0, Y: 0, Z: 15},
			Radius: 4,
			Color: color.RGBA{255, 0, 0, 255}}}
	camera := lib.Camera{
		Position: lib.Vec{X: 0, Y: 0, Z: 0.3},
		ScreenWidth: width,
		ScreenHeight: height}

	canvas := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var rTotal float64
			var gTotal float64
			var bTotal float64
			for i := 0; i < samples; i++ {
				r, g, b, _ := camera.RenderPixel(surfaces, x, y, maxDepth).RGBA()
				rTotal += math.Sqrt(float64(r) / 65535.0) * 255.0
				gTotal += math.Sqrt(float64(g) / 65535.0) * 255.0
				bTotal += math.Sqrt(float64(b) / 65535.0) * 255.0
			}
			color := color.RGBA{
				uint8(rTotal / samples),
				uint8(gTotal / samples),
				uint8(bTotal / samples),
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
