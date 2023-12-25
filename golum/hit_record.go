package golum

type HitRecord struct {
	in_ray            Ray
	intersect, normal Vec3
	inside            bool
	material          Material
	shape             Shape
}
