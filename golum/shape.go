package golum

import "math"

type Shape interface {
	Hit(in_ray Ray) *HitRecord
}

type Sphere struct {
	center   Vec3
	radius   float64
	material Material
}

func NewSphere(center Vec3, radius float64, material Material) Sphere {
	return Sphere{
		center:   center,
		radius:   radius,
		material: material,
	}
}

func (s Sphere) Hit(in_ray Ray) *HitRecord {
	a := in_ray.D.NormSquared()
	b := 2 * in_ray.D.Dot(in_ray.O.SubVec(s.center))
	c := in_ray.O.SubVec(s.center).NormSquared() - s.radius*s.radius

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return nil
	}

	t := (-b - math.Sqrt(discriminant)) / (2 * a)
	if t < 0 {
		return nil
	}

	intersect := in_ray.PointAt(t)
	normal := intersect.SubVec(s.center).Normalize()
	inside := in_ray.D.Dot(normal) > 0
	return &HitRecord{
		in_ray:    in_ray,
		intersect: intersect,
		normal:    normal,
		inside:    inside,
		material:  s.material,
		shape:     s,
	}
}
