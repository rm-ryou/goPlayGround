package model

import "testing"

func TestWriteColor(t *testing.T) {
	color := Color{0, 0.5, 1}
	expected := "0 128 255\n"

	res := color.WriteColor()
	if res != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", res)
	}
}
