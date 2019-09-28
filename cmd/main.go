package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/shaders"
	"github.com/neverix/pathcaster/surfaces"
	"github.com/neverix/pathcaster/surfaces/transform"
	"github.com/neverix/pathcaster/textures"
	"github.com/neverix/pathcaster/util"
)

const (
	width    = 200
	height   = 100
	samples  = 6
	maxDepth = 3
)

func main() {
	model, err := surfaces.ParseOBJFile("data/big/simulecow.obj")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Parsed!")
	var world surfaces.SurfaceList = []pathcaster.Surface{
		&surfaces.Background{Texture: &textures.VGradient{
			Top: util.Color{
				R: 0.4, G: 0.6, B: 0.95},
			Bottom: util.Color{
				R: 0.3, G: 0.5, B: 0.9}}},
		&surfaces.Sphere{
			Position: util.Vec{
				X: 0, Y: -204, Z: 15},
			Radius: 200,
			Shader: &shaders.DiffuseShader{}},
		&surfaces.Sphere{
			Position: util.Vec{
				X: -8, Y: 0, Z: 11},
			Radius: 4,
			Shader: &shaders.DiffuseShader{}},
		&surfaces.Sphere{
			Position: util.Vec{
				X: 8, Y: 0, Z: 11},
			Radius: 4,
			Shader: &shaders.DiffuseShader{Reflection: 0.9}},
		&transform.Offset{
			Child: &transform.Scale{
				Child: model.ToSurface(),
				Scale: util.Vec{
					X: 1, Y: 1, Z: -1}},
			Offset: util.Vec{
				X: 0, Y: 0, Z: 3}}}

	camera := pathcaster.Camera{
		Position:     util.Vec{X: -1, Y: 1, Z: -8},
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
