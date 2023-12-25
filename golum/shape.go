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

type Triangle struct {
	a, b, c  Vec3
	normal   Vec3
	material Material
}

func NewTriangle(a, b, c Vec3, material Material) Triangle {
	normal := b.SubVec(a).Cross(c.SubVec(a)).Normalize()
	return Triangle{
		a:        a,
		b:        b,
		c:        c,
		normal:   normal,
		material: material,
	}
}

func (t Triangle) Flip() Triangle {
	return NewTriangle(t.c, t.b, t.a, t.material)
}

func (t Triangle) SetNormal(normal Vec3) Triangle {
	t.normal = normal
	return t
}

func (t Triangle) Hit(in_ray Ray) *HitRecord {
	// moller-trumbore algorithm
	edge1 := t.b.SubVec(t.a)
	edge2 := t.c.SubVec(t.a)
	h := in_ray.D.Cross(edge2)
	a := edge1.Dot(h)
	if math.Abs(a) < EPS {
		return nil
	}

	f := 1 / a
	s := in_ray.O.SubVec(t.a)
	u := f * s.Dot(h)
	if u < 0 || u > 1 {
		return nil
	}

	q := s.Cross(edge1)
	v := f * in_ray.D.Dot(q)
	if v < 0 || u+v > 1 {
		return nil
	}

	tt := f * edge2.Dot(q)
	if tt < EPS {
		return nil
	}

	intersect := in_ray.PointAt(tt)
	inside := in_ray.D.Dot(t.normal) > 0
	return &HitRecord{
		in_ray:    in_ray,
		intersect: intersect,
		normal:    t.normal,
		inside:    inside,
		material:  t.material,
		shape:     t,
	}
}
