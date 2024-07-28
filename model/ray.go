package model

type Ray struct {
	Orig Vec3D
	Dir  Vec3D
}

func (r Ray) At(t float64) Vec3D {
	return r.Orig.Add(r.Dir.MultiNum(t))
}

func (r Ray) isHitSphere(center Vec3D, radius float64) bool {
	oc := center.Sub(r.Orig)
	a := r.Dir.Dot(r.Dir)
	b := -2 * r.Dir.Dot(oc)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c

	return discriminant >= 0
}
func (r Ray) Color() Color {
	if r.isHitSphere(Vec3D{0, 0, -1}, 0.5) {
		return Color{1, 0, 0}
	}

	unitDir := r.Dir.Norm()
	a := 0.5 * (unitDir.Y + 1)
	color := Color{Vec3D: Vec3D{1, 1, 1}.MultiNum(1 - a).Add(Vec3D{0.5, 0.7, 1}.MultiNum(a))}

	return color
}
