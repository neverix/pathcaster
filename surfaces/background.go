package surfaces

import (
	"math"

	"github.com/neverix/pathcaster/util"

	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/shaders"
)

// Background is a background surface
type Background struct {
	Texture util.Texture
}

// Hit is an implementation of the hit method for backgrounds
func (b *Background) Hit(r *util.Ray, tMin, tMax float64) *pathcaster.Hit {
	if math.IsInf(tMax, 1) {
		return &pathcaster.Hit{
			Position: r.AtPosition(math.Pow(2.0, 64)),
			Normal:   r.Direction.Neg(),
			Shader:   &shaders.EmissiveShader{},
			UV: util.UV{
				U: polar(r.Direction.X, r.Direction.Z),
				V: r.Direction.Y},
			Texture: b.Texture}
	}
	return nil
}

func polar(x, y float64) float64 {
	return math.Atan2(y, x)/math.Pi/2 + 0.5
}
