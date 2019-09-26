package surfaces

import (
	"github.com/neverix/pathcaster/pathcaster"
	"github.com/neverix/pathcaster/util"
)

// Triangle is a triangular shape
type Triangle struct {
	A              util.Vec
	B              util.Vec
	C              util.Vec
	Shader         pathcaster.Shader
	normal         util.Vec
	normalComputed bool
}

// Hit is an implementation of the hit method for a triangle
func (t *Triangle) Hit(r *util.Ray, tMin, tMax float64) *pathcaster.Hit {
	if !t.normalComputed {
		t.normal = t.C.Sub(t.B).Cross(t.A.Sub(t.B)).Norm()
		t.normalComputed = true
	}
	d := -t.A.Dot(t.normal)
	x := r.Direction.Dot(t.normal)
	if x == 0 {
		return nil
	}
	m := -(r.Position.Dot(t.normal) + d) / x
	if m < tMin || m > tMax {
		return nil
	}
	p := r.AtPosition(m)
	if checkIfOutsideTriangle(t.A, t.B, p, t.normal) ||
		checkIfOutsideTriangle(t.C, t.A, p, t.normal) ||
		checkIfOutsideTriangle(t.B, t.C, p, t.normal) {
		return nil
	}
	return &pathcaster.Hit{
		Position: p,
		Normal:   t.normal,
		Shader:   t.Shader}
}

func checkIfOutsideTriangle(a, b, p, n util.Vec) bool {
	return n.Dot(b.Sub(a).Cross(p.Sub(a)).Norm()) < 0
}
