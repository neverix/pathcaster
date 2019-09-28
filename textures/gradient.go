package textures

import (
	"github.com/neverix/pathcaster/util"
)

// VGradient is a simple vertical linear gradient
type VGradient struct {
	Top    util.Color
	Bottom util.Color
}

// At is a texture implementation for a vertical gradient
func (v *VGradient) At(uv util.UV) util.Color {
	return v.Top.Lerp(v.Bottom, uv.V)
}
