package lib

import "image/color"

// Hit is a data type for a ray hit
type Hit struct {
	position Vec
	normal Vec
	color color.Color
}
