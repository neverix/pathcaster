package shaders

import (
	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/util"
)

// ReflectiveShader is a metallic shader
type ReflectiveShader struct {
	Color util.Color
}

// Scatter is an reflective material scatter renderer
func (f *ReflectiveShader) Scatter(r *util.Ray, h *pathcaster.Hit) *pathcaster.ScatterResult {
	v := r.Direction.Norm()
	n := h.Normal
	return &pathcaster.ScatterResult{
		Albedo: f.Color,
		Scattered: &util.Ray{
			Position:  h.Position,
			Direction: v.Sub(n.Mul(v.Dot(n) * 2))}}
}
