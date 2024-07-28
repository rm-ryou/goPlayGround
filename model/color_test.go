package model

import "testing"

func TestWriteColor(t *testing.T) {
	color := Color{0, 0.16, 1}

	// gamma space
	// r: 0
	// g: 0.4
	// b: 1
	// rGammaSpace: 0
	// gGammaSpace: 0.4 * 256 = 102
	// bGammaSpace: 0.999 * 256 = 255
	expected := "0 102 255\n"

	res := color.WriteColor()
	if res != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", res)
	}
}

func TestAddColor(t *testing.T) {
	color1 := Color{0, 1, 3}
	color2 := Color{0, 2, 9}
	expected := Color{0, 3, 12}

	res := color1.Add(color2)
	if res != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", res)
	}
}

func TestMultiColor(t *testing.T) {
	color1 := Color{0, 1, 3}
	color2 := Color{0, 2, 9}
	expected := Color{0, 2, 27}

	res := color1.Multi(color2)
	if res != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", res)
	}
}

func TestMultiNumColor(t *testing.T) {
	color := Color{0, 1, 3}
	expected := Color{0, 2, 6}

	res := color.MultiNum(2)
	if res != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", res)
	}
}
