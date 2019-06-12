package lib

import "image/color"

// Hit is a data type for a ray hit
type Hit struct {
	Position Vec
	Normal Vec
	Color color.Color
}
