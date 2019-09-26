package shaders

import (
	"github.com/neverix/pathcaster/util"
	"math/rand"
	"sort"

	"github.com/neverix/pathcaster/pathcaster"
)

// MixedShader is a shader that chooses between two shaders randomly
type MixedShader struct {
	X           []pathcaster.Shader
	P           []float64
	initialized bool
	totalSum    float64
}

// Scatter is a mixed shader renderer that chooses between
// shaders X with the probability of choosing the shader
// X[n] equal to P[n].
func (m *MixedShader) Scatter(r *util.Ray, h *pathcaster.Hit) *pathcaster.ScatterResult {
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
