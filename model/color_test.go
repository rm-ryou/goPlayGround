package model

import "testing"

func TestWriteColor(t *testing.T) {
	color := Color{Vec3D: Vec3D{0, 0.5, 1}}
	expected := "0 127 255\n"

	res := color.WriteColor()
	if res != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", res)
	}
}
