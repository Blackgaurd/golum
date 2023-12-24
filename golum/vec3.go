package golum

import "math"

type Vec3 struct {
	x, y, z float64
}

func V(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func VFrom(val float64) Vec3 {
	return Vec3{val, val, val}
}

func (v Vec3) AddVec(other Vec3) Vec3 {
	return V(v.x+other.x, v.y+other.y, v.z+other.z)
}

func (v Vec3) SubVec(other Vec3) Vec3 {
	return V(v.x-other.x, v.y-other.y, v.z-other.z)
}

func (v Vec3) MulVec(other Vec3) Vec3 {
	return V(v.x*other.x, v.y*other.y, v.z*other.z)
}

func (v Vec3) DivVec(other Vec3) Vec3 {
	return V(v.x/other.x, v.y/other.y, v.z/other.z)
}

func (v Vec3) AddScalar(val float64) Vec3 {
	return V(v.x+val, v.y+val, v.z+val)
}

func (v Vec3) SubScalar(val float64) Vec3 {
	return V(v.x-val, v.y-val, v.z-val)
}

func (v Vec3) MulScalar(val float64) Vec3 {
	return V(v.x*val, v.y*val, v.z*val)
}

func (v Vec3) DivScalar(val float64) Vec3 {
	return V(v.x/val, v.y/val, v.z/val)
}

func (v Vec3) Cross(other Vec3) Vec3 {
	return V(v.y*other.z-v.z*other.y, v.z*other.x-v.x*other.z, v.x*other.y-v.y*other.x)
}

func (v Vec3) Dot(other Vec3) float64 {
	return v.x*other.x + v.y*other.y + v.z*other.z
}

func (v Vec3) NormSquared() float64 {
	return v.Dot(v)
}

func (v Vec3) Norm() float64 {
	return float64(math.Sqrt(float64(v.NormSquared())))
}

func (v Vec3) Normalize() Vec3 {
	norm := v.Norm()
	return v.DivScalar(norm)
}
