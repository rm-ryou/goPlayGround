package model

import (
	"math"
	"testing"
)

func TestHitObjectList(t *testing.T) {
	orig := Vec3D{0, 0, -100}
	dir := Vec3D{0, 0, 1}
	ray := Ray{orig, dir}

	t.Run("オブジェクトが光と交差するときtrueを返す", func (t *testing.T) {
		objectList := new(HittableList)
		objectList.AddObject(Sphere{Vec3D{0, 0, 0}, 10})
		objectList.AddObject(Sphere{Vec3D{0, 0, 0}, 50})
		rec := new(HitRecord)

		// colision point of sphere
		// oc = {0, 0, 100}
		// a = 1
		// h = 100
		// c = 7500
		// discriminant = 2500
		// sqrt of discriminant = 50

		// info of HitRecord
		// T = 50 because the second sphere is in front of the first
		// Point = {0, 0, -50}
		// Norm = Point / 50 = {0, 0, -1}
		// IsFrontFace = true

		res := objectList.Hit(ray, math.Inf(1), 0, rec)
		if res != true {
			t.Errorf("result: %v", res)
		}

		t.Run("HitRecordの値が変化する", func(t *testing.T) {
			point := Vec3D{0, 0, -50}
			norm := Vec3D{0, 0, -1}
			var param float64 = 50
			expected := HitRecord{point, norm, param, true}

			if *rec != expected {
				t.Errorf("expected: %v", expected)
				t.Errorf("result: %v", *rec)
			}
		})
	})

	t.Run("オブジェクトが光と交差しない時falseを返す", func (t *testing.T) {
		objectList := new(HittableList)
		objectList.AddObject(Sphere{Vec3D{10, 0, -1}, 1})
		objectList.AddObject(Sphere{Vec3D{-10, 0, -1}, 2})
		rec := new(HitRecord)

		res := objectList.Hit(ray, math.Inf(1), 0, rec)
		if res != false {
			t.Errorf("result: %v", res)
		}

		t.Run("HitRecordの値は変化しない", func(t *testing.T) {
			expected := new(HitRecord)

			if *rec != *expected {
				t.Errorf("expected: %v", *expected)
				t.Errorf("result: %v", *rec)
			}
		})
	})
}
