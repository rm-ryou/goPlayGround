package model

import "testing"

func TestSetFaceNorm(t *testing.T) {
	orig := Vec3D{0, 0, -1}
	dir := Vec3D{0, 0, 1}
	ray := Ray{orig, dir}

	// This value is any value
	// because this test not use this value
	hitRecord := HitRecord{
		Point:       Vec3D{0, 0, 0},
		Norm:        Vec3D{0, 0, 0},
		T:           1,
		IsFrontFace: false,
	}

	t.Run("光が球の内側にある時は負の値になる", func(t *testing.T) {
		outwardNorm := Vec3D{0, 0, 2}
		expected := outwardNorm.MultiNum(-1)

		hitRecord.SetFaceNorm(ray, outwardNorm)
		if hitRecord.Norm != expected {
			t.Errorf("expected %v", expected)
			t.Errorf("result   %v", hitRecord.Norm)
		}
	})

	t.Run("光が球の外側にある時は正の値になる", func(t *testing.T) {
		outwardNorm := Vec3D{0, 0, -2}
		expected := outwardNorm

		hitRecord.SetFaceNorm(ray, outwardNorm)
		if hitRecord.Norm != expected {
			t.Errorf("expected %v", expected)
			t.Errorf("result   %v", hitRecord.Norm)
		}
	})
}
