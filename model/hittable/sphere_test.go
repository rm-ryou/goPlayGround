package hittable

import (
	"testing"

	"github.com/rm-ryou/goPlayGround/model"
)

func TestHit(t *testing.T) {
	orig := model.Vec3D{0, 0, -1}
	dir := model.Vec3D{0, 0, 1}
	ray := model.Ray{orig, dir}

	t.Run("光が球と交差するときはtrueを返す", func(t *testing.T) {
		center := model.Vec3D{0, 0, 0}
		var radius float64 = 1
		sh := Sphere{center, radius}

		point := model.Vec3D{0, 0, 0}
		norm := model.Vec3D{0, 0, 0}
		var param float64 = 0
		hitRecord := model.HitRecord{point, norm, param, false}

		expectedRes := true

		res := sh.Hit(ray, 1, -1, &hitRecord)
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
			expectPoint := model.Vec3D{0, 0, -1}
			expectNorm := model.Vec3D{0, 0, -1}
			var expectParam float64 = 0
			expectHR := model.HitRecord{expectPoint, expectNorm, expectParam, false}

			if hitRecord != expectHR {
				t.Errorf("expectHR: %v", expectHR)
				t.Errorf("result: %v", hitRecord)
			}
		})
	})

	t.Run("光が急と交差しない時はfalseを返す", func(t *testing.T) {
		center := model.Vec3D{2, 2, 0}
		var radius float64 = 1
		sh := Sphere{center, radius}

		point := model.Vec3D{0, 0, 0}
		norm := model.Vec3D{0, 0, 0}
		var param float64 = 0
		hitRecord := model.HitRecord{point, norm, param, false}

		expected := false

		res := sh.Hit(ray, 1, 1, &hitRecord)
		if res != expected {
			t.Errorf("result: %v", res)
		}

		t.Run("hitReocrdの値は変化しない", func(t *testing.T) {
			expectPoint := model.Vec3D{0, 0, 0}
			expectNorm := model.Vec3D{0, 0, 0}
			var expectParam float64 = 0
			expectHR := model.HitRecord{expectPoint, expectNorm, expectParam, false}

			if hitRecord != expectHR {
				t.Errorf("expectHR: %v", expectHR)
				t.Errorf("result: %v", hitRecord)
			}
		})
	})
}
