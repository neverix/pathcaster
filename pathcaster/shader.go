package pathcaster

import "github.com/neverix/pathcaster/util"

// Shader is a data type for shaders
type Shader interface {
	Scatter(*util.Ray, *Hit) *ScatterResult
}

// ScatterResult is a data type for scatter results
type ScatterResult struct {
	Albedo    util.Color
	Scattered *util.Ray
}
