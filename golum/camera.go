package golum

import (
	"math"
	"math/rand"
)

type Camera struct {
	pos, forward, up, right Vec3
	fovx                    float64
}

func LookAt(pos, look_at, up Vec3, fovx float64) Camera {
	forward := look_at.SubVec(pos)
	return Camera{
		pos,
		forward.Normalize(),
		up.Normalize(),
		forward.Cross(up).Normalize(),
		fovx * math.Pi / 180,
	}
}

func (cam Camera) GetRay(w, h, x, y int, rng *rand.Rand) Ray {
	forward_dis := float64(w) / math.Tan(cam.fovx/2)
	right_dis := float64(x) - float64(w)/2 + rng.Float64()
	up_dis := float64(y) - float64(h)/2 + rng.Float64()

	ray_d := Vec3{}
	ray_d = ray_d.AddVec(cam.forward.MulScalar(forward_dis))
	ray_d = ray_d.AddVec(cam.right.MulScalar(right_dis))
	ray_d = ray_d.AddVec(cam.up.MulScalar(up_dis))
	ray_d = ray_d.Normalize()

	return Ray{O: cam.pos, D: ray_d}
}
