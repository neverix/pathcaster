package lib

import "math/rand"

// Shader is a data type for shaders
type Shader interface {
	Scatter(*Ray, *Hit) *ScatterResult
}

// ScatterResult is a data type for scatter results
type ScatterResult struct {
	Albedo Color
	Scattered *Ray
}

// EmissiveShader is an emissive shader
type EmissiveShader struct {
	Color Color
}

// Scatter is an emissive material scatter renderer
func (d *EmissiveShader) Scatter(r *Ray, h *Hit) *ScatterResult {
	return &ScatterResult{d.Color, nil}
}

// DiffuseShader is an diffuse shader
type DiffuseShader struct {
	Color Color
}

// Scatter is an diffuse shader scatter renderer
func (d *DiffuseShader) Scatter(r *Ray, h *Hit) *ScatterResult {
	offset := h.Normal.Add(Vec{
		rand.Float64(),
		rand.Float64(),
		rand.Float64()}).Norm()
	ray := &Ray{h.Position, offset}
	color := d.Color
	return &ScatterResult{color, ray}
}
