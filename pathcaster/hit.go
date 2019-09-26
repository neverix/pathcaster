package pathcaster

import "github.com/neverix/pathcaster/util"

// Hit is a data type for a ray hit
type Hit struct {
	Position util.Vec
	Normal   util.Vec
	Shader   Shader
}
