package surfaces

import (
	"github.com/neverix/pathcaster/util"
	"math"

	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/shaders"
)

// Background is a background surface
type Background struct {
	Color util.Color
}

// Hit is an implementation of the hit method for backgrounds
func (b *Background) Hit(r *util.Ray, tMin, tMax float64) *pathcaster.Hit {
	if math.IsInf(tMax, 1) {
		return &pathcaster.Hit{
			Position: r.AtPosition(math.Pow(2.0, 64)),
			Normal:   r.Direction.Neg(),
			Shader: &shaders.EmissiveShader{
				Color: b.Color}}
	}
	return nil
}
