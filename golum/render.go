package golum

import (
	"math/rand"
)

func trace(ray Ray, shapes []Shape, rng *rand.Rand, depth int) Vec3 {
	if depth <= 0 {
		return Vec3{}
	}

	hit := &HitRecord{}
	for _, shape := range shapes {
		if h := shape.Hit(ray); h != nil {
			hit = h
		}
	}
	if hit.material == nil {
		return Vec3{}
	}
	result := hit.material.Scatter(ray, hit, rng)
	if result.absorbed {
		return result.emitted
	}

	cos_theta := hit.normal.Dot(result.scattered.D)
	result_color := trace(result.scattered, shapes, rng, depth-1)
	result_color = result_color.MulVec(result.color)
	result_color = result_color.MulScalar(cos_theta * 2)
	result_color = result_color.AddVec(result.emitted)
	return result_color
}

func Render(width, height int, camera Camera, shapes []Shape, rng *rand.Rand) ImageBuffer {
	img := NewImageBuffer(width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ray := camera.GetRay(width, height, x, y, rng)
			color := trace(ray, shapes, rng, 5)
			img.AddData(x, y, color)
		}
	}
	return img
}
