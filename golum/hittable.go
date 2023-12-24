package golum

import (
	"math"
)

type Hittable interface {
	Intersect() Vec3
	Normal() Vec3
}

type Sphere struct {
	center Vec3
	radius float64
}

func NewSphere(center Vec3, radius float64) Sphere {
	return Sphere{center, radius}
}

func (s Sphere) Intersect(ray Ray) (bool, float64) {
	a := ray.D.NormSquared()
	b := 2 * ray.D.Dot(ray.O.SubVec(s.center))
	c := ray.O.SubVec(s.center).NormSquared() - s.radius*s.radius

	d := b*b - 4*a*c
	if d < 0 {
		return false, 0
	}

	t := (-b - math.Sqrt(d)) / (2 * a)
	if t < 0 {
		return false, 0
	}
	return true, t
}

func (s Sphere) Normal(point Vec3) Vec3 {
	return point.SubVec(s.center)
}
