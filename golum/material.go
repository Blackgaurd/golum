package golum

import (
	"math"
	"math/rand"
)

type ScatterResult struct {
	absorbed  bool
	scattered Ray
	color     Vec3
	emitted   Vec3
}

type Material interface {
	Scatter(in_ray Ray, hit *HitRecord, rng *rand.Rand) ScatterResult
}

type Emit struct {
	color Vec3
}

func NewEmit(color Vec3) Emit {
	return Emit{color}
}

func (e Emit) Scatter(in_ray Ray, hit *HitRecord, rng *rand.Rand) ScatterResult {
	return ScatterResult{
		absorbed: true,
		emitted:  e.color,
	}
}

type Diffuse struct {
	color Vec3
}

func NewDiffuse(color Vec3) Diffuse {
	return Diffuse{color}
}

func hemisphereSample(normal Vec3, rng *rand.Rand) Vec3 {
	u := rng.Float64()
	v := rng.Float64()
	theta := math.Acos(2*u-1) - math.Pi/2
	phi := 2 * math.Pi * v
	sample := Vec3{
		math.Cos(phi) * math.Cos(theta),
		math.Sin(phi) * math.Cos(theta),
		math.Sin(theta),
	}

	if sample.Dot(normal) < 0 {
		return sample.Neg()
	}
	return sample
}

func (d Diffuse) Scatter(in_ray Ray, hit *HitRecord, rng *rand.Rand) ScatterResult {
	if hit.inside {
		return ScatterResult{absorbed: true}
	}

	scattered_o := hit.intersect.AddVec(hit.normal.MulScalar(EPS))
	scattered_d := hemisphereSample(hit.normal, rng)
	return ScatterResult{
		scattered: Ray{scattered_o, scattered_d},
		color:     d.color,
	}
}
