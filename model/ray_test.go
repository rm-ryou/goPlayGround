package model

import (
	"math"
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
	orig := Vec3D{0, 0, 0}
	dir := Vec3D{3, 4, 0}
	// |dir| = 5, ie. dir.Norm = {0.6, 0.8, 0}
	// a = 0.5 * (0.8 + 1) = 0.9
	// 0.1 * {1, 1, 1} + 0.9 * {0.5, 0.7, 1}
	// {0.55, 0.73, 1}
	ray := Ray{orig, dir}
	expected := Color{Vec3D: Vec3D{0.55, 0.73, 1}}

	res := ray.Color()
	if res != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", res)
	}
}

func TestIsHitSphere(t *testing.T) {
	orig := Vec3D{0, 0, -1}
	dir := Vec3D{0, 0, 1}
	ray := Ray{orig, dir}

	t.Run("光が球と交差するときは解を返す", func(t *testing.T) {
		center := Vec3D{0, 0, 0}
		radius := 1
		var expected float64 = 0

		res := ray.isHitSphere(center, float64(radius))
		if res != expected {
			t.Errorf("expected: %v", expected)
			t.Errorf("result: %v", res)
		}
	})

	t.Run("光が急と交差しない時はNaNを返す", func(t *testing.T) {
		center := Vec3D{2, 2, 0}
		radius := 1

		res := ray.isHitSphere(center, float64(radius))
		if !math.IsNaN(res) {
			t.Errorf("result: %v", res)
		}
	})
}
