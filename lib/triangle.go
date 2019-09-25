package lib

// Triangle is a triangular shape
type Triangle struct {
	A              Vec
	B              Vec
	C              Vec
	Shader         Shader
	normal         Vec
	normalComputed bool
}

// Hit is an implementation of the hit method for a triangle
func (t *Triangle) Hit(r *Ray, tMin, tMax float64) *Hit {
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
	return &Hit{
		p,
		t.normal,
		t.Shader}
}

func checkIfOutsideTriangle(a, b, p, n Vec) bool {
	return n.Dot(b.Sub(a).Cross(p.Sub(a)).Norm()) < 0
}
