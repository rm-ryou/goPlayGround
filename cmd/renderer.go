package cmd

import (
	"strconv"
	"strings"
)

func createFormat(width, height int) string {
	var format strings.Builder

	format.WriteString("P3\n" + strconv.Itoa(width) + " " + strconv.Itoa(height) + "\n255\n")
	return format.String()
}

func createImage(width, height int) string {
	var image strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var r float64 = float64(j) / (float64(width) - 1)
			var g float64 = float64(i) / (float64(height) - 1)
			var b float64 = 0.0

			ir := int(r * 255.999)
			ig := int(g * 255.999)
			ib := int(b * 255.999)

			image.WriteString(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n")
		}
	}
	return image.String()
}

func CreateData(width, height int) string {
	return createFormat(width, height) + createImage(width, height)
}
