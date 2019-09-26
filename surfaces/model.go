package surfaces

import (
	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/util"
)

// Model is a representation of a 3D model file, including
// Wavefront OBJ
type Model struct {
	Vertices []util.Vec
	Faces    []Face
	Shader   pathcaster.Shader
}

// Face is a face of a 3D model, consisting of three vertices
type Face struct {
	A, B, C int64
}

// ToSurface makes a model a renderable surface
func (m *Model) ToSurface() pathcaster.Surface {
	model := make([]pathcaster.Surface, len(m.Faces))
	for i, face := range m.Faces {
		model[i] = &Triangle{
			A:      m.Vertices[face.A-1],
			B:      m.Vertices[face.B-1],
			C:      m.Vertices[face.C-1],
			Shader: m.Shader}
	}
	return SurfaceList(model)
}
