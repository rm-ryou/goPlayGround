package model

import "testing"

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
	expected := Color{0.55, 0.73, 1}

	res := ray.Color()
	if res != expected {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}
