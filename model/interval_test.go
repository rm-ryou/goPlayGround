package model

import "testing"

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
