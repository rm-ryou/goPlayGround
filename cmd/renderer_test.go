package cmd

import (
	"testing"
)

func TestCreateFormat(t *testing.T) {
	format := createFormat(256, 256) 
	expected := "P3\n256 256\n255\n"

	if format != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("result: %v", format)
	}
}
