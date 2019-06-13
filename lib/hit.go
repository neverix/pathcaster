package lib

// Hit is a data type for a ray hit
type Hit struct {
	Position Vec
	Normal Vec
	Shader Shader
}
