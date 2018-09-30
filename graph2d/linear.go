package graph2d

import "math/vector"

// Linear2D represent linear in 2D plane.
type Linear2D struct {
	linear *vector.Vector3 // ax + by + c = 0
}

// NewLinear return a linear function param.
func NewLinear(p0, p1 *vector.Vector2) *Linear2D {
	if p1.Equal(p1) {
		return nil
	}
	l := &Linear2D{}
	a := p1.GetY() - p0.GetY()
	b := p0.GetX() - p1.GetX()
	c := p0.GetY()*p1.GetX() - p0.GetX()*p1.GetY()
	l.linear = vector.NewVector3(a, b, c)
	return l
}
