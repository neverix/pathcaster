package lib

// Color is an HDR replacement to Go's image/color.Color
type Color struct {
	R float64
	G float64
	B float64
}

// Unwrap turns a color into a multiple-return
func (c Color) Unwrap() (float64, float64, float64) {
	return c.R, c.G, c.B
}

// Multiply multiplies colors
func (c Color) Multiply(o Color) Color {
	return Color{c.R * o.R, c.G * o.G, c.B * o.B}
}
