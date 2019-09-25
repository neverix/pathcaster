package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/neverix/pathcaster/lib"
)

const (
	width    = 200
	height   = 100
	samples  = 32
	maxDepth = 3
)

func main() {
	pyramid, err := lib.ParseOBJFile("data/pyramid.obj")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var surfaces lib.SurfaceList = []lib.Surface{
		&lib.Background{Color: lib.Color{
			R: 0.5, G: 0.7, B: 1}},
		&lib.Sphere{
			Position: lib.Vec{
				X: 0, Y: -204, Z: 15},
			Radius: 200,
			Shader: &lib.DiffuseShader{Color: lib.Color{
				R: 0.4, G: 0.8, B: 0}}},
		&lib.Sphere{
			Position: lib.Vec{
				X: -8, Y: 0, Z: 11},
			Radius: 4,
			Shader: &lib.DiffuseShader{Color: lib.Color{
				R: 0.9, G: 0.1, B: 0}}},
		&lib.Sphere{
			Position: lib.Vec{
				X: 8, Y: 0, Z: 11},
			Radius: 4,
			Shader: &lib.ReflectiveShader{Color: lib.Color{
				R: 0.8, G: 0.6, B: 0.2}}},
		pyramid.ToSurface()}

	camera := lib.Camera{
		Position:     lib.Vec{X: 0, Y: 0, Z: 0},
		ScreenWidth:  width,
		ScreenHeight: height}
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
