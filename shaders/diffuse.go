package shaders

import (
	"math"
	"math/rand"

	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/util"
)

// DiffuseShader is an diffuse shader
type DiffuseShader struct {
	Reflection float64
}

// Scatter is an diffuse shader scatter renderer
func (d *DiffuseShader) Scatter(r *util.Ray, h *pathcaster.Hit) *pathcaster.ScatterResult {
	v := r.Direction.Norm()
	n := h.Normal
	reflected := v.Sub(n.Mul(v.Dot(n) * 2))
	reflected = n.Lerp(reflected, math.Min(1, d.Reflection*2))
	offset := reflected.Add(util.Vec{
		X: rand.Float64(),
		Y: rand.Float64(),
		Z: rand.Float64()}.Norm().
		Mul(math.Min(1, (2 - d.Reflection*2))))
	ray := &util.Ray{
		Position:  h.Position,
		Direction: offset}
	return &pathcaster.ScatterResult{
		Albedo:    h.Texture.At(h.UV),
		Scattered: ray}
}
