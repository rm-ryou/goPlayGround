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
