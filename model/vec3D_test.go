package model

import (
	"fmt"
	"math"
	"testing"
)

func isEqualVector(vec1, vec2 Vec3D) bool {
	v1X := fmt.Sprintf("%5f", vec1.X)
	v1Y := fmt.Sprintf("%5f", vec1.Y)
	v1Z := fmt.Sprintf("%5f", vec1.Z)
	v2X := fmt.Sprintf("%5f", vec2.X)
	v2Y := fmt.Sprintf("%5f", vec2.Y)
	v2Z := fmt.Sprintf("%5f", vec2.Z)
	return v1X == v2X && v1Y == v2Y && v1Z == v2Z
}

func TestAdd(t *testing.T) {
	vec1 := Vec3D{1, 2, 3}
	vec2 := Vec3D{-1, 2, 8}
	expected := Vec3D{0, 4, 11}

	res := vec1.Add(vec2)
	if !isEqualVector(expected, res) {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", res)
	}
}

func TestSub(t *testing.T) {
	vec1 := Vec3D{1, 2, 3}
	vec2 := Vec3D{-1, 2, 6}
	expected := Vec3D{2, 0, -3}

	res := vec1.Sub(vec2)
	if !isEqualVector(expected, res) {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}

func TestMultiVec(t *testing.T) {
	vec1 := Vec3D{1, -2, 3}
	vec2 := Vec3D{0, -2, 6}
	expected := Vec3D{0, 4, 18}

	res := vec1.MultiVec(vec2)
	if !isEqualVector(expected, res) {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}

func TestDivVec(t *testing.T) {
	vec1 := Vec3D{1, -2, 3}
	vec2 := Vec3D{-1, -2, 6}
	expected := Vec3D{-1, 1, 0.5}

	res := vec1.DivVec(vec2)
	if !isEqualVector(expected, res) {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}

func TestMultiNum(t *testing.T) {
	vec := Vec3D{0, 2, -3}
	expected := Vec3D{0, 6, -9}

	res := vec.MultiNum(3)
	if !isEqualVector(expected, res) {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}

func TestDivNum(t *testing.T) {
	vec := Vec3D{0, 3, -6}
	expected := Vec3D{0, 1, -2}

	res := vec.DivNum(3)
	if !isEqualVector(expected, res) {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}

func TestDot(t *testing.T) {
	vec1 := Vec3D{0, 2, -3}
	vec2 := Vec3D{0, 2, -3}
	var expected float64 = 13

	res := vec1.Dot(vec2)
	if res != expected {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}
func TestCross(t *testing.T) {
	vec1 := Vec3D{6, 2, 1}
	vec2 := Vec3D{4, 5, 1}
	expected := Vec3D{-3, -2, 22}

	res := vec1.Cross(vec2)
	if !isEqualVector(expected, res) {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}

func TestLength(t *testing.T) {
	vec := Vec3D{1, 1, 1}
	expected := math.Sqrt(3)

	res := vec.Length()
	if res != expected {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}

func TestNorm(t *testing.T) {
	vec := Vec3D{3, 3, 3}
	expected := Vec3D{1, 1, 1}.DivNum(math.Sqrt(3))

	res := vec.Norm()
	if !isEqualVector(expected, res) {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}

func TestVec3DToColor(t *testing.T) {
	vec := Vec3D{-1, 1, 256}
	expected := Color{0, 1, 255}

	res := vec.Vec3DToColor()
	if res != expected {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}

func TestGenRandomVecWithIn(t *testing.T) {
	expected := Vec3D{5, 5, 5}
	var max, min float64 = 10, 0
	isTest := true

	res := GenRandomVecWithIn(max, min, isTest)
	if res != expected {
		t.Errorf("expected %v", expected)
		t.Errorf("result   %v", res)
	}
}
