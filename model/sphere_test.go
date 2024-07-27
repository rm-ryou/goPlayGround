package model

import (
	"testing"
)

func TestHitSphere(t *testing.T) {
	orig := Vec3D{0, 0, -1}
	dir := Vec3D{0, 0, 1}
	ray := Ray{orig, dir}

	t.Run("光が球と交差するときはtrueを返す", func(t *testing.T) {
		center := Vec3D{0, 0, 0}
		var radius float64 = 1
		sh := Sphere{center, radius}

		point := Vec3D{0, 0, 0}
		norm := Vec3D{0, 0, 0}
		var param float64 = 0
		hitRecord := HitRecord{point, norm, param, false}

		expectedRes := true

		res := sh.Hit(ray, Interval{1, -1}, &hitRecord)
		if res != expectedRes {
			t.Errorf("result: %v", res)
		}

		// oc = 0, 0, 1
		// a = 0
		// h = 0
		// c = 0
		// discriminant = 0
		// root = 0
		// rec.Point = 0, 0, -1
		// rec.Norm = {0, 0, -1} / 1

		t.Run("hitReocrdの値は変化する", func(t *testing.T) {
			expectPoint := Vec3D{0, 0, -1}
			expectNorm := Vec3D{0, 0, -1}
			var expectParam float64 = 0
			expectHR := HitRecord{expectPoint, expectNorm, expectParam, true}

			if hitRecord != expectHR {
				t.Errorf("expectHR: %v", expectHR)
				t.Errorf("result: %v", hitRecord)
			}
		})
	})

	t.Run("光が急と交差しない時はfalseを返す", func(t *testing.T) {
		center := Vec3D{2, 2, 0}
		var radius float64 = 1
		sh := Sphere{center, radius}

		point := Vec3D{0, 0, 0}
		norm := Vec3D{0, 0, 0}
		var param float64 = 0
		hitRecord := HitRecord{point, norm, param, false}

		expected := false

		res := sh.Hit(ray, Interval{1, 1}, &hitRecord)
		if res != expected {
			t.Errorf("result: %v", res)
		}

		t.Run("hitReocrdの値は変化しない", func(t *testing.T) {
			expectPoint := Vec3D{0, 0, 0}
			expectNorm := Vec3D{0, 0, 0}
			var expectParam float64 = 0
			expectHR := HitRecord{expectPoint, expectNorm, expectParam, false}

			if hitRecord != expectHR {
				t.Errorf("expectHR: %v", expectHR)
				t.Errorf("result: %v", hitRecord)
			}
		})
	})
}
