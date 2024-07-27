package model

import (
	"testing"
)

func TestCalcRay(t *testing.T) {
	var x, y float64 = 0, 0
	orig := Vec3D{0, 0, 0}
	expected := Ray{orig, Camera.Pixel00Loc()}

	res := Camera.CalcRay(x, y, true)
	if res != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", res)
	}
}
