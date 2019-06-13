package lib

// Color is an HDR replacement to Go's image/color.Color
type Color Vec

// Unwrap turns a color into a multiple-return
func (c Color) Unwrap() (float64, float64, float64) {
	return c.X, c.Y, c.Z
}

// Multiply multiplies colors
func (c Color) Multiply(o Color) Color {
	return Color{c.X * o.X, c.Y * o.Y, c.Z * o.Z}
}
