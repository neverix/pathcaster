package util

// Texture is a texture for 3D objects
type Texture interface {
	At(uv UV) Color
}
