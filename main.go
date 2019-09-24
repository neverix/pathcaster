package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/neverix/pathcaster/lib"
)

const (
	width    = 400
	height   = 200
	samples  = 100
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
	renderConfig := lib.RenderConfig{
		MaxDepth: maxDepth,
		Samples:  samples}
	canvas := camera.RenderSurface(surfaces, renderConfig)

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
