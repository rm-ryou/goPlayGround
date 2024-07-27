package model

type HitRecord struct {
	Point       Vec3D
	Norm        Vec3D
	T           float64
	IsFrontFace bool
}

func (hr *HitRecord) SetFaceNorm(ray Ray, outwardNorm Vec3D) {
	hr.IsFrontFace = ray.Dir.Dot(outwardNorm) < 0

	if hr.IsFrontFace {
		hr.Norm = outwardNorm
	} else {
		hr.Norm = outwardNorm.MultiNum(-1)
	}
}
