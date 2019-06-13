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
	width = 200
	height = 100
	samples = 5
)

func main() {
	var surfaces lib.SurfaceList = []lib.Surface{
		&lib.Background{Color: color.White},
		&lib.Sphere{
			Position: lib.Vec{
				X: 0, Y: 0, Z: 15},
			Radius: 4,
			Color: color.RGBA{128, 128, 128, 255}},
		&lib.Sphere{
			Position: lib.Vec{
				X: 0, Y: 94, Z: 15},
			Radius: 90,
			Color: color.RGBA{0, 255, 0, 255}}}

	canvas := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var rTotal float64
			var gTotal float64
			var bTotal float64
			for i := 0; i < samples; i++ {
				ray := &lib.Ray{
					Position: lib.Vec{
						X: 0,
						Y: 0,
						Z: 0},
					Direction: lib.Vec{
						X: float64(x) / float64(width) * 
							float64(width) / float64(height) -
							float64(width) / float64(height) / 2.0,
						Y: float64(y) / float64(height) - 0.5,
						Z: 1}}
				hit := surfaces.Hit(ray, 0.001, math.MaxFloat64)
				color := hit.Material.Render(hit, &surfaces)
				r, g, b, _ := color.RGBA()
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
