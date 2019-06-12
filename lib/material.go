package lib

import "image/color"

// Material is a data type for materials
type Material interface {
	Render(world Surface) color.Color
}

// EmissiveMaterial is an emissive shader
type EmissiveMaterial struct {
	Color color.Color
}

// Render is an emissive material renderer
func (d *EmissiveMaterial) Render(w Surface) color.Color {
	return d.Color
}
