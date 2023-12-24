package golum

import "math"

type Hittable interface {
	Intersect() Vec3
	Normal() Vec3
}

type Sphere struct {
	center Vec3
	radius float64
}

func (s Sphere) Intersect(ray_d, ray_o Vec3) (bool, float64) {
	a := ray_d.NormSquared()
	b_vec := ray_d.MulVec(ray_o.SubVec(s.center))
	b := 2*b_vec.x + b_vec.y + b_vec.z
	c := s.center.NormSquared() + ray_o.NormSquared()*s.center.Dot(ray_o) - s.radius*s.radius

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
