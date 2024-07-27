package cmd

import (
	"strconv"
	"strings"

	"github.com/rm-ryou/goPlayGround/model"
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
			color := model.Color{
				R: float64(j) / (float64(width) - 1),
				G: float64(i) / (float64(height) - 1),
				B: 0,
			}

			image.WriteString(color.WriteColor())
		}
	}
	return image.String()
}

func CreateData(width, height int) string {
	return createFormat(width, height) + createImage(width, height)
}
