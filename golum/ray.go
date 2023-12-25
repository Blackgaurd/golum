package golum

type Ray struct {
	O, D Vec3
}

func (r Ray) PointAt(t float64) Vec3 {
	return r.O.AddVec(r.D.MulScalar(t))
}
