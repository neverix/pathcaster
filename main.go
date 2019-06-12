package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"github.com/neverix/pathcaster/lib"
)

const (
	width = 200
	height = 100
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
			hit := surfaces.Hit(ray)
			canvas.Set(x, y, hit.Material.Render(&surfaces))
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
