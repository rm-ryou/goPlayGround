package cmd

import (
	"strconv"
	"strings"
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

func TestCreateImage(t *testing.T) {
	image := createImage(256, 256)

	for _, row := range strings.Split(image, "\n") {
		for _, pixel := range strings.Fields(row) {
			color, err := strconv.Atoi(pixel)
			if err != nil {
				t.Error(err)
			}
			if color < 0 || color > 255 {
				t.Errorf("Color is greater than or equal to 0 and less than or equal to 255.")
				t.Errorf("result: %d", color)
			}
		}
	}
}
