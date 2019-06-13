package lib

import "image/color"

// Material is a data type for materials
type Material interface {
	Render(h *Hit, world Surface) color.Color
}

// EmissiveMaterial is an emissive shader
type EmissiveMaterial struct {
	Color color.Color
}

// Render is an emissive material renderer
func (d *EmissiveMaterial) Render(h *Hit, w Surface) color.Color {
	return d.Color
}

// DiffuseMaterial is an diffuse shader
type DiffuseMaterial struct {
	Color color.Color
}

// Render is an diffuse material renderer
func (d *DiffuseMaterial) Render(h *Hit, w Surface) color.Color {
	target := h.Position.Add(h.Normal).Add(randInUnitSphere())
	hit := w.Hit(&Ray{h.Position, target.Sub(h.Position)})
	r, g, b, _ := hit.Material.Render(hit, w).RGBA()
	dr, dg, db, _ := d.Color.RGBA()
	return color.RGBA{
		uint8(float64(r) / 65535.0 * float64(dr) / 65535.0 * 255.0),
		uint8(float64(g) / 65535.0 * float64(dg) / 65535.0 * 255.0),
		uint8(float64(b) / 65535.0 * float64(db) / 65535.0 * 255.0),
		255}
}
