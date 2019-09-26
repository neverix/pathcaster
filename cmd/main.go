package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/neverix/pathcaster/util"

	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/shaders"
	"github.com/neverix/pathcaster/surfaces"
)

const (
	width    = 100
	height   = 50
	samples  = 3
	maxDepth = 3
)

func main() {
	model, err := surfaces.ParseOBJFile("data/big/teapot.obj")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Parsed!")
	var world surfaces.SurfaceList = []pathcaster.Surface{
		&surfaces.Background{Color: util.Color{
			R: 0.5, G: 0.7, B: 1}},
		&surfaces.Sphere{
			Position: util.Vec{
				X: 0, Y: -204, Z: 15},
			Radius: 200,
			Shader: &shaders.DiffuseShader{Color: util.Color{
				R: 0.4, G: 0.8, B: 0}}},
		&surfaces.Sphere{
			Position: util.Vec{
				X: -8, Y: 0, Z: 11},
			Radius: 4,
			Shader: &shaders.DiffuseShader{Color: util.Color{
				R: 0.9, G: 0.1, B: 0}}},
		&surfaces.Sphere{
			Position: util.Vec{
				X: 8, Y: 0, Z: 11},
			Radius: 4,
			Shader: &shaders.ReflectiveShader{Color: util.Color{
				R: 0.8, G: 0.6, B: 0.2}}},
		model.ToSurface()}

	camera := pathcaster.Camera{
		Position:     util.Vec{X: 0, Y: 1, Z: -6},
		ScreenWidth:  width,
		ScreenHeight: height}
	renderConfig := pathcaster.RenderConfig{
		MaxDepth: maxDepth,
		Samples:  samples}
	fmt.Println("Rendering...")
	canvas := camera.RenderSurface(world, renderConfig)

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
