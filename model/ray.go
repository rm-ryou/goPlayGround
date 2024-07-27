package model

type Ray struct {
	Orig Vec3D
	Dir  Vec3D
}

func (r Ray) At(t float64) Vec3D {
	return r.Orig.Add(r.Dir.MultiNum(t))
}

func (r Ray) Color() Color {
	unitDir := r.Dir.Norm()
	a := 0.5 * (unitDir.Y + 1)
	vec := Vec3D{1, 1, 1}.MultiNum(1 - a).Add(Vec3D{0.5, 0.7, 1}.MultiNum(a))

	return vec.Vec3DToColor()
}
