package model

import (
	"testing"
)

func TestSize(t *testing.T) {
	interval := Interval{10, -10}
	var expected float64 = 20

	res := interval.Size()
	if res != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", res)
	}
}

func TestIsContains(t *testing.T) {
	var n float64 = 0
	t.Run("nが[min, max]に収まる時trueを返す", func(t *testing.T) {
		interval := Interval{100, -100}

		res := interval.IsContains(n)
		if res != true {
			t.Errorf("result: %v", res)
		}
	})

	t.Run("nが[min, max]に収まらない時falseを返す", func(t *testing.T) {
		interval := Interval{-1, -100}

		res := interval.IsContains(n)
		if res != false {
			t.Errorf("result: %v", res)
		}
	})
}

func TestIsSurrounds(t *testing.T) {
	var n float64 = 0
	t.Run("nが(min, max)に収まる時trueを返す", func(t *testing.T) {
		interval := Interval{100, -100}

		res := interval.IsContains(n)
		if res != true {
			t.Errorf("result: %v", res)
		}
	})

	t.Run("nが(min, max)に収まらない時falseを返す", func(t *testing.T) {
		interval := Interval{-1, -100}

		res := interval.IsContains(n)
		if res != false {
			t.Errorf("result: %v", res)
		}
	})
}

func TestClamp(t *testing.T) {
	intensity := Interval{0.999, 0}

	t.Run("nがintensity.Minより小さい時はintensity.Minを返す", func(t *testing.T) {
		var n float64 = -1
		expected := intensity.Min

		res := intensity.Clamp(n)
		if res != expected {
			t.Errorf("expected: %v", expected)
			t.Errorf("result: %v", res)
		}
	})

	t.Run("nがintensity.Maxより大きい時はintensity.Maxを返す", func(t *testing.T) {
		var n float64 = 1
		expected := intensity.Max

		res := intensity.Clamp(n)
		if res != expected {
			t.Errorf("expected: %v", expected)
			t.Errorf("result: %v", res)
		}
	})

	t.Run("それ以外の時はnを返す", func(t *testing.T) {
		var n float64 = 0.5
		var expected float64 = 0.5

		res := intensity.Clamp(n)
		if res != expected {
			t.Errorf("expected: %v", expected)
			t.Errorf("result: %v", res)
		}
	})
}
