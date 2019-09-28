package shaders

import (
	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/util"
)

// EmissiveShader is an emissive shader
type EmissiveShader struct{}

// Scatter is an emissive material scatter renderer
func (d *EmissiveShader) Scatter(r *util.Ray, h *pathcaster.Hit) *pathcaster.ScatterResult {
	return &pathcaster.ScatterResult{
		Albedo:    h.Texture.At(h.UV),
		Scattered: nil}
}
