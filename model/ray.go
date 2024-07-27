package model

import "math"

type Ray struct {
	Orig Vec3D
	Dir  Vec3D
}

func (r Ray) At(t float64) Vec3D {
	return r.Orig.Add(r.Dir.MultiNum(t))
}

func (r Ray) isHitSphere(center Vec3D, radius float64) float64 {
	oc := center.Sub(r.Orig)
	a := r.Dir.Dot(r.Dir)
	b := -2 * r.Dir.Dot(oc)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return math.NaN()
	} else {
		return (-b - math.Sqrt(discriminant)) / (2 * a)
	}
}
func (r Ray) Color() Color {
	t := r.isHitSphere(Vec3D{0, 0, -1}, 0.5)
	if t > 0 {
		n := r.At(t).Sub(Vec3D{0, 0, -1}).Norm()
		return Vec3D{n.X + 1, n.Y + 1, n.Z + 1}.MultiNum(0.5).Vec3DToColor()
	}

	unitDir := r.Dir.Norm()
	a := 0.5 * (unitDir.Y + 1)
	vec := Vec3D{1, 1, 1}.MultiNum(1 - a).Add(Vec3D{0.5, 0.7, 1}.MultiNum(a))

	return vec.Vec3DToColor()
}
