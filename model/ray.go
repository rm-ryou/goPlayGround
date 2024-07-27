package model

type Ray struct {
	Orig Vec3D
	Dir  Vec3D
}

func (r Ray) At(t float64) Vec3D {
	return r.Orig.Add(r.Dir.MultiNum(t))
}
