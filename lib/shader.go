package lib

import (
	"math/rand"
	"sort"
)

// Shader is a data type for shaders
type Shader interface {
	Scatter(*Ray, *Hit) *ScatterResult
}

// ScatterResult is a data type for scatter results
type ScatterResult struct {
	Albedo    Color
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
	return &ScatterResult{d.Color, ray}
}

// MixedShader is a shader that chooses between two shaders randomly
type MixedShader struct {
	X           []Shader
	P           []float64
	initialized bool
	totalSum    float64
}

// Scatter is a mixed shader renderer that chooses between
// shaders X with the probability of choosing the shader
// X[n] equal to P[n].
func (m *MixedShader) Scatter(r *Ray, h *Hit) *ScatterResult {
	if !m.initialized {
		sort.Slice(m.X, func(a, b int) bool {
			return m.P[a] < m.P[b]
		})
		sort.Slice(m.P, func(a, b int) bool {
			return m.P[a] < m.P[b]
		})
		var currentSum float64
		for i, p := range m.P {
			currentSum += p
			m.P[i] = currentSum
		}
		m.totalSum = currentSum
		m.initialized = true
	}
	key := rand.Float64() * m.totalSum
	shader := m.X[sort.SearchFloat64s(m.P, key)]
	return shader.Scatter(r, h)
}

// ReflectionShader is a metallic shader
type ReflectionShader struct {
	Color Color
}

// Scatter is an reflective material scatter renderer
func (f *ReflectionShader) Scatter(r *Ray, h *Hit) *ScatterResult {
	v := r.Direction.Norm()
	n := h.Normal
	return &ScatterResult{
		f.Color,
		&Ray{
			r.Position,
			v.Sub(n.Mul(v.Dot(h.Normal) * 2))}}
}
