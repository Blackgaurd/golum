package golum

import "math"

type Vec3 struct {
	X, Y, Z float64
}

func V(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func VFrom(val float64) Vec3 {
	return Vec3{val, val, val}
}

func (v Vec3) AddVec(other Vec3) Vec3 {
	return V(v.X+other.X, v.Y+other.Y, v.Z+other.Z)
}

func (v Vec3) SubVec(other Vec3) Vec3 {
	return V(v.X-other.X, v.Y-other.Y, v.Z-other.Z)
}

func (v Vec3) MulVec(other Vec3) Vec3 {
	return V(v.X*other.X, v.Y*other.Y, v.Z*other.Z)
}

func (v Vec3) DivVec(other Vec3) Vec3 {
	return V(v.X/other.X, v.Y/other.Y, v.Z/other.Z)
}

func (v Vec3) AddScalar(val float64) Vec3 {
	return V(v.X+val, v.Y+val, v.Z+val)
}

func (v Vec3) SubScalar(val float64) Vec3 {
	return V(v.X-val, v.Y-val, v.Z-val)
}

func (v Vec3) MulScalar(val float64) Vec3 {
	return V(v.X*val, v.Y*val, v.Z*val)
}

func (v Vec3) DivScalar(val float64) Vec3 {
	return V(v.X/val, v.Y/val, v.Z/val)
}

func (v Vec3) Cross(other Vec3) Vec3 {
	return V(v.Y*other.Z-v.Z*other.Y, v.Z*other.X-v.X*other.Z, v.X*other.Y-v.Y*other.X)
}

func (v Vec3) Dot(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
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

func (v Vec3) Abs() Vec3 {
	return V(math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z))
}

func (v Vec3) Neg() Vec3 {
	return V(-v.X, -v.Y, -v.Z)
}
