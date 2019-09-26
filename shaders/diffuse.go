package shaders

import (
	"github.com/neverix/pathcaster/util"
	"math/rand"

	"github.com/neverix/pathcaster/pathcaster"
)

// DiffuseShader is an diffuse shader
type DiffuseShader struct {
	Color util.Color
}

// Scatter is an diffuse shader scatter renderer
func (d *DiffuseShader) Scatter(r *util.Ray, h *pathcaster.Hit) *pathcaster.ScatterResult {
	offset := h.Normal.Add(util.Vec{
		X: rand.Float64(),
		Y: rand.Float64(),
		Z: rand.Float64()}).Norm()
	ray := &util.Ray{
		Position:  h.Position,
		Direction: offset}
	return &pathcaster.ScatterResult{
		Albedo:    d.Color,
		Scattered: ray}
}
