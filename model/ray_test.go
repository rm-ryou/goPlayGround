package model

import (
	"testing"
)

func TestAt(t *testing.T) {
	orig := Vec3D{1, 2, 5}
	dir := Vec3D{1, -1, 0}
	ray := Ray{orig, dir}
	expected := Vec3D{4, -1, 5}

	res := ray.At(3)
	if res != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", res)
	}
}

func TestColor(t *testing.T) {
	isTest := true
	t.Run("球と交差するときは球の色を出力", func(t *testing.T) {
		orig := Vec3D{0, 0, 1}
		dir := Vec3D{0, 0, -1}
		ray := Ray{orig, dir}
		objectList := new(HittableList)
		objectList.AddObject(Sphere{Vec3D{0, 0, -1}, 1})

		// point of hit sphere
		// oc = {0, 0, -2}
		// a = 1
		// h = 2
		// c = 3
		// dirsc = 1
		// root = 2 - 1 = 1
		// hitRecord.T = 1
		// hitRecord.Point = {0, 0, 0}
		// hitRecord.Norm = {0, 0, 1}

		// info of HitRecord
		// T = 50 because the second sphere is in front of the first
		// Point = {0, 0, -50}
		// Norm = Point / 50 = {0, 0, -1}
		// IsFrontFace = true

		// expect
		// ({1, 1, 1} + {0, 0, -1}) * 0.5 = {0.5, 0.5, 0}
		expected := Color{0.5, 0.5, 0.5}

		res := ray.Color(*objectList, 1, isTest)
		if res != expected {
			t.Errorf("expected: %v", expected)
			t.Errorf("result: %v", res)
		}
	})

	t.Run("球と交差しないときは背景色を出力", func(t *testing.T) {
		orig := Vec3D{0, 0, 0}
		dir := Vec3D{3, 4, 0}
		// |dir| = 5, ie. dir.Norm = {0.6, 0.8, 0}
		// a = 0.5 * (0.8 + 1) = 0.9
		// 0.1 * {1, 1, 1} + 0.9 * {0.5, 0.7, 1}
		// {0.55, 0.73, 1}
		ray := Ray{orig, dir}
		expected := Color{0.55, 0.73, 1}

		res := ray.Color(HittableList{}, 1, isTest)
		if res != expected {
			t.Errorf("expected: %v", expected)
			t.Errorf("result: %v", res)
		}
	})
}
